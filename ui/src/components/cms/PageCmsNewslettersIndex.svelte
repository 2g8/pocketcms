<script>
    import { tick } from "svelte";
    import { querystring } from "svelte-spa-router";
    import CommonHelper from "@/utils/CommonHelper";    
    import ApiClient from "@/utils/ApiClient";
    import tooltip from "@/actions/tooltip";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import RefreshButton from "@/components/base/RefreshButton.svelte";
    import Searchbar from "@/components/base/Searchbar.svelte";
    import CollectionDocsPanel from "@/components/collections/CollectionDocsPanel.svelte";
    import CollectionUpsertPanel from "@/components/collections/CollectionUpsertPanel.svelte";
    import CollectionsSidebar from "@/components/collections/CollectionsSidebar.svelte";
    import RecordPreviewPanel from "@/components/records/RecordPreviewPanel.svelte";
    import RecordUpsertPanel from "@/components/records/RecordUpsertPanel.svelte";
    import RecordsCount from "@/components/records/RecordsCount.svelte";
    import RecordsList from "@/components/records/RecordsList.svelte";
    import { hideControls, pageTitle } from "@/stores/app";
    import { addSuccessToast, addErrorToast } from "@/stores/toasts";
    import { superuser } from "@/stores/superuser";
    import {
        activeCollection,
        changeActiveCollectionByIdOrName,
        collections,
        isCollectionsLoading,
        loadCollections,
    } from "@/stores/collections";

    const initialQueryParams = new URLSearchParams($querystring);

    let collectionUpsertPanel;
    let collectionDocsPanel;
    let recordUpsertPanel;
    let recordPreviewPanel;
    let recordsList;
    let recordsCount;
    let filter = initialQueryParams.get("filter") || "";
    let sort = initialQueryParams.get("sort") || "-@rowid";
    let selectedCollectionIdOrName = 'newsletters'; // 修改为 newsletters
    let totalCount = 0; // used to manully change the count without the need of reloading the recordsCount component

    loadCollections(selectedCollectionIdOrName);

    $: reactiveParams = new URLSearchParams($querystring);

    $: collectionQueryParam = 'newsletters'; // 修改为 newsletters

    $: if (
        !$isCollectionsLoading &&
        collectionQueryParam &&
        collectionQueryParam != selectedCollectionIdOrName &&
        collectionQueryParam != $activeCollection?.id &&
        collectionQueryParam != $activeCollection?.name
    ) {
        changeActiveCollectionByIdOrName(collectionQueryParam);
    }

    // reset filter and sort on collection change
    $: if (
        $activeCollection?.id &&
        selectedCollectionIdOrName != $activeCollection.id &&
        selectedCollectionIdOrName != $activeCollection.name
    ) {
        reset();
    }

    $: if ($activeCollection?.id) {
        normalizeSort();
    }

    $: if (!$isCollectionsLoading && initialQueryParams.get("recordId")) {
        showRecordById(initialQueryParams.get("recordId"));
    }

    // keep the url params in sync
    $: if (!$isCollectionsLoading && (sort || filter || $activeCollection?.id)) {
        //updateQueryParams();
    }

    $: $pageTitle = "Newsletters"; // 修改为 Newsletters

    async function showRecordById(recordId) {
        await tick(); // ensure that the reactive component params are resolved

        $activeCollection?.type === "view"
            ? recordPreviewPanel.show(recordId)
            : recordUpsertPanel?.show(recordId);
    }

    function reset() {
        selectedCollectionIdOrName = $activeCollection?.id;
        filter = "";
        sort = "-@rowid";

        normalizeSort();

        updateQueryParams({ recordId: null });

        // close any open collection panels
        collectionUpsertPanel?.forceHide();
        collectionDocsPanel?.hide();
    }

    // ensures that the sort fields exist in the collection
    async function normalizeSort() {
        if (!sort) {
            return; // nothing to normalize
        }

        const collectionFields = CommonHelper.getAllCollectionIdentifiers($activeCollection);

        const sortFields = sort.split(",").map((f) => {
            if (f.startsWith("+") || f.startsWith("-")) {
                return f.substring(1);
            }
            return f;
        });

        // invalid sort expression or missing sort field
        if (sortFields.filter((f) => collectionFields.includes(f)).length != sortFields.length) {
            if ($activeCollection?.type != "view") {
                sort = "-@rowid"; // all collections with exception to the view has this field
            } else if (collectionFields.includes("created")) {
                // common autodate field
                sort = "-created";
            } else {
                sort = "";
            }
        }
    }

    function updateQueryParams(extra = {}) {
        const queryParams = Object.assign(
            {
                collection: $activeCollection?.id || "",
                filter: filter,
                sort: sort,
            },
            extra,
        );

        CommonHelper.replaceHashQueryParams(queryParams);
    }
</script>

{#if $isCollectionsLoading && !$collections.length}
    <PageWrapper center>
        <div class="placeholder-section m-b-base">
            <span class="loader loader-lg" />
            <h1>Loading collections...</h1>
        </div>
    </PageWrapper>
{:else if !$collections.length}
    <PageWrapper center>
        <div class="placeholder-section m-b-base">
            <div class="icon">
                <i class="ri-database-2-line" />
            </div>
            {#if $hideControls}
                <h1 class="m-b-10">You don't have any collections yet.</h1>
            {:else}
                <h1 class="m-b-10">Create your first collection to add records!</h1>
                <button
                    type="button"
                    class="btn btn-expanded-lg btn-lg"
                    on:click={() => collectionUpsertPanel?.show()}
                >
                    <i class="ri-add-line" />
                    <span class="txt">Create new collection</span>
                </button>
            {/if}
        </div>
    </PageWrapper>
{:else}

    <PageWrapper class="flex-content">
        <header class="page-header">
            <nav class="breadcrumbs">
                <div class="breadcrumb-item">PocketCMS</div>
                <div class="breadcrumb-item">{$activeCollection.name ? $activeCollection.name.charAt(0).toUpperCase() + $activeCollection.name.slice(1) : ''}</div>
            </nav>

            <div class="inline-flex gap-5">
                {#if !$hideControls}
                    <button
                        type="button"
                        aria-label="Edit collection"
                        class="btn btn-transparent btn-circle"
                        use:tooltip={{ text: "Edit collection", position: "right" }}
                        on:click={() => collectionUpsertPanel?.show($activeCollection)}
                    >
                        <i class="ri-settings-4-line" />
                    </button>
                {/if}

                <RefreshButton
                    on:refresh={() => {
                        recordsList?.load();
                        recordsCount?.reload();
                    }}
                />
            </div>

            <div class="btns-group">
                {#if $activeCollection.type !== "view"}
                    <button type="button" class="btn btn-expanded" on:click={() => recordUpsertPanel?.show()}>
                        <i class="ri-add-line" />
                        <span class="txt">Add Newsletter</span>
                    </button>
                {/if}
            </div>
        </header>

        <div class="clearfix m-b-sm" />
            <Searchbar
                value={filter}
                autocompleteCollection={$activeCollection}
                on:submit={(e) => (filter = e.detail)}
            />
            <RecordsList
                bind:this={recordsList}
                collection={$activeCollection}
                showColumns="id,name,status,visibility,created_at,updated_at"
                bind:filter
                bind:sort
                on:select={(e) => {
                    updateQueryParams({
                        recordId: e.detail.id,
                    });

                    let showModel = e.detail._partial ? e.detail.id : e.detail;

                    $activeCollection.type === "view"
                        ? recordPreviewPanel?.show(showModel)
                        : recordUpsertPanel?.show(showModel);
                }}
                on:delete={() => {
                    recordsCount?.reload();
                }}
                on:new={() => recordUpsertPanel?.show()}
            />

        <svelte:fragment slot="footer">
            <RecordsCount
                bind:this={recordsCount}
                class="m-r-auto txt-sm txt-hint"
                collection={$activeCollection}
                {filter}
                bind:totalCount
            />
        </svelte:fragment>
    </PageWrapper>
{/if}

<CollectionUpsertPanel
    bind:this={collectionUpsertPanel}
    on:truncate={() => {
        recordsList?.load();
        recordsCount?.reload();
    }}
/>

<CollectionDocsPanel bind:this={collectionDocsPanel} />

<RecordUpsertPanel
    bind:this={recordUpsertPanel}
    collection={$activeCollection}
    on:hide={() => {
        updateQueryParams({ recordId: null });
    }}
    on:save={(e) => {
        if (filter) {
            // if there is applied filter, reload the count since we
            // don't know after the save whether the record satisfies it
            recordsCount?.reload();
        } else if (e.detail.isNew) {
            totalCount++;
        }

        recordsList?.reloadLoadedPages();
    }}
    on:delete={(e) => {
        if (!filter || recordsList?.hasRecord(e.detail.id)) {
            totalCount--;
        }

        recordsList?.reloadLoadedPages();
    }}
/>

<RecordPreviewPanel
    bind:this={recordPreviewPanel}
    collection={$activeCollection}
    on:hide={() => {
        updateQueryParams({ recordId: null });
    }}
/>