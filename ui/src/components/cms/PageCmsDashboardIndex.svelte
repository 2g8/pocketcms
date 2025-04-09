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
        <header class="page-header">
            <nav class="breadcrumbs">
                <div class="breadcrumb-item">PocketCMS</div>
                <div class="breadcrumb-item">Dashboard</div>
            </nav>
        </header>
    </header>

    <div class="wrapper">
        {#if isLoading}
            <div class="loader"></div>
        {:else}
            <div class="grid">
                <div class="card-width">
                    <div class="card stats-card">
                        <div class="card-body">
                            <div class="card-content">
                                <div class="card-title">Posts</div>
                                <div class="card-value">5{stats.posts}</div>
                                <div class="card-change positive"><i class="ri-arrow-up-line"></i>+12.5%<span style="color:#999">since last week</span></div>
                            </div>
                            <div class="card-icon">
                                <i class="ri-article-line" aria-hidden="true"></i>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="card-width">
                    <div class="card stats-card">
                        <div class="card-body">
                            <div class="card-content">
                                <div class="card-title">Members</div>
                                <div class="card-value">1111{stats.members}</div>
                                <div class="card-change positive"><i class="ri-arrow-up-line"></i>+12.5%<span style="color:#999">since last week</span></div>
                            </div>
                            <div class="card-icon">
                                <i class="ri-group-line" aria-hidden="true"></i>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="card-width">
                    <div class="card stats-card">
                        <div class="card-body">
                            <div class="card-content">
                                <div class="card-title">Newsletters</div>
                                <div class="card-value">22222{stats.newsletters}</div>
                                <div class="card-change negative"><i class="ri-arrow-down-line"></i>-2.4%<span style="color:#999">since last week</span></div>
                            </div>
                            <div class="card-icon">
                                <i class="ri-mail-send-line" aria-hidden="true"></i>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="card-width">
                    <div class="card stats-card">
                        <div class="card-body">
                            <div class="card-content">
                                <div class="card-title">Subscriptions</div>
                                <div class="card-value">32323{stats.subscriptions}</div>
                                <div class="card-change positive"><i class="ri-arrow-up-line"></i>+15.7%<span style="color:#999">since last week</span></div>
                            </div>
                            <div class="card-icon">
                                <i class="ri-exchange-dollar-line" aria-hidden="true"></i>
                            </div>
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
        background: var(--card-bg, #fff);
        border-radius: 0.75rem;
        box-shadow: 0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1);
        transition: all 0.2s ease-in-out;
        overflow: hidden;
    }

    .stats-card:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
    }

    i:before {
        margin-top: -1px;
    }

    .stats-card .card-body {
        padding: 1.5rem;
        display: flex;
        justify-content: space-between;
        align-items: flex-start;
    }

    .stats-card .card-content {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }

    .stats-card .card-title {
        font-weight: 600;
        font-size: 1rem;
        color: var(--text-color, #1f2937);
    }

    .stats-card .card-value {
        font-size: 1.75rem;
        font-weight: 700;
        color: var(--heading-color, #111827);
        line-height: 1.2;
    }

    .stats-card .card-change {
        font-size: 0.875rem;
        font-weight: 500;
        display: flex;
        align-items: center;
        gap: 0.25rem;
    }

    .stats-card .card-change.positive {
        color: var(--success-color, #10b981);
    }

    .stats-card .card-change.negative {
        color: var(--danger-color, #ef4444);
    }

    .stats-card .card-icon {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 2rem;
        height: 2rem;
        border-radius: 0.5rem;
    }

    .stats-card .card-icon i {
        font-size: 1.5rem;
        color: var(--primary-color);
    }

    .loader {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 200px;
    }

    .grid {
        display: grid;
        gap: 1.5rem;
        width: 100%;
    }

    @media (min-width: 640px) {
        .card-width {
            grid-column: span 6 / span 6;
        }
        .grid {
            gap: 0rem;
            grid-template-columns: repeat(12, minmax(0, 1fr));
        }
    }

    @media (min-width: 1024px) {
        .card-width {
            grid-column: span 3 / span 3;
        }
        .grid {
            gap: 0rem;
            grid-template-columns: repeat(12, minmax(0, 1fr));
        }
    }
</style>