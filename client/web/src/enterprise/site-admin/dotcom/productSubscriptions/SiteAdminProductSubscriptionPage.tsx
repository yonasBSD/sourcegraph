import React, { useCallback, useEffect, useMemo, useState } from 'react'

import type { ConnectError } from '@connectrpc/connect'
import { mdiInformationOutline, mdiCircle, mdiPlus, mdiPencil } from '@mdi/js'
import { QueryClientProvider, type UseQueryResult } from '@tanstack/react-query'
import { useLocation, useNavigate, useParams, useSearchParams } from 'react-router-dom'

import { Timestamp } from '@sourcegraph/branded/src/components/Timestamp'
import { logger } from '@sourcegraph/common'
import type { TelemetryV2Props } from '@sourcegraph/shared/src/telemetry'
import {
    Alert,
    Button,
    Container,
    ErrorAlert,
    H3,
    Icon,
    Link,
    LoadingSpinner,
    PageHeader,
    Text,
    Tooltip,
} from '@sourcegraph/wildcard'

import { Collapsible } from '../../../../components/Collapsible'
import {
    ConnectionContainer,
    ConnectionError,
    ConnectionList,
    ConnectionLoading,
} from '../../../../components/FilteredConnection/ui'
import { PageTitle } from '../../../../components/PageTitle'
import { Timeline, type TimelineStage } from '../../../../components/Timeline'
import { useScrollToLocationHash } from '../../../../components/useScrollToLocationHash'
import { isProductLicenseExpired } from '../../../../productSubscription/helpers'
import {
    ProductSubscriptionLabel,
    productSubscriptionLabel,
} from '../../../dotcom/productSubscriptions/ProductSubscriptionLabel'
import { LicenseGenerationKeyWarning } from '../../../productSubscription/LicenseGenerationKeyWarning'

import { CodyServicesSection } from './CodyServicesSection'
import {
    queryClient,
    useArchiveEnterpriseSubscription,
    useGetEnterpriseSubscription,
    useListEnterpriseSubscriptionLicenses,
    useUpdateEnterpriseSubscription,
    type EnterprisePortalEnvironment,
} from './enterpriseportal'
import { EnterprisePortalEnvSelector, getDefaultEnterprisePortalEnv } from './EnterprisePortalEnvSelector'
import { EnterprisePortalEnvWarning } from './EnterprisePortalEnvWarning'
import {
    type EnterpriseSubscriptionCondition,
    type EnterpriseSubscriptionLicenseCondition,
    EnterpriseSubscriptionCondition_Status,
    EnterpriseSubscriptionLicenseType,
    type ListEnterpriseSubscriptionLicensesResponse,
    EnterpriseSubscriptionLicenseCondition_Status,
    type EnterpriseSubscriptionLicenseKey,
    EnterpriseSubscriptionInstanceType,
    type EnterpriseSubscriptionLicense,
} from './enterpriseportalgen/subscriptions_pb'
import { InstanceTypeBadge } from './InstanceTypeBadge'
import { SiteAdminGenerateProductLicenseForSubscriptionForm } from './SiteAdminGenerateProductLicenseForSubscriptionForm'
import { SiteAdminProductLicenseNode } from './SiteAdminProductLicenseNode'

interface Props extends TelemetryV2Props {}

export const SiteAdminProductSubscriptionPage: React.FunctionComponent<React.PropsWithChildren<Props>> = props => (
    <QueryClientProvider client={queryClient}>
        <Page {...props} />
    </QueryClientProvider>
)

const QUERY_PARAM_ENV = 'env'

/**
 * Displays a product subscription in the site admin area.
 */
const Page: React.FunctionComponent<React.PropsWithChildren<Props>> = ({ telemetryRecorder }) => {
    const navigate = useNavigate()
    const { subscriptionUUID: paramSubscriptionUUID = '' } = useParams<{ subscriptionUUID: string }>()
    useEffect(() => telemetryRecorder.recordEvent('admin.productSubscription', 'view'), [telemetryRecorder])

    const [searchParams, setSearchParams] = useSearchParams()
    const [env, setEnv] = useState<EnterprisePortalEnvironment>(
        (searchParams.get(QUERY_PARAM_ENV) as EnterprisePortalEnvironment) || getDefaultEnterprisePortalEnv()
    )
    useEffect(() => {
        const currentEnv = searchParams.get(QUERY_PARAM_ENV) as EnterprisePortalEnvironment

        searchParams.set(QUERY_PARAM_ENV, env)
        setSearchParams(searchParams)

        // HACK: env state doesn't propagate to hooks correctly, so conditionally
        // reload the page.
        // Required until we fix https://linear.app/sourcegraph/issue/CORE-245
        if (env !== currentEnv) {
            window.location.reload()
            return
        }
    }, [env, setSearchParams, searchParams])

    const { data, isFetching: isLoading, error, refetch } = useGetEnterpriseSubscription(env, paramSubscriptionUUID)

    const [showGenerate, setShowGenerate] = useState<boolean>(false)

    const licenses = useListEnterpriseSubscriptionLicenses(
        env,
        [
            {
                filter: {
                    case: 'subscriptionId',
                    value: paramSubscriptionUUID,
                },
            },
            {
                filter: {
                    // This UI only manages old-school license keys.
                    case: 'type',
                    value: EnterpriseSubscriptionLicenseType.KEY,
                },
            },
        ],
        { limit: 100, shouldLoad: !!data }
    )

    const {
        mutateAsync: archiveProductSubscription,
        isPending: archiveLoading,
        error: archiveError,
    } = useArchiveEnterpriseSubscription(env)

    const subscription = data?.subscription

    const onArchive = useCallback(async () => {
        if (!subscription) {
            return
        }
        const reason = window.prompt(
            'Do you really want to PERMANENTLY archive this subscription? All licenses associated with this subscription will be PERMANENTLY revoked, it will no longer be available for various Sourcegraph services, and changes can no longer be made to this subscription.\n\nHowever, it does NOT refund payment or cancel billing for you.\n\nEnter a revocation reason to continue.'
        )
        if (!reason || reason.length <= 3) {
            window.alert('Aborting.')
            return
        }
        try {
            telemetryRecorder.recordEvent('admin.productSubscription', 'archive')
            await archiveProductSubscription({
                reason,
                subscriptionId: subscription.id,
            })
            navigate('/site-admin/dotcom/product/subscriptions')
        } catch (error) {
            logger.error(error)
        }
    }, [subscription, archiveProductSubscription, navigate, telemetryRecorder])

    const toggleShowGenerate = useCallback((): void => setShowGenerate(previousValue => !previousValue), [])

    const {
        mutateAsync: updateEnterpriseSubscription,
        isPending: subscriptionUpdating,
        error: subscriptionUpdateError,
    } = useUpdateEnterpriseSubscription(env)

    const onLicenseUpdate = useCallback(async () => {
        await licenses.refetch()
        setShowGenerate(false)
    }, [licenses])

    const isAnythingLoading = isLoading || licenses.isLoading || subscriptionUpdating || archiveLoading

    const created = subscription?.conditions?.find(
        condition => condition.status === EnterpriseSubscriptionCondition_Status.CREATED
    )
    const archived = subscription?.conditions?.find(
        condition => condition.status === EnterpriseSubscriptionCondition_Status.ARCHIVED
    )

    const activeLicense = getActiveLicense(licenses?.data?.licenses)

    return (
        <div className="site-admin-product-subscription-page">
            <PageTitle title="Enterprise subscription" />
            <PageHeader
                headingElement="h2"
                path={[
                    { text: 'Enterprise subscriptions', to: `/site-admin/dotcom/product/subscriptions?env=${env}` },
                    { text: subscription?.displayName || subscription?.id || paramSubscriptionUUID },
                ]}
                description="This subscription tracks a single Enterprise Sourcegraph instance."
                byline={
                    subscription &&
                    created?.lastTransitionTime && (
                        <span className="text-muted">
                            Created <Timestamp date={created.lastTransitionTime.toDate()} />
                        </span>
                    )
                }
                actions={
                    <div className="align-items-end d-flex">
                        <EnterprisePortalEnvSelector env={env} setEnv={setEnv} />
                        <div>
                            <Button
                                onClick={onArchive}
                                disabled={archiveLoading || !!archived}
                                variant="danger"
                                display="block"
                            >
                                Archive
                            </Button>
                        </div>
                    </div>
                }
                className="mb-3"
            />
            {isAnythingLoading && <LoadingSpinner />}

            {archiveError && <ErrorAlert className="mt-2" error={archiveError} />}
            {subscriptionUpdateError && <ErrorAlert className="mt-2" error={subscriptionUpdateError} />}
            {error && <ErrorAlert className="mt-2" error={error} />}

            {archived && <Alert variant="danger">This subscription has been permanently archived.</Alert>}

            {subscription && (
                <>
                    <H3 className="mt-2">Details</H3>
                    <Container className="mb-3">
                        <EnterprisePortalEnvWarning env={env} actionText="managing a subscription" />
                        <table className="table mb-0">
                            <tbody>
                                <tr>
                                    <th className="text-nowrap">
                                        Subscription ID{' '}
                                        <Tooltip content="This immutable identifier represents a subscription for a single Enterprise Sourcegraph instance tracked by the Enterprise Portal service.">
                                            <Icon aria-label="Show help text" svgPath={mdiInformationOutline} />
                                        </Tooltip>
                                    </th>
                                    <td className="w-100">
                                        <span className="text-monospace">{subscription?.id}</span>
                                    </td>
                                </tr>
                                <tr>
                                    <th className="text-nowrap">
                                        Display name{' '}
                                        <Tooltip content="Brief, human-friendly, globally unique name for this subscription.">
                                            <Icon aria-label="Show help text" svgPath={mdiInformationOutline} />
                                        </Tooltip>
                                    </th>
                                    <td className="w-100">
                                        {subscription?.displayName ? (
                                            <>{subscription?.displayName}</>
                                        ) : (
                                            <span className="text-muted">Not set</span>
                                        )}
                                        {!archived && (
                                            <EditAttributeButton
                                                attribute="displayName"
                                                label="Edit display name"
                                                refetch={refetch}
                                                disabled={isAnythingLoading}
                                                telemetryRecorder={telemetryRecorder}
                                                onClick={async () => {
                                                    const displayName = window.prompt(
                                                        'Enter instance display name to assign:',
                                                        subscription?.displayName
                                                    )
                                                    if (displayName === null) {
                                                        return
                                                    }
                                                    await updateEnterpriseSubscription({
                                                        subscription: { id: subscription?.id, displayName },
                                                    })
                                                }}
                                            />
                                        )}
                                    </td>
                                </tr>
                                <tr>
                                    <th className="text-nowrap">
                                        Salesforce subscription{' '}
                                        <Tooltip content="The ID of the corresponding Salesforce subscription, of the format 'a1a...'.">
                                            <Icon aria-label="Show help text" svgPath={mdiInformationOutline} />
                                        </Tooltip>
                                    </th>
                                    <td className="w-100">
                                        {subscription?.salesforce?.subscriptionId ? (
                                            <span className="text-monospace">
                                                {subscription?.salesforce?.subscriptionId}
                                            </span>
                                        ) : (
                                            <span className="text-muted">Not set</span>
                                        )}
                                        {!archived && (
                                            <EditAttributeButton
                                                attribute="salesforceSubscriptionID"
                                                label="Edit Salesforce subscription ID"
                                                refetch={refetch}
                                                disabled={isAnythingLoading}
                                                telemetryRecorder={telemetryRecorder}
                                                onClick={async () => {
                                                    const salesforceSubscriptionID = window.prompt(
                                                        'Enter the Salesforce subscription ID to assign:',
                                                        subscription?.salesforce?.subscriptionId
                                                    )
                                                    if (salesforceSubscriptionID === null) {
                                                        return
                                                    }
                                                    if (salesforceSubscriptionID === '') {
                                                        await updateEnterpriseSubscription({
                                                            subscription: {
                                                                id: subscription?.id,
                                                            },
                                                            updateMask: {
                                                                paths: ['salesforce.subscription_id'],
                                                            },
                                                        })
                                                    } else {
                                                        await updateEnterpriseSubscription({
                                                            subscription: {
                                                                id: subscription?.id,
                                                                salesforce: {
                                                                    subscriptionId: salesforceSubscriptionID,
                                                                },
                                                            },
                                                        })
                                                    }
                                                }}
                                            />
                                        )}
                                    </td>
                                </tr>
                                <tr>
                                    <th className="text-nowrap">
                                        Instance domain{' '}
                                        <Tooltip content="The known 'external URL' of this Sourcegraph instance. This must be set manually, and is required for Cody Analytics.">
                                            <Icon aria-label="Show help text" svgPath={mdiInformationOutline} />
                                        </Tooltip>
                                    </th>
                                    <td className="w-100">
                                        {subscription?.instanceDomain ? (
                                            <Link to={`https://${subscription?.instanceDomain}`}>
                                                {subscription?.instanceDomain}
                                            </Link>
                                        ) : (
                                            <span className="text-muted">Not set</span>
                                        )}
                                        {!archived && (
                                            <EditAttributeButton
                                                attribute="instanceDomain"
                                                label="Edit instance domain"
                                                refetch={refetch}
                                                disabled={isAnythingLoading}
                                                telemetryRecorder={telemetryRecorder}
                                                onClick={async () => {
                                                    const instanceDomain = window.prompt(
                                                        'Enter instance domain to assign (leave empty to unassign):',
                                                        subscription?.instanceDomain
                                                    )
                                                    if (instanceDomain === null) {
                                                        return
                                                    }
                                                    if (instanceDomain === '') {
                                                        await updateEnterpriseSubscription({
                                                            subscription: {
                                                                id: subscription?.id,
                                                            },
                                                            updateMask: {
                                                                paths: ['instance_domain'],
                                                            },
                                                        })
                                                    } else {
                                                        await updateEnterpriseSubscription({
                                                            subscription: { id: subscription?.id, instanceDomain },
                                                        })
                                                    }
                                                }}
                                            />
                                        )}
                                    </td>
                                </tr>
                                <tr>
                                    <th className="text-nowrap">
                                        Instance type{' '}
                                        <Tooltip content="This indicates what this subscription's instance is used for.">
                                            <Icon aria-label="Show help text" svgPath={mdiInformationOutline} />
                                        </Tooltip>
                                    </th>
                                    <td className="w-100">
                                        {subscription?.instanceType ? (
                                            <InstanceTypeBadge instanceType={subscription.instanceType} />
                                        ) : (
                                            <span className="text-muted">Not set</span>
                                        )}
                                        {!archived && (
                                            <EditAttributeButton
                                                attribute="instanceType"
                                                label="Edit instance type"
                                                refetch={refetch}
                                                disabled={isAnythingLoading}
                                                telemetryRecorder={telemetryRecorder}
                                                onClick={async () => {
                                                    const types = [
                                                        EnterpriseSubscriptionInstanceType.PRIMARY,
                                                        EnterpriseSubscriptionInstanceType.SECONDARY,
                                                        EnterpriseSubscriptionInstanceType.INTERNAL,
                                                    ]
                                                    const instanceType = window.prompt(
                                                        `Enter an instance type to assign (one of: ${types
                                                            .map(type => EnterpriseSubscriptionInstanceType[type])
                                                            .join(', ')}). Leave empty to unassign.`
                                                    )
                                                    if (instanceType === null) {
                                                        return
                                                    }
                                                    if (instanceType === '') {
                                                        await updateEnterpriseSubscription({
                                                            subscription: {
                                                                id: subscription?.id,
                                                            },
                                                            updateMask: { paths: ['instance_type'] },
                                                        })
                                                        return
                                                    }
                                                    const type = types.find(
                                                        type =>
                                                            EnterpriseSubscriptionInstanceType[type].toLowerCase() ===
                                                            instanceType.toLowerCase()
                                                    )
                                                    if (!type) {
                                                        window.alert(`Invalid instance type ${instanceType}`)
                                                        return
                                                    }
                                                    await updateEnterpriseSubscription({
                                                        subscription: {
                                                            id: subscription?.id,
                                                            instanceType: type,
                                                        },
                                                    })
                                                }}
                                            />
                                        )}
                                    </td>
                                </tr>
                                <tr>
                                    <th className="text-nowrap">
                                        Active license{' '}
                                        <Tooltip content="The most recently created, non-expired, non-revoked license is considered the 'active license'.">
                                            <Icon aria-label="Show help text" svgPath={mdiInformationOutline} />
                                        </Tooltip>
                                    </th>
                                    <td className="w-100">
                                        {activeLicense ? (
                                            <>
                                                <ProductSubscriptionLabel
                                                    productName={activeLicense.license.value?.planDisplayName}
                                                    userCount={activeLicense.license.value?.info?.userCount}
                                                />{' '}
                                                - <Link to={`#${activeLicense.id}`}>view license</Link>
                                            </>
                                        ) : (
                                            <span className="text-muted">No active license</span>
                                        )}
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </Container>

                    <Collapsible
                        title={<H3>History</H3>}
                        titleAtStart={true}
                        defaultExpanded={!!archived}
                        className="mb-3"
                    >
                        <Container className="mb-3">
                            {subscription && licenses.data ? (
                                <ConditionsTimeline
                                    subscriptionConditions={subscription.conditions}
                                    licensesConditions={licenses.data.licenses.flatMap(({ id, conditions, license }) =>
                                        conditions.map(condition => ({
                                            licenseID: id,
                                            license: license.value,
                                            condition,
                                        }))
                                    )}
                                />
                            ) : (
                                <LoadingSpinner />
                            )}
                        </Container>
                    </Collapsible>

                    <CodyServicesSection
                        enterprisePortalEnvironment={env}
                        viewerCanAdminister={true}
                        productSubscriptionUUID={subscription?.id}
                        telemetryRecorder={telemetryRecorder}
                    />

                    <H3 className="d-flex align-items-start">
                        Licenses
                        <Button
                            className="ml-auto"
                            onClick={toggleShowGenerate}
                            variant="primary"
                            disabled={!!archived || archiveLoading}
                        >
                            <Icon aria-hidden={true} svgPath={mdiPlus} /> New license key
                        </Button>
                    </H3>
                    <EnterprisePortalEnvWarning env={env} actionText="managing licenses" />
                    <LicenseGenerationKeyWarning className="mb-2" />
                    <Container className="mb-2">
                        <ProductSubscriptionLicensesConnection
                            env={env}
                            licenses={licenses}
                            toggleShowGenerate={toggleShowGenerate}
                            telemetryRecorder={telemetryRecorder}
                        />
                    </Container>
                </>
            )}
            {subscription && showGenerate && (
                <SiteAdminGenerateProductLicenseForSubscriptionForm
                    env={env}
                    subscription={subscription}
                    latestLicense={licenses.data?.licenses[0] ?? undefined}
                    onGenerate={onLicenseUpdate}
                    onCancel={() => setShowGenerate(false)}
                    telemetryRecorder={telemetryRecorder}
                />
            )}
        </div>
    )
}

function getActiveLicense(
    licenses: EnterpriseSubscriptionLicense[] | undefined
): EnterpriseSubscriptionLicense | undefined {
    return licenses?.find(
        // Exists if it is the first license, has an expiry, and expiry is before
        // now
        ({ license, conditions }, idx) =>
            idx === 0 &&
            license?.value?.info?.expireTime &&
            !isProductLicenseExpired(license?.value?.info?.expireTime?.toDate()) &&
            !conditions.find(condition => condition.status === EnterpriseSubscriptionLicenseCondition_Status.REVOKED)
    )
}

interface ProductSubscriptionLicensesConnectionProps extends TelemetryV2Props {
    env: EnterprisePortalEnvironment
    licenses: UseQueryResult<ListEnterpriseSubscriptionLicensesResponse, ConnectError>
    toggleShowGenerate: () => void
}

const ProductSubscriptionLicensesConnection: React.FunctionComponent<ProductSubscriptionLicensesConnectionProps> = ({
    env,
    licenses: { data, refetch, error, isLoading },
    toggleShowGenerate,
    telemetryRecorder,
}) => {
    const location = useLocation()
    const licenseIDFromLocationHash = useMemo(() => {
        if (location.hash.length > 1) {
            return decodeURIComponent(location.hash.slice(1))
        }
        return
    }, [location.hash])
    useScrollToLocationHash(location)

    return (
        <ConnectionContainer>
            {error && <ConnectionError errors={[error.message]} />}
            {isLoading && !data && <ConnectionLoading />}
            <ConnectionList as="ul" className="list-group list-group-flush mb-0" aria-label="Subscription licenses">
                {data?.licenses?.map(node => (
                    <SiteAdminProductLicenseNode
                        env={env}
                        key={node.id}
                        node={node}
                        defaultExpanded={node.id === licenseIDFromLocationHash}
                        showSubscription={false}
                        onRevokeCompleted={refetch}
                        telemetryRecorder={telemetryRecorder}
                        isActiveLicense={node.id === getActiveLicense(data?.licenses)?.id}
                    />
                ))}
            </ConnectionList>
            {data?.licenses?.length === 0 && <NoProductLicense toggleShowGenerate={toggleShowGenerate} />}
        </ConnectionContainer>
    )
}

const NoProductLicense: React.FunctionComponent<{
    toggleShowGenerate: () => void
}> = ({ toggleShowGenerate }) => (
    <>
        <Text className="text-muted">No license key has been generated yet.</Text>
        <Button onClick={toggleShowGenerate} variant="primary">
            <Icon aria-hidden={true} svgPath={mdiPlus} /> New license key
        </Button>
    </>
)

interface ConditionsTimelineProps {
    subscriptionConditions: EnterpriseSubscriptionCondition[]
    /**
     * Combined conditions of all licenses found.
     */
    licensesConditions: {
        licenseID: string
        license: EnterpriseSubscriptionLicenseKey | undefined
        condition: EnterpriseSubscriptionLicenseCondition
    }[]
}

const ConditionsTimeline: React.FunctionComponent<ConditionsTimelineProps> = ({
    subscriptionConditions,
    licensesConditions,
}) => {
    const getSubscriptionConditionBackgroundClassName = (status: EnterpriseSubscriptionCondition_Status): string => {
        switch (status) {
            case EnterpriseSubscriptionCondition_Status.CREATED: {
                return 'bg-success'
            }
            case EnterpriseSubscriptionCondition_Status.ARCHIVED: {
                return 'bg-danger'
            }
            default: {
                return 'bg-info'
            }
        }
    }
    const getLicenseConditionBackgroundClassName = (status: EnterpriseSubscriptionLicenseCondition_Status): string => {
        switch (status) {
            case EnterpriseSubscriptionLicenseCondition_Status.CREATED: {
                return 'bg-success'
            }
            case EnterpriseSubscriptionLicenseCondition_Status.REVOKED: {
                return 'bg-danger'
            }
            default: {
                return 'bg-info'
            }
        }
    }

    const allConditions: {
        lastTransitionTime: Date
        summary: string
        details: React.ReactNode
        className: string
    }[] = subscriptionConditions
        .map(condition => ({
            lastTransitionTime: condition.lastTransitionTime!.toDate(),
            summary: `Subscription ${EnterpriseSubscriptionCondition_Status[condition.status].toLowerCase()}`,
            details: (
                <>
                    {condition.message ? (
                        <>{condition.message}</>
                    ) : (
                        <span className="text-muted">No details provided.</span>
                    )}
                </>
            ),
            className: getSubscriptionConditionBackgroundClassName(condition.status),
        }))
        .concat(
            ...licensesConditions.map(({ licenseID, license, condition }) => ({
                lastTransitionTime: condition.lastTransitionTime!.toDate(),
                summary: `License ${EnterpriseSubscriptionLicenseCondition_Status[condition.status]
                    .toLowerCase()
                    .replaceAll('_', ' ')}: ${productSubscriptionLabel(
                    license?.planDisplayName,
                    license?.info?.userCount
                )}`,
                details: (
                    <>
                        {condition.message ? (
                            <>{condition.message}</>
                        ) : (
                            <span className="text-muted">No details provided.</span>
                        )}
                        <div className="mt-3">
                            <Link to={`#${licenseID}`}>
                                View license <span className="text-monospace">{licenseID}</span>
                            </Link>
                        </div>
                    </>
                ),
                className: getLicenseConditionBackgroundClassName(condition.status),
            }))
        )
        .sort((a, b) => (a.lastTransitionTime > b.lastTransitionTime ? -1 : 1))

    const stages = allConditions?.map(
        (condition): TimelineStage => ({
            icon: <Icon aria-label="event" svgPath={mdiCircle} />,
            date: condition.lastTransitionTime.toISOString(),
            className: condition.className,

            text: condition.summary,
            details: <Container>{condition.details}</Container>,
        })
    )

    return <Timeline showDurations={false} stages={stages} />
}

interface EditAttributeButtonProps extends TelemetryV2Props {
    attribute: string
    label: string
    onClick: () => Promise<void>
    refetch: () => void
    disabled?: boolean
}

const EditAttributeButton: React.FunctionComponent<EditAttributeButtonProps> = ({
    attribute,
    label,
    onClick,
    refetch,
    disabled,
    telemetryRecorder,
}) => (
    <Button
        size="sm"
        variant="link"
        aria-label={label}
        className="ml-1"
        disabled={disabled}
        onClick={async () => {
            await onClick()
            telemetryRecorder.recordEvent('admin.productSubscription', 'edit', {
                privateMetadata: { attribute },
            })
            refetch()
        }}
    >
        <Icon aria-hidden={true} svgPath={mdiPencil} />
    </Button>
)
