<script>
    import { pageTitle } from "@/stores/app";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import ApiClient from "@/utils/ApiClient";
    import { onMount } from "svelte";
    import { link } from "svelte-spa-router";
    import MenuPosts from "@/components/cms/MenuPosts.svelte";

    $pageTitle = "Posts";

    let posts = [];
    let isLoading = true;
    let totalItems = 0;
    let page = 1;
    let perPage = 20;

    onMount(async () => {
        await loadPosts();
    });

    async function loadPosts() {
        isLoading = true;
        try {
            // 这里可以添加获取文章列表的API调用
            // 例如：const result = await ApiClient.posts.getList(page, perPage);
            // posts = result.items;
            // totalItems = result.totalItems;
            posts = []; // 示例数据，实际应从API获取
            totalItems = 0;
            isLoading = false;
        } catch (err) {
            console.error("Failed to load posts", err);
            isLoading = false;
        }
    }

    function handlePageChange(newPage) {
        page = newPage;
        loadPosts();
    }

    async function deletePost(id) {
        if (!confirm("Are you sure you want to delete this post?")) {
            return;
        }

        try {
            // 这里可以添加删除文章的API调用
            // 例如：await ApiClient.posts.delete(id);
            await loadPosts();
        } catch (err) {
            console.error("Failed to delete post", err);
        }
    }
</script>

<MenuPosts />

<PageWrapper>
    <header class="page-header">
        <header class="page-header">
            <nav class="breadcrumbs">
                <div class="breadcrumb-item">Posts</div>
                <div class="breadcrumb-item">All Posts</div>
            </nav>
        </header>
    </header>

    <div class="wrapper">
        {#if isLoading}
            <div class="loader"></div>
        {:else if posts.length === 0}
            <div class="alert alert-info">
                <i class="ri-information-line" aria-hidden="true"></i>
                <span class="txt">No posts found.</span>
            </div>
        {:else}
            <div class="table-wrapper">
                <table class="table">
                    <thead>
                        <tr>
                            <th>Title</th>
                            <th>Status</th>
                            <th>Created</th>
                            <th>Updated</th>
                            <th class="col-action"></th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each posts as post (post.id)}
                            <tr>
                                <td>
                                    <a href="/posts/{post.id}" use:link>{post.title}</a>
                                </td>
                                <td>
                                    <span class="label {post.status === 'published' ? 'success' : 'warning'}">
                                        {post.status}
                                    </span>
                                </td>
                                <td>
                                    <span class="txt-hint">{new Date(post.created).toLocaleString()}</span>
                                </td>
                                <td>
                                    <span class="txt-hint">{new Date(post.updated).toLocaleString()}</span>
                                </td>
                                <td class="col-action">
                                    <div class="actions-group">
                                        <a href="/posts/{post.id}/edit" class="btn btn-sm btn-outline" use:link>
                                            <i class="ri-pencil-line" aria-hidden="true"></i>
                                        </a>
                                        <button type="button" class="btn btn-sm btn-outline btn-danger" on:click={() => deletePost(post.id)}>
                                            <i class="ri-delete-bin-line" aria-hidden="true"></i>
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>

            <!-- Pagination -->
            {#if totalItems > perPage}
                <div class="pagination">
                    <button 
                        class="btn btn-sm btn-outline" 
                        disabled={page === 1} 
                        on:click={() => handlePageChange(page - 1)}
                    >
                        <i class="ri-arrow-left-s-line" aria-hidden="true"></i>
                    </button>
                    
                    <span class="pagination-info">
                        Page {page} of {Math.ceil(totalItems / perPage)}
                    </span>
                    
                    <button 
                        class="btn btn-sm btn-outline" 
                        disabled={page >= Math.ceil(totalItems / perPage)} 
                        on:click={() => handlePageChange(page + 1)}
                    >
                        <i class="ri-arrow-right-s-line" aria-hidden="true"></i>
                    </button>
                </div>
            {/if}
        {/if}
    </div>
</PageWrapper>

<style>
    .pagination {
        display: flex;
        align-items: center;
        justify-content: center;
        margin-top: 20px;
        gap: 10px;
    }
    .pagination-info {
        margin: 0 10px;
    }
    .label {
        padding: 3px 8px;
        border-radius: 4px;
        font-size: 0.85em;
    }
    .label.success {
        background-color: var(--successColor);
        color: white;
    }
    .label.warning {
        background-color: var(--warningColor);
        color: white;
    }
    .loader {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 200px;
    }
</style>