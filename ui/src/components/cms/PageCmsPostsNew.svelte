<script>
    import { onDestroy, onMount } from "svelte";
    import { push } from "svelte-spa-router";
    import { addSuccessToast, addErrorToast } from "@/stores/toasts";
    import { pageTitle } from "@/stores/app";
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import Field from "@/components/base/Field.svelte";
    import FileField from "@/components/records/fields/FileField.svelte";
    import EditorField from "@/components/records/fields/EditorField.svelte";
    import MenuPosts from "@/components/cms/MenuPosts.svelte";
    import PostSettingPanel from "@/components/cms/PostSettingPanel.svelte";
    
    $pageTitle = "New Post";
    onMount(() => {
        // 设置CSS变量，在当前页面z-index为120，其他页面为0
        document.documentElement.style.setProperty('--page-wrapper-z-index', '120');
        requestAnimationFrame(() => {
            // 延迟到下一帧获取宽度，布局计算可能已完成
            let pageSidebar = document.querySelector('.page-sidebar');
            let sidebarWidth = pageSidebar.offsetWidth;
            document.documentElement.style.setProperty('--pageSidebarWidth', `${sidebarWidth}px`);
        });
    });
    onDestroy(() => {
        // 页面销毁时，恢复CSS变量
        document.documentElement.style.setProperty('--page-wrapper-z-index', '0');
        document.documentElement.style.setProperty('--pageSidebarWidth', `235px`);
    })
    
    let settingPanel;

    let title = "";
    let slug = "";
    let lexical = ""; // 使用lexical替代content
    let feature_image = "";
    let featured = false;
    let type = "post";
    let status = "draft"; // draft, published
    let visibility = "public";
    let custom_excerpt = "";
    let codeinjection_head = "";
    let codeinjection_foot = "";
    let canonical_url = "";
    let show_title_and_feature_image = true;
    let isSubmitting = false;
    
    // 文件上传相关变量
    let uploadedFiles = [];
    let deletedFileNames = [];
    let record = { id: "new_post", collectionId: "posts" };
    
    function generateSlug() {
        if (title && !slug) {
            slug = CommonHelper.slugify(title.toLowerCase(),'-');
        }
    }

    async function handleSubmit() {
        if (!title || !lexical) {
            addErrorToast("Please fill in title and content");
            return;
        }
        
        // 如果没有slug，根据标题生成
        if (!slug) {
            generateSlug();
        }

        isSubmitting = true;
        try {
            // 创建FormData对象来处理文件上传
            const formData = new FormData();
            
            // 添加基本字段
            formData.append("title", title);
            formData.append("slug", slug);
            formData.append("lexical", lexical);
            formData.append("feature_image", feature_image);
            formData.append("featured", featured);
            formData.append("type", type);
            formData.append("status", status);
            formData.append("visibility", visibility);
            formData.append("custom_excerpt", custom_excerpt);
            formData.append("codeinjection_head", codeinjection_head);
            formData.append("codeinjection_foot", codeinjection_foot);
            formData.append("canonical_url", canonical_url);
            formData.append("show_title_and_feature_image", show_title_and_feature_image);
            
            // 添加上传的文件
            for (const file of uploadedFiles) {
                formData.append("feature_image+", file);
            }
            
            // 处理删除的文件
            for (const name of deletedFileNames) {
                formData.append("feature_image-", name);
            }
            
            await ApiClient.collection("posts").create(formData);

            addSuccessToast("Post created successfully");
            push("/cms/posts");
        } catch (error) {
            console.error("Failed to create post:", error);
            addErrorToast(error.message || "Failed to save, please try again");
        } finally {
            isSubmitting = false;
        }
    }
</script>

<MenuPosts />

<PageWrapper>
    <header class="page-header">
        <nav class="breadcrumbs">
            <div class="breadcrumb-item">Posts</div>
            <div class="breadcrumb-item">New Post</div>
        </nav>
        <div class="flex-fill"></div>
        <button 
            type="button" 
            class="btn btn-circle btn-transparent" 
            on:click={() => settingPanel.show()}
            aria-label="Post settings"
        >
            <i class="ri-layout-right-2-line" aria-hidden="true"></i>
        </button>
    </header>

    <div class="wrapper">
        <form class="panel" autocomplete="off" on:submit|preventDefault={handleSubmit}>
            <div class="grid">
                <div class="col-lg-12">
                    <FileField 
                        record={record}
                        field={{ 
                            name: "Feature Image",
                            required: false,
                            maxSelect: 1,
                            mimeTypes: ["image/jpeg", "image/png", "image/gif", "image/webp"]
                        }}
                        bind:value={feature_image}
                        bind:uploadedFiles={uploadedFiles}
                        bind:deletedFileNames={deletedFileNames}
                    />
                </div>
                <div class="col-lg-12">
                    <Field class="form-field required" name="title" let:uniqueId>
                        <label for={uniqueId}>Title</label>
                        <input 
                            type="text"
                            id={uniqueId}
                            bind:value={title}
                            placeholder="Enter post title"
                            required
                            on:blur={generateSlug}
                        />
                    </Field>
                </div>

                <div class="col-lg-12">
                    <EditorField
                        field={{
                            name: "Content",
                            required: true,
                            convertURLs: false
                        }}
                        bind:value={lexical}
                    />
                </div>
            </div>

            <div class="flex m-t-base">
                <div class="flex-fill"></div>
                <button
                    type="submit"
                    class="btn btn-expanded"
                    class:btn-loading={isSubmitting}
                    disabled={isSubmitting}
                >
                    <span class="txt">{isSubmitting ? "Saving..." : "Save"}</span>
                </button>
            </div>
        </form>
    </div>
</PageWrapper>

<PostSettingPanel 
    bind:this={settingPanel} 
    postData={{ 
        slug, 
        publishDate: "", 
        tags: "", 
        visibility, 
        excerpt: custom_excerpt, 
        author: "" 
    }} 
    on:save={(e) => {
        slug = e.detail.slug;
        visibility = e.detail.visibility;
        custom_excerpt = e.detail.excerpt;
    }}
/>

<style>
    :global(.page-wrapper) {        
        z-index: var(--page-wrapper-z-index);
    }
    :global(.tox.tox-tinymce.tox-fullscreen){
        margin-left: calc(var(--appSidebarWidth) + var(--pageSidebarWidth) + 1px);
    }
</style>