import * as React from 'react'

import { Badge, LinkOrSpan } from '@sourcegraph/wildcard'

import type { EnterprisePortalEnvironment } from './enterpriseportal'
import {
    type EnterpriseSubscription,
    EnterpriseSubscriptionCondition_Status,
} from './enterpriseportalgen/subscriptions_pb'
import { InstanceTypeBadge } from './InstanceTypeBadge'

import styles from './SiteAdminProductSubscriptionNode.module.scss'

export const SiteAdminProductSubscriptionNodeHeader: React.FunctionComponent<React.PropsWithChildren<unknown>> = () => (
    <thead>
        <tr>
            <th>Display name</th>
            <th>Salesforce subscription</th>
            <th>Instance type</th>
            <th>Instance domain</th>
        </tr>
    </thead>
)

export interface SiteAdminProductSubscriptionNodeProps {
    env: EnterprisePortalEnvironment
    node: EnterpriseSubscription
}

/**
 * Displays a product subscription in a connection in the site admin area.
 */
export const SiteAdminProductSubscriptionNode: React.FunctionComponent<
    React.PropsWithChildren<SiteAdminProductSubscriptionNodeProps>
> = ({ env, node }) => {
    const archived = node.conditions.find(
        condition => condition.status === EnterpriseSubscriptionCondition_Status.ARCHIVED
    )

    return (
        <tr className={styles.row}>
            <td>
                {archived && (
                    <Badge variant="danger" small={true} className="mr-2">
                        Archived
                    </Badge>
                )}
                <LinkOrSpan to={`/site-admin/dotcom/product/subscriptions/${node.id}?env=${env}`} className="mr-2">
                    <strong>{node.displayName}</strong>
                </LinkOrSpan>
            </td>
            <td className="text-nowrap">
                {node?.salesforce?.subscriptionId ? (
                    <span className="text-monospace mr-2">{node?.salesforce?.subscriptionId}</span>
                ) : (
                    <span className="text-muted mr-2">Not set</span>
                )}
            </td>
            <td className="text-nowrap">
                {node?.instanceType ? (
                    <InstanceTypeBadge instanceType={node.instanceType} className="mr-2" />
                ) : (
                    <span className="text-muted mr-2">Not set</span>
                )}
            </td>
            <td className="text-nowrap">
                {node?.instanceDomain ? (
                    <small>{node?.instanceDomain}</small>
                ) : (
                    <span className="text-muted mr-2">Not set</span>
                )}
            </td>
        </tr>
    )
}
