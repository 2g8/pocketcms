<script>
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";
    import tooltip from "@/actions/tooltip";
    import { addSuccessToast } from "@/stores/toasts";
    import { pageTitle } from "@/stores/app";
    import { setErrors } from "@/stores/errors";
    import Field from "@/components/base/Field.svelte";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import SettingsSidebar from "@/components/settings/SettingsSidebar.svelte";

    $pageTitle = "SEO settings";

    let originalFormSettings = {};
    let formSettings = {};
    let isLoading = false;
    let isSaving = false;
    let initialHash = "";

    $: initialHash = JSON.stringify(originalFormSettings);

    $: hasChanges = initialHash != JSON.stringify(formSettings);

    loadSettings();

    async function loadSettings() {
        isLoading = true;

        try {
            const settings = (await ApiClient.settings.getAll()) || {};
            init(settings);
        } catch (err) {
            ApiClient.error(err);
        }

        isLoading = false;
    }

    async function save() {
        if (isSaving || !hasChanges) {
            return;
        }

        isSaving = true;

        try {
            const settings = await ApiClient.settings.update(CommonHelper.filterRedactedProps(formSettings));
            init(settings);

            setErrors({});

            addSuccessToast("Successfully saved SEO settings.");
        } catch (err) {
            ApiClient.error(err);
        }

        isSaving = false;
    }

    function init(settings = {}) {
        formSettings = {
            seo: settings?.seo || {
                title: "",
                description: "",
                keywords: "",
                author: "",
                robots: "",
                viewport: "",
                canonical: "",
                favicon: "",
                alternate: "",
                og: "",
                twitter: ""
            }
        };

        originalFormSettings = JSON.parse(JSON.stringify(formSettings));
    }

    function reset() {
        formSettings = JSON.parse(JSON.stringify(originalFormSettings || {}));
    }
</script>

<SettingsSidebar />

<PageWrapper>
    <header class="page-header">
        <nav class="breadcrumbs">
            <div class="breadcrumb-item">Settings</div>
            <div class="breadcrumb-item">Seo settings</div>
        </nav>
    </header>

    <div class="wrapper">
        <form class="panel" autocomplete="off" on:submit|preventDefault={save}>
            {#if isLoading}
                <div class="loader" />
            {:else}
                <div class="grid">
                    <div class="col-lg-6">
                        <Field class="form-field required" name="seo.title" let:uniqueId>
                            <label for={uniqueId}>Title</label>
                            <input
                                type="text"
                                id={uniqueId}
                                required
                                bind:value={formSettings.seo.title}
                            />
                            <div class="help-block">
                                The title tag that appears in search engine results.
                            </div>
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field" name="seo.description" let:uniqueId>
                            <label for={uniqueId}>Description</label>
                            <input
                                type="text"
                                id={uniqueId}
                                bind:value={formSettings.seo.description}
                            />
                            <div class="help-block">
                                A brief description of your site for search engine results.
                            </div>
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field" name="seo.keywords" let:uniqueId>
                            <label for={uniqueId}>Keywords</label>
                            <input
                                type="text"
                                id={uniqueId}
                                bind:value={formSettings.seo.keywords}
                            />
                            <div class="help-block">
                                Comma-separated keywords related to your site content.
                            </div>
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field" name="seo.author" let:uniqueId>
                            <label for={uniqueId}>Author</label>
                            <input
                                type="text"
                                id={uniqueId}
                                bind:value={formSettings.seo.author}
                            />
                            <div class="help-block">
                                Author is commonly used to specify the author of a web page.
                            </div>
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field" name="seo.robots" let:uniqueId>
                            <label for={uniqueId}>Robots</label>
                            <input
                                type="text"
                                id={uniqueId}
                                bind:value={formSettings.seo.robots}
                            />
                            <div class="help-block">
                                Controls how search engines crawl and index your site (e.g., "index, follow").
                            </div>
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field" name="seo.viewport" let:uniqueId>
                            <label for={uniqueId}>Viewport</label>
                            <input
                                type="text"
                                id={uniqueId}
                                bind:value={formSettings.seo.viewport}
                            />
                            <div class="help-block">
                                Defines how your site displays on mobile devices (e.g., "width=device-width, initial-scale=1").
                            </div>
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field" name="seo.canonical" let:uniqueId>
                            <label for={uniqueId}>Canonical URL</label>
                            <input
                                type="text"
                                id={uniqueId}
                                bind:value={formSettings.seo.canonical}
                            />
                            <div class="help-block">
                                The preferred URL for your site to avoid duplicate content issues.
                            </div>
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field" name="seo.favicon" let:uniqueId>
                            <label for={uniqueId}>Favicon</label>
                            <input
                                type="text"
                                id={uniqueId}
                                bind:value={formSettings.seo.favicon}
                            />
                            <div class="help-block">
                                Path to your site's favicon.
                            </div>
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field" name="seo.alternate" let:uniqueId>
                            <label for={uniqueId}>Alternate</label>                            
                            <textarea
                                id={uniqueId}
                                class="textarea"
                                rows="3"
                                bind:value={formSettings.seo.alternate}
                            ></textarea>
                            <div class="help-block">
                                Alternate versions of your site (e.g., different languages).
                            </div>
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field" name="seo.og" let:uniqueId>
                            <label for={uniqueId}>Open Graph</label>
                            <textarea
                                id={uniqueId}
                                class="textarea"
                                rows="3"
                                bind:value={formSettings.seo.og}
                            ></textarea>
                            <div class="help-block">
                                Open Graph metadata for social media sharing.
                            </div>
                        </Field>
                    </div>

                    <div class="col-lg-6">
                        <Field class="form-field" name="seo.twitter" let:uniqueId>
                            <label for={uniqueId}>Twitter Card</label>
                            <textarea
                                id={uniqueId}
                                class="textarea"
                                rows="3"
                                bind:value={formSettings.seo.twitter}
                            ></textarea>
                            <div class="help-block">
                                Twitter Card metadata for Twitter sharing.
                            </div>
                        </Field>
                    </div>
                </div>

                <div class="flex m-t-base">
                    <div class="flex-fill" />

                    {#if hasChanges}
                        <button
                            type="button"
                            class="btn btn-transparent btn-hint"
                            disabled={isSaving}
                            on:click={() => reset()}
                        >
                            <span class="txt">Cancel</span>
                        </button>
                    {/if}

                    <button
                        type="submit"
                        class="btn btn-expanded"
                        class:btn-loading={isSaving}
                        disabled={!hasChanges || isSaving}
                        on:click={() => save()}
                    >
                        <span class="txt">Save changes</span>
                    </button>
                </div>
            {/if}
        </form>
    </div>
</PageWrapper>