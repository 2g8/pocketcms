<script>
    import PageSidebar from "@/components/base/PageSidebar.svelte";
    import { hideControls } from "@/stores/app";
    import { link } from "svelte-spa-router";
    import active from "svelte-spa-router/active";
</script>

<PageSidebar class="posts-sidebar">
    <div class="sidebar-content">
        <div class="sidebar-title">Posts</div>
        <a href="/posts" class="sidebar-list-item" use:active={{ path: "/posts" }} use:link>
            <i class="ri-article-line" aria-hidden="true" />
            <span class="txt">All Posts</span>
        </a>
        <a
            href="/posts/published"
            class="sidebar-list-item"
            use:active={{ path: "/posts/published/?.*" }}
            use:link
        >
            <i class="ri-check-line" aria-hidden="true" />
            <span class="txt">Published</span>
        </a>
        <a
            href="/posts/scheduled"
            class="sidebar-list-item"
            use:active={{ path: "/posts/scheduled/?.*" }}
            use:link
        >
            <i class="ri-time-line" aria-hidden="true" />
            <span class="txt">Scheduled</span>
        </a>
        <a
            href="/posts/free"
            class="sidebar-list-item"
            use:active={{ path: "/posts/free/?.*" }}
            use:link
        >
            <i class="ri-price-tag-3-line" aria-hidden="true" />
            <span class="txt">Free Posts</span>
        </a>
        <a
            href="/posts/paid"
            class="sidebar-list-item"
            use:active={{ path: "/posts/paid/?.*" }}
            use:link
        >
            <i class="ri-money-dollar-circle-line" aria-hidden="true" />
            <span class="txt">Paid Posts</span>
        </a>
    </div>

    <footer class="sidebar-footer">
        <button type="button" class="btn btn-block btn-outline">
            <i class="ri-add-line" />
            <span class="txt">New post</span>
        </button>
    </footer>
</PageSidebar>