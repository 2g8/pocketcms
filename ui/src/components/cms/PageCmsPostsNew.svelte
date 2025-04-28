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
    
    $pageTitle = "New Post";

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
            slug = CommonHelper.slugify(title);
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
    </header>

    <div class="wrapper">
        <form class="panel" autocomplete="off" on:submit|preventDefault={handleSubmit}>
            <div class="grid">
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
                    <Field class="form-field required" name="slug" let:uniqueId>
                        <label for={uniqueId}>Slug</label>
                        <input 
                            type="text"
                            id={uniqueId}
                            bind:value={slug}
                            placeholder="Post URL identifier"
                            required
                        />
                    </Field>
                </div>

                <div class="col-lg-12">
                    <Field class="form-field" name="custom_excerpt" let:uniqueId>
                        <label for={uniqueId}>Excerpt</label>
                        <textarea 
                            id={uniqueId}
                            bind:value={custom_excerpt}
                            placeholder="Post excerpt"
                            rows="3"
                        ></textarea>
                    </Field>
                </div>

                <div class="col-lg-12">
                    <Field class="form-field required" name="lexical" let:uniqueId>
                        <label for={uniqueId}>Content</label>
                        <textarea 
                            id={uniqueId}
                            bind:value={lexical}
                            placeholder="Enter post content"
                            rows="10"
                            required
                        ></textarea>
                    </Field>
                </div>

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

                <div class="col-lg-6">
                    <Field class="form-field" name="status" let:uniqueId>
                        <label for={uniqueId}>Status</label>
                        <select id={uniqueId} bind:value={status}>
                            <option value="draft">Draft</option>
                            <option value="published">Published</option>
                        </select>
                    </Field>
                </div>

                <div class="col-lg-6">
                    <Field class="form-field" name="visibility" let:uniqueId>
                        <label for={uniqueId}>Visibility</label>
                        <select id={uniqueId} bind:value={visibility}>
                            <option value="public">Public</option>
                            <option value="members">Members only</option>
                            <option value="paid">Paid members</option>
                        </select>
                    </Field>
                </div>

                <div class="col-lg-6">
                    <Field class="form-field form-field-toggle" name="featured" let:uniqueId>
                        <input 
                            type="checkbox"
                            id={uniqueId}
                            bind:checked={featured}
                        />
                        <label for={uniqueId}>
                            <span class="txt">Featured post</span>
                        </label>
                    </Field>
                </div>

                <div class="col-lg-6">
                    <Field class="form-field form-field-toggle" name="show_title_and_feature_image" let:uniqueId>
                        <input 
                            type="checkbox"
                            id={uniqueId}
                            bind:checked={show_title_and_feature_image}
                        />
                        <label for={uniqueId}>
                            <span class="txt">Show title and featured image</span>
                        </label>
                    </Field>
                </div>

                <div class="col-lg-12">
                    <div class="accordions">
                        <details class="accordion">
                            <summary class="accordion-header">Advanced Settings</summary>
                            <div class="accordion-content">
                                <div class="grid">
                                    <div class="col-lg-6">
                                        <Field class="form-field" name="canonical_url" let:uniqueId>
                                            <label for={uniqueId}>规范URL</label>
                                            <input 
                                                type="text"
                                                id={uniqueId}
                                                bind:value={canonical_url}
                                                placeholder="Canonical URL"
                                            />
                                        </Field>
                                    </div>
                                    
                                    <div class="col-lg-12">
                                        <Field class="form-field" name="codeinjection_head" let:uniqueId>
                                            <label for={uniqueId}>Head代码注入</label>
                                            <textarea 
                                                id={uniqueId}
                                                bind:value={codeinjection_head}
                                                placeholder="Code to inject in <head>"
                                                rows="3"
                                            ></textarea>
                                        </Field>
                                    </div>

                                    <div class="col-lg-12">
                                        <Field class="form-field" name="codeinjection_foot" let:uniqueId>
                                            <label for={uniqueId}>Footer代码注入</label>
                                            <textarea 
                                                id={uniqueId}
                                                bind:value={codeinjection_foot}
                                                placeholder="Code to inject before </body>"
                                                rows="3"
                                            ></textarea>
                                        </Field>
                                    </div>
                                </div>
                            </div>
                        </details>
                    </div>
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