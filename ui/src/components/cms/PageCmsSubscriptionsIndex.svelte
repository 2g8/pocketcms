<script>
    import { onMount } from "svelte";
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import { pageTitle } from "@/stores/app";
    import { addSuccessToast, addErrorToast } from "@/stores/toasts";
    import { confirm } from "@/stores/confirmation";
    import Field from "@/components/base/Field.svelte";
    import PageWrapper from "@/components/base/PageWrapper.svelte";

    $pageTitle = "Subscription Management";

    let subscriptions = [];
    let isLoading = false;
    let searchText = "";
    let page = 1;
    let perPage = 20;
    let totalItems = 0;
    let statusFilter = "all";

    $: totalPages = Math.ceil(totalItems / perPage);

    onMount(() => {
        loadSubscriptions();
    });

    async function loadSubscriptions() {
        isLoading = true;

        try {
            let filter = "";
            
            if (searchText) {
                filter += `user.name ~ "${searchText}" || user.email ~ "${searchText}"`;
            }
            
            if (statusFilter !== "all") {
                if (filter) filter += " && ";
                filter += `status = "${statusFilter}"`;
            }

            const result = await ApiClient.collection("subscriptions").getList(page, perPage, {
                filter: filter,
                sort: "-created",
                expand: "user",
            });

            subscriptions = result.items;
            totalItems = result.totalItems;
        } catch (err) {
            console.error("Failed to load subscriptions:", err);
            addErrorToast("Failed to load subscription list");
        }

        isLoading = false;
    }

    async function cancelSubscription(subscription) {
        if (!await confirm({
            title: "Cancel Subscription",
            content: `Are you sure you want to cancel this subscription? This action cannot be undone.`,
            confirmText: "Cancel Subscription",
            cancelText: "Back",
        })) {
            return;
        }

        try {
            // 这里应该调用取消订阅的API
            // 例如：await ApiClient.subscriptions.cancel(subscription.id);
            subscription.status = "cancelled";
            addSuccessToast("Subscription successfully cancelled");
        } catch (err) {
            console.error("Failed to cancel subscription:", err);
            addErrorToast("Failed to cancel subscription");
        }
    }

    function handleSearch() {
        page = 1;
        loadSubscriptions();
    }

    function handleStatusChange() {
        page = 1;
        loadSubscriptions();
    }

    function changePage(newPage) {
        if (newPage < 1 || newPage > totalPages || newPage === page) {
            return;
        }
        page = newPage;
        loadSubscriptions();
    }

    function getStatusLabel(status) {
        switch (status) {
            case "active": return "Active";
            case "cancelled": return "Cancelled";
            case "expired": return "Expired";
            default: return status;
        }
    }

    function getStatusClass(status) {
        switch (status) {
            case "active": return "success";
            case "cancelled": return "danger";
            case "expired": return "warning";
            default: return "";
        }
    }
</script>

<PageWrapper>
    <header class="page-header">
        <h1>
            <i class="ri-exchange-dollar-line" aria-hidden="true"></i>
            <span class="txt">Subscription Management</span>
        </h1>
    </header>

    <div class="wrapper">
        <div class="filters">
            <div class="search-bar">
                <form on:submit|preventDefault={handleSearch}>
                    <Field type="text" bind:value={searchText} placeholder="Search members..." />
                    <button type="submit" class="btn btn-secondary">Search</button>
                </form>
            </div>
            
            <div class="status-filter">
                <label for="status-filter">Status Filter:</label>
                <select id="status-filter" bind:value={statusFilter} on:change={handleStatusChange}>
                    <option value="all">All</option>
                    <option value="active">Active</option>
                    <option value="cancelled">Cancelled</option>
                    <option value="expired">Expired</option>
                </select>
            </div>
        </div>

        {#if isLoading}
            <div class="loader"></div>
        {:else if subscriptions.length === 0}
            <div class="alert alert-info">
                <i class="ri-information-line" aria-hidden="true"></i>
                <span class="txt">No subscriptions found.</span>
            </div>
        {:else}
            <div class="table-wrapper">
                <table class="table">
                    <thead>
                        <tr>
                            <th>Member</th>
                            <th>Plan</th>
                            <th>Status</th>
                            <th>Start Date</th>
                            <th>End Date</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each subscriptions as subscription (subscription.id)}
                            <tr>
                                <td>
                                    {#if subscription.expand?.user}
                                        {subscription.expand.user.name || subscription.expand.user.email}
                                    {:else}
                                        --
                                    {/if}
                                </td>
                                <td>{subscription.plan || "--"}</td>
                                <td>
                                    <span class="label {getStatusClass(subscription.status)}">
                                        {getStatusLabel(subscription.status)}
                                    </span>
                                </td>
                                <td>{CommonHelper.formatDate(subscription.startDate)}</td>
                                <td>{subscription.endDate ? CommonHelper.formatDate(subscription.endDate) : "--"}</td>
                                <td>
                                    <div class="actions">
                                        {#if subscription.status === "active"}
                                            <button class="btn btn-sm btn-outline" on:click={() => cancelSubscription(subscription)}>
                                                <i class="ri-close-circle-line" aria-hidden="true"></i>
                                                <span class="txt">Cancel</span>
                                            </button>
                                        {/if}
                                        <a href="/subscriptions/{subscription.id}" class="btn btn-sm btn-outline">
                                            <i class="ri-eye-line" aria-hidden="true"></i>
                                            <span class="txt">View</span>
                                        </a>
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
    .filters {
        display: flex;
        flex-wrap: wrap;
        gap: 20px;
        margin-bottom: 20px;
        align-items: center;
    }
    .search-bar form {
        display: flex;
        gap: 10px;
        max-width: 500px;
    }
    .status-filter {
        display: flex;
        align-items: center;
        gap: 10px;
    }
    .status-filter select {
        padding: 8px;
        border-radius: 4px;
        border: 1px solid #ddd;
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