<script>
    import { pageTitle, pocketCMSSettings, setPocketCmsSettings } from "@/stores/app";
    import PageWrapper from "@/components/base/PageWrapper.svelte";
    import { onMount } from "svelte";
    import { push, replace } from "svelte-spa-router";
    import { setErrors } from "@/stores/errors";
    import ApiClient from "@/utils/ApiClient";
    import CommonHelper from "@/utils/CommonHelper";

    $pageTitle = "Welcome to PocketCMS";    

    let isLoading = false;
    let isSaving = false;
    let formSettings = {};
    
    loadSettings();

    async function loadSettings() {
        isLoading = true;
        try {
            const settings = (await ApiClient.settings.getAll()) || {};
            init(settings?.pocketcms);
        } catch (err) {
            ApiClient.error(err);
        }

        isLoading = false;
    }

    let steps = [
        {
            title: "Setup Site Info",
            description: "Start a new PocketCMS site",
            icon: "ri-add-line",
            completed: $pocketCMSSettings?.flagOnboardSetting || false,
            action: "/settings/seo"
        },
        {
            title: "Customize Design",
            description: "Personalize your website appearance",
            icon: "ri-palette-line",
            completed: $pocketCMSSettings?.flagOnboardTheme || false,
            action: "/settings/theme"
        },
        {
            title: "Explore Editor",
            description: "Learn how to create content",
            icon: "ri-edit-2-line",
            completed: $pocketCMSSettings?.flagOnboardEditor || false,
            action: "/posts/new"
        },
        {
            title: "Build Audience",
            description: "Start attracting your readers",
            icon: "ri-group-line",
            completed: $pocketCMSSettings?.flagOnboardMember || false,
            action: "/members"
        },
        {
            title: "Launch Your Website",
            description: "Take down your Coming Soon page and start building your community",
            icon: "ri-global-line",
            completed: $pocketCMSSettings?.flagWebsiteOnline || false,
            action: "/set_website_online"
        }
    ];

    function handleStepClick(action, completed) {
        if (completed) return;
        switch (action) {
            case "/settings/seo":
                formSettings.pocketcms.flagOnboardSetting = true;
                break;
            case "/settings/theme":
                formSettings.pocketcms.flagOnboardTheme = true;
                break;
            case "/posts/new":
                formSettings.pocketcms.flagOnboardEditor = true;
                break;
            case "/members":
                formSettings.pocketcms.flagOnboardMember = true;
                break;
            case "/dashboard":
                formSettings.pocketcms.flagOnboardSkip = true;
                break;
            case "/set_website_online":
                formSettings.pocketcms.flagWebsiteOnline = true;
                break;
        }
        save(action);
    }

    async function save(action) {
        if (isSaving) {
            return;
        }

        isSaving = true;

        try {
            const settings = await ApiClient.settings.update(CommonHelper.filterRedactedProps(formSettings));
            init(settings?.pocketcms);

            setErrors({});
            
            push(action);
        } catch (err) {
            ApiClient.error(err);
        }

        isSaving = false;
    }

    function init(settings = {}) {
        setPocketCmsSettings(settings);
        if (settings.flagOnboardSkip == true) {
            replace('/dashboard');
        }
        formSettings = {
            pocketcms: settings
        };
        steps = steps.map((step, index) => {
            let namemap = {
                0: "flagOnboardSetting",
                1: "flagOnboardTheme",
                2: "flagOnboardEditor",
                3: "flagOnboardMember",
                4: "flagWebsiteOnline",
                5: "flagOnboardSkip"
            }
            return {
                ...step,
                completed: $pocketCMSSettings[namemap[index]] || false
            };
        })
    }

</script>
<PageWrapper>
    <div class="wrapper">
        {#if isLoading}
            <div class="loader"></div>
        {:else}
            <div class="onboarding-container">
                <div class="onboarding-header">
                    <h1>Get Started with PocketCMS</h1>
                    <p>Welcome! Let's help you set up your new site.</p>
                </div>

                <div class="steps-container">
                    {#each steps as step, index}
                        <div class="step-card {step.completed ? 'completed' : ''}" on:click={() => handleStepClick(step.action, step.completed)}>
                            <div class="step-number">{index + 1}</div>
                            <div class="step-content">
                                <div class="step-title">{step.title}</div>
                                <div class="step-description">{step.description}</div>
                            </div>
                            <div class="step-icon">
                                <i class="{step.icon}" aria-hidden="true"></i>
                                {#if step.completed}
                                    <div class="step-check">
                                        <i class="ri-check-line" aria-hidden="true"></i>
                                    </div>
                                {/if}
                            </div>
                        </div>
                    {/each}
                </div>

                <div class="help-section">                    
                    {#if ($pocketCMSSettings.flagOnboardSetting && $pocketCMSSettings.flagOnboardTheme && $pocketCMSSettings.flagOnboardEditor && $pocketCMSSettings.flagOnboardMember && $pocketCMSSettings.flagWebsiteOnline) }
                        <h2>Great! You're all set.</h2>
                        <button class="btn btn-block btn-outline"  on:click={() => handleStepClick('/dashboard',false)}>
                            <i class="ri-eye-line" />
                            <span class="txt">Visit Dashboard</span>
                        </button>
                    {:else}
                        <p>More questions? Check our <a href="https://pocketcms.dev/help?utm_source=admin&utm_campaign=onboard">Help Center</a>.</p>
                        <button class="btn btn-block btn-outline"  on:click={() => handleStepClick('/dashboard',false)}>
                            <i class="ri-add-line" />
                            <span class="txt">Skip Onboarding</span>
                        </button>
                    {/if}
                </div>
            </div>
        {/if}
    </div>
</PageWrapper>

<style>
    .onboarding-container {
        max-width: 60%;
        margin: 0 auto;
        padding: 2rem 1rem;
    }

    .onboarding-header {
        text-align: center;
        margin-bottom: 3rem;
    }

    .logo {
        margin: 0 auto 1.5rem;
        width: 60px;
        height: 60px;
    }

    .onboarding-header h1 {
        font-size: 2.5rem;
        margin-bottom: 0.5rem;
    }

    .onboarding-header p {
        font-size: 1.125rem;
        margin: 15px 0;
    }

    .steps-container {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .step-card {
        background: var(--card-bg, #fff);
        border-radius: 0.75rem;
        box-shadow: 0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1);
        padding: 1.5rem;
        display: flex;
        align-items: center;
        transition: all 0.2s ease-in-out;
        cursor: pointer;
        position: relative;
        overflow: hidden;
    }

    .step-card:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
    }

    .step-card.completed {
        border-left: 4px solid var(--success-color, #10b981);
        color: #999999;
        cursor: default;
    }

    .step-card.completed:hover {
        transform: none;
        box-shadow: 0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1);
    }

    .step-number {
        width: 2rem;
        height: 2rem;
        border-radius: 50%;
        background-color: var(--primary-color-light, #e0e7ff);
        display: flex;
        align-items: center;
        justify-content: center;
        font-weight: 600;
        margin-right: 1rem;
    }

    .step-content {
        flex: 1;
    }

    .step-title {
        font-weight: 600;
        font-size: 1.125rem;
        margin-bottom: 0.25rem;
    }

    .step-description {
        font-size: 0.875rem;
    }

    .step-icon {
        position: relative;
        width: 2.5rem;
        height: 2.5rem;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .step-icon i {
        font-size: 1.5rem;
    }

    .step-check {
        position: absolute;
        bottom: 0.25rem;
        right: 0.25rem;
        width: 2rem;
        height: 2rem;
        border-radius: 50%;
        background-color: var(--success-color, #10b981);
        color: white;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .step-check i {
        font-size: 0.875rem;
        color: white;
    }

    .help-section {
        width:50%;
        margin: 0 auto;
        text-align: center;
    }

    .help-section p{
        margin: 15px 0;
    }

    .help-section h2{
        font-size: 2.2em;
        margin: 30px 0;
    }

    .help-section a:hover {
        text-decoration: none;
    }

    .skip-button {
        margin-top: 1rem;
        background: none;
        border: none;
        color: var(--text-color-secondary, #4b5563);
        cursor: pointer;
        font-size: 0.875rem;
        text-decoration: underline;
    }

    .skip-button:hover {
        color: var(--text-color, #1f2937);
    }

    .loader {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 200px;
    }

    @media (max-width: 640px) {
        .onboarding-header h1 {
            font-size: 1.5rem;
        }

        .onboarding-header p {
            font-size: 1rem;
        }

        .step-card {
            padding: 1rem;
        }

        .step-title {
            font-size: 1rem;
        }
    }
</style>
