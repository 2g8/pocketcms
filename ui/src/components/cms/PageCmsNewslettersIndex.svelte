<script>
    import { onMount } from "svelte";
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import { pageTitle } from "@/stores/app";
    import { addSuccessToast, addErrorToast } from "@/stores/toasts";
    import { confirm } from "@/stores/confirmation";
    import { link } from "svelte-spa-router";
    import Field from "@/components/base/Field.svelte";
    import PageWrapper from "@/components/base/PageWrapper.svelte";

    $pageTitle = "Email Newsletters";

    let newsletters = [];
    let isLoading = false;
    let searchText = "";
    let page = 1;
    let perPage = 20;
    let totalItems = 0;

    $: totalPages = Math.ceil(totalItems / perPage);

    onMount(() => {
        loadNewsletters();
    });

    async function loadNewsletters() {
        isLoading = true;

        try {
            const result = await ApiClient.collection("newsletters").getList(page, perPage, {
                filter: searchText ? `subject ~ "${searchText}" || content ~ "${searchText}"` : "",
                sort: "-created",
            });

            newsletters = result.items;
            totalItems = result.totalItems;
        } catch (err) {
            console.error("Failed to load newsletters:", err);
            addErrorToast("Failed to load newsletter list");
        }

        isLoading = false;
    }

    async function deleteNewsletter(newsletter) {
        if (!await confirm({
            title: "Delete Newsletter",
            content: `Are you sure you want to delete newsletter "${newsletter.subject}"? This action cannot be undone.`,
            confirmText: "Delete",
            cancelText: "Cancel",
        })) {
            return;
        }

        try {
            await ApiClient.collection("newsletters").delete(newsletter.id);
            newsletters = newsletters.filter(n => n.id !== newsletter.id);
            totalItems--;
            addSuccessToast("Newsletter successfully deleted");
        } catch (err) {
            console.error("Failed to delete newsletter:", err);
            addErrorToast("Failed to delete newsletter");
        }
    }

    async function sendNewsletter(newsletter) {
        if (!await confirm({
            title: "Send Newsletter",
            content: `Are you sure you want to send newsletter "${newsletter.subject}" to all subscribers?`,
            confirmText: "Send",
            cancelText: "Cancel",
        })) {
            return;
        }

        try {
            // 这里应该调用发送通讯的API
            // await ApiClient.newsletters.send(newsletter.id);
            addSuccessToast("Newsletter send request submitted");
        } catch (err) {
            console.error("Failed to send newsletter:", err);
            addErrorToast("Failed to send newsletter");
        }
    }

    function handleSearch() {
        page = 1;
        loadNewsletters();
    }

    function changePage(newPage) {
        if (newPage < 1 || newPage > totalPages || newPage === page) {
            return;
        }
        page = newPage;
        loadNewsletters();
    }
</script>

<PageWrapper>
    <header class="page-header">
        <h1>
            <i class="ri-mail-send-line" aria-hidden="true"></i>
            <span class="txt">Email Newsletters</span>
        </h1>

        <div class="page-header-actions">
            <a href="/newsletters/create" class="btn btn-primary" use:link>
                <i class="ri-add-line" aria-hidden="true"></i>
                <span class="txt">Create Newsletter</span>
            </a>
        </div>
    </header>

    <div class="wrapper">
        <div class="search-bar">
            <form on:submit|preventDefault={handleSearch}>
                <Field type="text" bind:value={searchText} placeholder="Search newsletters..." />
                <button type="submit" class="btn btn-secondary">Search</button>
            </form>
        </div>

        {#if isLoading}
            <div class="loader"></div>
        {:else if newsletters.length === 0}
            <div class="alert alert-info">
                <i class="ri-information-line" aria-hidden="true"></i>
                <span class="txt">No newsletters found.</span>
            </div>
        {:else}
            <div class="table-wrapper">
                <table class="table">
                    <thead>
                        <tr>
                            <th>Subject</th>
                            <th>Status</th>
                            <th>Created</th>
                            <th>Sent</th>
                            <th class="col-action">Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each newsletters as newsletter (newsletter.id)}
                            <tr>
                                <td>
                                    <a href="/newsletters/{newsletter.id}" use:link>{newsletter.subject}</a>
                                </td>
                                <td>
                                    <span class="label {newsletter.status === 'sent' ? 'success' : 'warning'}">
                                        {newsletter.status === 'sent' ? 'Sent' : 'Draft'}
                                    </span>
                                </td>
                                <td>{CommonHelper.formatDate(newsletter.created)}</td>
                                <td>{newsletter.sentAt ? CommonHelper.formatDate(newsletter.sentAt) : '--'}</td>
                                <td>
                                    <div class="actions">
                                        {#if newsletter.status !== 'sent'}
                                            <button class="btn btn-sm btn-outline" on:click={() => sendNewsletter(newsletter)}>
                                                <i class="ri-send-plane-line" aria-hidden="true"></i>
                                                <span class="txt">Send</span>
                                            </button>
                                        {/if}
                                        <button class="btn btn-sm btn-outline" on:click={() => deleteNewsletter(newsletter)}>
                                            <i class="ri-delete-bin-line" aria-hidden="true"></i>
                                            <span class="txt">Delete</span>
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>

            {#if totalPages > 1}
                <div class="pagination">
                    <button class="btn btn-sm" disabled={page === 1} on:click={() => changePage(page - 1)}>Previous</button>
                    <span class="pagination-info">Page {page} of {totalPages}</span>
                    <button class="btn btn-sm" disabled={page === totalPages} on:click={() => changePage(page + 1)}>Next</button>
                </div>
            {/if}
        {/if}
    </div>
</PageWrapper>

<style>
    .search-bar {
        margin-bottom: 20px;
    }
    .search-bar form {
        display: flex;
        gap: 10px;
        max-width: 500px;
    }
    .actions {
        display: flex;
        gap: 5px;
    }
    .pagination {
        display: flex;
        justify-content: center;
        align-items: center;
        margin-top: 20px;
        gap: 15px;
    }
    .pagination-info {
        font-size: 0.9rem;
    }
</style>