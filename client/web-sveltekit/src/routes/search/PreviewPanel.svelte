<script lang="ts" context="module">
    import type { Range } from '$lib/codemirror/static-highlights'

    function extractHighlightedRanges(result: ContentMatch | SymbolMatch | PathMatch): Range[] {
        if (result.type !== 'content') {
            return []
        }
        return result.chunkMatches?.flatMap(chunkMatch => chunkMatch.ranges) || []
    }
</script>

<script lang="ts">
    import { from } from 'rxjs'

    import { noOpTelemetryRecorder } from '@sourcegraph/shared/src/telemetry'

    import CodeMirrorBlob from '$lib/CodeMirrorBlob.svelte'
    import { isErrorLike } from '$lib/common'
    import { getGraphQLClient, mapOrThrow, toGraphQLResult } from '$lib/graphql'
    import Icon from '$lib/Icon.svelte'
    import LoadingSpinner from '$lib/LoadingSpinner.svelte'
    import {
        getFileMatchUrl,
        createCodeIntelAPI,
        type ContentMatch,
        type PathMatch,
        type SymbolMatch,
    } from '$lib/shared'
    import { settings } from '$lib/stores'
    import { toReadable } from '$lib/utils/promises'
    import { Alert } from '$lib/wildcard'

    // We are reusing the queries from the blob page to ensure we take advantage of
    // the cache. It may make sense to pass loaders into this function
    // rather than adding a dependency on the blob page.
    import {
        BlobFileViewBlobQuery,
        BlobFileViewHighlightedFileQuery,
    } from '../[...repo=reporev]/(validrev)/(code)/-/blob/[...path]/page.gql'

    import { getSearchResultsContext } from './searchResultsContext'

    export let result: ContentMatch | SymbolMatch | PathMatch

    const searchResultContext = getSearchResultsContext()

    const client = getGraphQLClient()
    $: codeIntelAPI = createCodeIntelAPI({
        settings: setting => (isErrorLike($settings?.final) ? undefined : $settings?.final?.[setting]),
        requestGraphQL(options) {
            return from(client.query(options.request, options.variables).then(toGraphQLResult))
        },
        telemetryRecorder: noOpTelemetryRecorder,
    })

    $: blobStore = toReadable(
        client
            .query(BlobFileViewBlobQuery, {
                repoName: result.repository,
                revspec: result.commit ?? '',
                path: result.path,
            })
            .then(mapOrThrow(result => result.data?.repository?.commit?.blob ?? null))
    )

    $: highlightStore = toReadable(
        client
            .query(BlobFileViewHighlightedFileQuery, {
                repoName: result.repository,
                revspec: result.commit ?? '',
                path: result.path,
                disableTimeout: false,
            })
            .then(mapOrThrow(result => result.data?.repository?.commit?.blob?.highlight.lsif ?? ''))
    )
</script>

<div class="preview-container">
    <div class="header">
        <h3>File Preview</h3>
        <button data-testid="preview-close" on:click={() => searchResultContext.setPreview(null)}>
            <Icon icon={ILucideX} --icon-size="16px" inline />
        </button>
    </div>
    <div class="file-link">
        <small>
            <a href={getFileMatchUrl(result)}>{result.path}</a>
        </small>
    </div>
    <div class="content">
        {#if $blobStore.pending}
            <LoadingSpinner />
        {:else if $blobStore.error}
            <Alert variant="danger">
                Unable to load file data: {$blobStore.error}
            </Alert>
        {:else if $blobStore.value}
            <CodeMirrorBlob
                blobInfo={{
                    repoName: result.repository,
                    commitID: result.commit ?? '',
                    revision: '',
                    filePath: result.path,
                    content: $blobStore.value.content ?? '',
                    languages: $blobStore.value.languages ?? [],
                }}
                highlights={$highlightStore.value ?? ''}
                {codeIntelAPI}
                staticHighlightRanges={extractHighlightedRanges(result)}
            />
        {/if}
    </div>
</div>

<style lang="scss">
    .preview-container {
        display: flex;
        flex-direction: column;
        height: 100%;
    }

    .header {
        // Keep in sync with the action bar
        height: 3rem;
        flex: none;

        display: flex;
        align-items: center;
        padding: 0 0.5rem;
        border-bottom: 1px solid var(--border-color);

        h3 {
            margin: 0;
            margin-right: auto;
            height: fit-content;
        }

        button {
            background-color: transparent;
            color: var(--text-muted);
            border: none;
            &:hover {
                color: var(--body-color);
            }
        }
    }

    .file-link {
        padding: 0.25rem 0.5rem;
        border-bottom: 1px solid var(--border-color);
        flex: none;
    }

    .content {
        flex: 1;
        min-height: 0;
        background-color: var(--code-bg);
    }
</style>
