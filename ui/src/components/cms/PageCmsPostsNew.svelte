<script>
    import { onDestroy, onMount } from "svelte";
    import { push } from "svelte-spa-router";
    import { addSuccessToast, addErrorToast } from "@/stores/toasts";
    import { pageTitle } from "@/stores/app";
    import { superuser } from "@/stores/superuser";
    import { confirm } from "@/stores/confirmation";
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import Field from "@/components/base/Field.svelte";
    import FileField from "@/components/records/fields/FileField.svelte";
    import EditorField from "@/components/records/fields/EditorField.svelte";
    import MenuPosts from "@/components/cms/MenuPosts.svelte";
    import PostSettingPanel from "@/components/cms/PostSettingPanel.svelte";
    import { loadCollections } from "@/stores/collections";
    
    LoadContentforfilepicker();

    async function LoadContentforfilepicker() {
        await loadCollections();
    }
    
    $pageTitle = "New Post";
    onMount(() => {
        // Set CSS variable, z-index is 120 for current page, 0 for other pages
        document.documentElement.style.setProperty('--page-wrapper-z-index', '120');
        requestAnimationFrame(() => {
            // Delay to next frame to get width, layout calculation may be completed
            let pageSidebar = document.querySelector('.page-sidebar');
            let sidebarWidth = pageSidebar.offsetWidth;
            document.documentElement.style.setProperty('--pageSidebarWidth', `${sidebarWidth}px`);
        });
    });
    onDestroy(() => {
        // When page is destroyed, restore CSS variables
        document.documentElement.style.setProperty('--page-wrapper-z-index', '0');
        document.documentElement.style.setProperty('--pageSidebarWidth', `235px`);
    })
    
    let settingPanel;

    let title = "";
    let slug = "";
    let lexical = ""; // Use lexical instead of content
    let feature_image = "";
    let featured = false;
    let type = "post";
    let status = "draft"; // draft, published
    let visibility = "public";
    let custom_excerpt = "";
    let codeinjection_head = "";
    let codeinjection_foot = "";
    let canonical_url = "";
    let published_at = Date.now();
    let tags = "";
    let created_by = "";
    let member_id = "";
    let email_recipient_filter = "all";
    let show_title_and_feature_image = true;
    let isSubmitting = false;
    
    // File upload related variables
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
        
        // If no slug, generate from title
        if (!slug) {
            generateSlug();
        }

        // Get current user information
        const email = $superuser.email;
        try {
            // 通过ApiClient读取members collection，检查是否存在email为当前superuser的记录
            const membersResponse = await ApiClient.collection('members').getList(1, 1, {
                filter: `email="${email}"`
            });
            
            // 如果不存在email为当前superuser的记录，则提示是否添加
            const userExist = membersResponse.items.length > 0 ? membersResponse.items[0] : null;
            
            // Check if user exists
            if (!userExist) {
                confirm("Add yourself as a member to publishing. Do it now?", addCurrentSuperuserToMember);
                return;
            }else{
                // If user exists, set created_by and member_id
                created_by = membersResponse.items[0].name;
                member_id = membersResponse.items[0].id;
            }
        } catch (error) {
            console.error("Failed to check member:", error);
            addErrorToast(error.message || "Failed to check member, please try again");
            return;
        }
        
        isSubmitting = true;
        try {
            // Create FormData object to handle file uploads
            const formData = new FormData();
            
            // Add basic fields
            formData.append("title", title);
            formData.append("slug", slug);
            formData.append("lexical", lexical);
            formData.append("feature_image", feature_image);
            formData.append("featured", featured);
            formData.append("type", type);
            formData.append("status", status);
            formData.append("published_at", published_at);
            formData.append("tags", tags);
            formData.append("created_by", created_by);
            formData.append("member_id", member_id);
            formData.append("email_recipient_filter", email_recipient_filter);
            formData.append("visibility", visibility);
            formData.append("custom_excerpt", custom_excerpt);
            formData.append("codeinjection_head", codeinjection_head);
            formData.append("codeinjection_foot", codeinjection_foot);
            formData.append("canonical_url", canonical_url);
            formData.append("show_title_and_feature_image", show_title_and_feature_image);
            
            // Add uploaded files
            for (const file of uploadedFiles) {
                formData.append("feature_image+", file);
            }
            
            // Process deleted files
            for (const name of deletedFileNames) {
                formData.append("feature_image-", name);
            }
            
            await ApiClient.collection("posts").create(formData);

            addSuccessToast("Post created successfully");
            push("/posts");
        } catch (error) {
            console.error("Failed to create post:", error);
            addErrorToast(error.message || "Failed to save, please try again");
        } finally {
            isSubmitting = false;
        }
    }

    // Add self to the member
    async function addCurrentSuperuserToMember(){
        if($superuser?.email === undefined){
            return;
        }
        const password = CommonHelper.randomString(13);
        const name = ($superuser.email).split("@")[0].trim();
        const username = "admin_"+name.toLowerCase();
        const data = {
            "email": $superuser.email,
            "name": name,
            "username": username,
            "slug": CommonHelper.slugify(username),
            "password": password,
            "passwordConfirm": password,
            "created_by": "admin",
            "email_disabled": false,
            "verified": true,
        }
        try {
            const record = await ApiClient.collection('members').create(data);
            handleSubmit();
        } catch (err) {
            console.error("Failed to add Current Superuser To Member:", err);
            addErrorToast("Failed to add Current Superuser To Member");
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
                        editorConfig={{
                            min_height: 480,
                            height: 480,
                            placeholder: "Press '/' for commands, or '++' for AI autocomplete...",
                            sparkConfig: {                      
                                appid: "9e868ce4",
                                apiKey: "a15d512f688abe35982932ce2f294121",
                                apiSecret: "N2Y5OTgwMDM4ODJkNWUxZjgwZWE1MzFj",
                                domain: '4.0Ultra', // Model version
                                temperature: 0.7,
                                maxTokens: 4096,
                            }
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
        published_at, 
        tags:"", 
        visibility, 
        excerpt: custom_excerpt, 
        created_by
    }} 
    on:save={(e) => {
        slug = e.detail.slug;
        tags = e.detail.tags;
        published_at = e.detail.published_at;
        visibility = e.detail.visibility;
        custom_excerpt = e.detail.excerpt;
    }}
/>

<style>
    :global(.page-wrapper) {        
        z-index: var(--page-wrapper-z-index);
    }
    :global(.tox.tox-tinymce.tox-fullscreen){
        margin-left: calc(var(--appSidebarWidth) + 1px);
        width: calc(100% - var(--appSidebarWidth) - 1px) !important;
        /* margin-left: calc(var(--appSidebarWidth) + var(--pageSidebarWidth) + 1px); 
        */
    }
    /*
    :global(.form-field.form-field-editor>label) {
        display: none !important;
    }
    :global(.form-field.form-field-editor>label~.tinymce-wrapper) {
        padding:2px!important;
    }
    */
</style>