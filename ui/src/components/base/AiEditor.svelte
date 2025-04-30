<script>
    import { onMount, onDestroy, createEventDispatcher } from "svelte";
    import CommonHelper from "@/utils/CommonHelper";
    import "aieditor/dist/style.css";

    // Component properties
    export let id = "aieditor_" + CommonHelper.randomString(7);
    export let placeholder = "Click to input content...";
    export let disabled = false;
    export let value = "";
    export let cssClass = "aieditor-wrapper";
    export let conf = {};

    // Internal state
    let container;
    let editorRef;
    let lastVal = value;

    const dispatch = createEventDispatcher();

    // Reactive editor content updates
    $: {
        try {
            if (editorRef && lastVal !== value) {
                editorRef.setContent(value);
            }
            if (editorRef && disabled !== undefined) {
                // Set editor readonly state
                editorRef.setReadOnly(disabled);
            }
        } catch (err) {
            console.warn("AiEditor reactive error:", err);
        }
    }

    // Initialize editor
    function initEditor() {
        if (!container) return;

        import("aieditor")
            .then(({ AiEditor }) => {
                const finalConfig = {
                    element: `#${id}`,
                    lang: "en",
                    fontFamily: {
                        values: [
                            {
                                name: "Arial",
                                value: "Arial",
                            },
                            {
                                name: "Times New Roman",
                                value: "Times New Roman",
                            },
                            {
                                name: "Verdana",
                                value: "Verdana",
                            },
                            {
                                name: "Tahoma",
                                value: "Tahoma",
                            },
                            {
                                name: "Georgia",
                                value: "Georgia",
                            },
                            {
                                name: "Helvetica",
                                value: "Helvetica",
                            },
                            {
                                name: "Noto Sans",
                                value: "Noto Sans",
                            },
                            {
                                name: "DejaVu Sans",
                                value: "DejaVu Sans",
                            },
                            {
                                name: "Liberation Sans",
                                value: "Liberation Sans",
                            },
                            {
                                name: "Ubuntu",
                                value: "Ubuntu",
                            },
                            {
                                name: "San Francisco",
                                value: "San Francisco",
                            },
                            {
                                name: "宋体",
                                value: "SimSun",
                            },
                            {
                                name: "黑体",
                                value: "SimHei",
                            },
                            {
                                name: "微软雅黑",
                                value: "Microsoft YaHei",
                            },
                            {
                                name: "仿宋",
                                value: "FangSong",
                            },
                            {
                                name: "楷体",
                                value: "KaiTi",
                            },
                            {
                                name: "苹方",
                                value: "PingFang SC",
                            },
                            {
                                name: "华文黑体",
                                value: "STHeiti",
                            },
                            {
                                name: "华文宋体",
                                value: "STSong",
                            },
                            {
                                name: "文泉驿正黑",
                                value: "WenQuanYi Zen Hei",
                            },
                        ],
                    },
                    ai:{
                        models:{
                            spark:{
                                appId: "9e868ce4",
                                apiKey: "a15d512f688abe35982932ce2f294121",
                                apiSecret: "N2Y5OTgwMDM4ODJkNWUxZjgwZWE1MzFj",
                                version: "v4.0"
                            }
                        }
                    },
                    placeholder: placeholder,
                    content: value,
                    readOnly: disabled,
                    ...conf,
                };

                editorRef = new AiEditor(finalConfig);

                // Listen for content changes
                editorRef.on("change", () => {
                    lastVal = editorRef.getContent();
                    if (lastVal !== value) {
                        value = lastVal;
                        dispatch("change", { value });
                    }
                });

                // Dispatch editor initialized event
                dispatch("init", { editor: editorRef });
            })
            .catch((err) => {
                console.error("Failed to initialize AiEditor:", err);
                dispatch("error", { error: err });
            });
    }

    onMount(() => {
        initEditor();
    });

    onDestroy(() => {
        try {
            if (editorRef) {
                editorRef.destroy();
                editorRef = null;
            }
        } catch (err) {
            console.warn("AiEditor destroy error:", err);
        }
    });
</script>

<div bind:this={container} class={cssClass}>
    <div {id}></div>
</div>

<style>
    :global(.aieditor-wrapper) {
        width: 100%;
        border: 1px solid #ddd;
        border-radius: 4px;
        overflow: hidden;
    }
</style>
