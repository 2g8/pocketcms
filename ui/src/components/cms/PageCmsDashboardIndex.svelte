<script>
    import { pageTitle } from "@/stores/app";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import ApiClient from "@/utils/ApiClient";
    import { onMount } from "svelte";

    $pageTitle = "Dashboard";

    let stats = {
        posts: 0,
        members: 0,
        newsletters: 0,
        subscriptions: 0
    };

    let isLoading = true;

    onMount(async () => {
        try {
            // 这里可以添加获取统计数据的API调用
            // 例如：const result = await ApiClient.stats.getAll();
            // stats = result.data;
            isLoading = false;
        } catch (err) {
            console.error("Failed to load dashboard stats", err);
            isLoading = false;
        }
    });
</script>

<PageWrapper>
    <header class="page-header">
        <h1>
            <i class="ri-dashboard-3-line" aria-hidden="true"></i>
            <span class="txt">Dashboard</span>
        </h1>
    </header>

    <div class="wrapper">
        {#if isLoading}
            <div class="loader"></div>
        {:else}
            <div class="grid">
                <div class="col-sm-6 col-lg-3">
                    <div class="card stats-card">
                        <div class="card-header">
                            <i class="ri-article-line" aria-hidden="true"></i>
                            <span>Posts</span>
                        </div>
                        <div class="card-body">
                            <h2>{stats.posts}</h2>
                        </div>
                    </div>
                </div>
                <div class="col-sm-6 col-lg-3">
                    <div class="card stats-card">
                        <div class="card-header">
                            <i class="ri-group-line" aria-hidden="true"></i>
                            <span>Members</span>
                        </div>
                        <div class="card-body">
                            <h2>{stats.members}</h2>
                        </div>
                    </div>
                </div>
                <div class="col-sm-6 col-lg-3">
                    <div class="card stats-card">
                        <div class="card-header">
                            <i class="ri-mail-send-line" aria-hidden="true"></i>
                            <span>Newsletters</span>
                        </div>
                        <div class="card-body">
                            <h2>{stats.newsletters}</h2>
                        </div>
                    </div>
                </div>
                <div class="col-sm-6 col-lg-3">
                    <div class="card stats-card">
                        <div class="card-header">
                            <i class="ri-exchange-dollar-line" aria-hidden="true"></i>
                            <span>Subscriptions</span>
                        </div>
                        <div class="card-body">
                            <h2>{stats.subscriptions}</h2>
                        </div>
                    </div>
                </div>
            </div>
        {/if}
    </div>
</PageWrapper>

<style>
    .stats-card {
        margin-bottom: 20px;
    }
    .stats-card .card-header {
        display: flex;
        align-items: center;
        gap: 10px;
        font-weight: bold;
    }
    .stats-card .card-body h2 {
        font-size: 2.5rem;
        margin: 0;
        text-align: center;
    }
    .loader {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 200px;
    }
</style>