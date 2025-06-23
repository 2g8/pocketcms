<script>
    import { createEventDispatcher } from "svelte";
    import OverlayPanel from "@/components/base/OverlayPanel.svelte";
    import Field from "@/components/base/Field.svelte";
    
    const dispatch = createEventDispatcher();
    
    export let postData = {
        slug: "",
        publishDate: "",
        tags: "",
        visibility: "public",
        excerpt: "",
        author: ""
    };
    
    let settingPanel;
    
    export function show() {
        return settingPanel?.show();
    }
    
    export function hide() {
        return settingPanel?.hide();
    }
    
    function handleSave() {
        dispatch('save', postData);
        hide();
    }
</script>

<OverlayPanel
    bind:this={settingPanel}
    class="overlay-panel-lg colored-header"
    on:hide
    on:show
>
    <svelte:fragment slot="header">
        <h4 class="upsert-panel-title">Post Settings</h4>
    </svelte:fragment>
    
    <div class="panel-content">
        <div class="grid">
            <div class="col-lg-12">
                <Field class="form-field" name="slug" let:uniqueId>
                    <label for={uniqueId}>Post URL</label>
                    <input 
                        type="text"
                        id={uniqueId}
                        bind:value={postData.slug}
                        placeholder="Post URL identifier"
                    />
                </Field>
            </div>
            
            <div class="col-lg-12">
                <Field class="form-field" name="publishDate" let:uniqueId>
                    <label for={uniqueId}>Publish Date</label>
                    <input 
                        type="datetime-local"
                        id={uniqueId}
                        bind:value={postData.publishDate}
                    />
                </Field>
            </div>
            
            <div class="col-lg-12">
                <Field class="form-field" name="tags" let:uniqueId>
                    <label for={uniqueId}>Tags</label>
                    <input 
                        type="text"
                        id={uniqueId}
                        bind:value={postData.tags}
                        placeholder="Separate multiple tags with commas"
                    />
                </Field>
            </div>
            
            <div class="col-lg-12">
                <Field class="form-field" name="visibility" let:uniqueId>
                    <label for={uniqueId}>Visibility</label>
                    <select id={uniqueId} bind:value={postData.visibility}>
                        <option value="public">Public</option>
                        <option value="members">Members Only</option>
                        <option value="paid">Paid Members</option>
                    </select>
                </Field>
            </div>
            
            <div class="col-lg-12">
                <Field class="form-field" name="excerpt" let:uniqueId>
                    <label for={uniqueId}>Excerpt</label>
                    <textarea 
                        id={uniqueId}
                        bind:value={postData.excerpt}
                        placeholder="Post excerpt"
                        rows="3"
                    ></textarea>
                </Field>
            </div>
            
            <div class="col-lg-12">
                <Field class="form-field" name="created_by" let:uniqueId>
                    <label for={uniqueId}>Author</label>
                    <input 
                        type="text"
                        id={uniqueId}
                        bind:value={postData.created_by}
                        placeholder="Post author"
                    />
                </Field>
            </div>
        </div>
    </div>
    
    <svelte:fragment slot="footer">
        <button type="button" class="btn btn-transparent" on:click={() => hide()}>
            <span class="txt">Cancel</span>
        </button>
        <button type="button" class="btn btn-expanded-sm" on:click={handleSave}>
            <span class="txt">Save Settings</span>
        </button>
    </svelte:fragment>
</OverlayPanel>

<style>
    .panel-content {
        padding: 15px;
    }
</style>