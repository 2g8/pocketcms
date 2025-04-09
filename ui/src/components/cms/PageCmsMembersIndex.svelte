<script>
    import { onMount } from "svelte";
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import { pageTitle } from "@/stores/app";
    import { addSuccessToast, addErrorToast } from "@/stores/toasts";
    import { confirm } from "@/stores/confirmation";
    import Field from "@/components/base/Field.svelte";
    import PageWrapper from "@/components/base/PageWrapper.svelte";

    $pageTitle = "Member Management";

    let members = [];
    let isLoading = false;
    let searchText = "";
    let page = 1;
    let perPage = 20;
    let totalItems = 0;

    $: totalPages = Math.ceil(totalItems / perPage);

    onMount(() => {
        loadMembers();
    });

    async function loadMembers() {
        isLoading = true;

        try {
            const result = await ApiClient.collection("users").getList(page, perPage, {
                filter: searchText ? `name ~ "${searchText}" || email ~ "${searchText}"` : "",
                sort: "-created",
            });

            members = result.items;
            totalItems = result.totalItems;
        } catch (err) {
            console.error("Failed to load members:", err);
            addErrorToast("Failed to load member list");
        }

        isLoading = false;
    }

    async function deleteMember(member) {
        if (!await confirm({
            title: "Delete Member",
            content: `Are you sure you want to delete member "${member.name || member.email}"? This action cannot be undone.`,
            confirmText: "Delete",
            cancelText: "Cancel",
        })) {
            return;
        }

        try {
            await ApiClient.collection("users").delete(member.id);
            members = members.filter(m => m.id !== member.id);
            totalItems--;
            addSuccessToast("Member successfully deleted");
        } catch (err) {
            console.error("Failed to delete member:", err);
            addErrorToast("Failed to delete member");
        }
    }

    function handleSearch() {
        page = 1;
        loadMembers();
    }

    function changePage(newPage) {
        if (newPage < 1 || newPage > totalPages || newPage === page) {
            return;
        }
        page = newPage;
        loadMembers();
    }
</script>

<PageWrapper>
    <header class="page-header">
        <h1>Member Management</h1>
        <div class="flex-fill"></div>
        <form class="search-form" on:submit|preventDefault={handleSearch}>
            <Field type="text" bind:value={searchText} placeholder="Search members..." class="search-input" />
            <button type="submit" class="btn btn-secondary btn-sm">Search</button>
        </form>
    </header>

    <div class="page-content">
        {#if isLoading}
            <div class="loader"></div>
        {:else if members.length === 0}
            <div class="block txt-center">
                <h6>No members found</h6>
                <p class="txt-hint">Try using different search criteria or add new members</p>
            </div>
        {:else}
            <div class="table-wrapper">
                <table class="table">
                    <thead>
                        <tr>
                            <th>Name</th>
                            <th>Email</th>
                            <th>Registration Date</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each members as member (member.id)}
                            <tr>
                                <td>{member.name || "--"}</td>
                                <td>{member.email}</td>
                                <td>{CommonHelper.formatDate(member.created)}</td>
                                <td>
                                    <div class="actions">
                                        <button class="btn btn-sm btn-outline" on:click={() => deleteMember(member)}>Delete</button>
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
    .page-header {
        display: flex;
        align-items: center;
        margin-bottom: 20px;
    }
    .flex-fill {
        flex: 1;
    }
    .search-form {
        display: flex;
        gap: 10px;
    }
    .search-input {
        width: 250px;
    }
    .table-wrapper {
        overflow-x: auto;
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