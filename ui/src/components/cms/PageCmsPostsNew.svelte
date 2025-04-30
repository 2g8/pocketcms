<script>
    import { onMount } from "svelte";
    import { push } from "svelte-spa-router";
    import { addSuccessToast, addErrorToast } from "@/stores/toasts";
    import { pageTitle } from "@/stores/app";
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import Field from "@/components/base/Field.svelte";
    import MenuPosts from "@/components/cms/MenuPosts.svelte";
    import PostSettingPanel from "@/components/cms/PostSettingPanel.svelte";
    import NovelEditor from "../base/NovelEditor.svelte";
    
    $pageTitle = "New Post";
    
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
            await ApiClient.collection("posts").create({
                title,
                slug,
                lexical,
                feature_image,
                featured,
                type,
                status,
                visibility,
                custom_excerpt,
                codeinjection_head,
                codeinjection_foot,
                canonical_url,
                show_title_and_feature_image
            });

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
                    <Field class="form-field" name="feature_image" let:uniqueId>
                        <label for={uniqueId}>Featured Image URL</label>
                        <input 
                            type="text"
                            id={uniqueId}
                            bind:value={feature_image}
                            placeholder="Enter image URL"
                        />
                    </Field>
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
                    <Field class="form-field required" name="lexical" let:uniqueId>
                        <label for={uniqueId}>Content</label>
                        <NovelEditor
                            bind:value={lexical}
                            id={uniqueId}
                            row="22"
                            placeholder="Enter post content"
                        />
                    </Field>
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
    .form-container {
        max-width: 800px;
        margin: 0 auto;
        padding: 20px;
    }

    .form-group {
        margin-bottom: 20px;
    }

    .form-actions {
        margin-top: 30px;
    }

    .accordion {
        margin-top: 30px;
        border: 1px solid #eee;
        border-radius: 4px;
    }

    .accordion summary {
        padding: 12px 15px;
        cursor: pointer;
        font-weight: 500;
        background-color: #f9f9f9;
        border-radius: 4px;
    }

    .accordion-content {
        padding: 15px;
    }
</style>