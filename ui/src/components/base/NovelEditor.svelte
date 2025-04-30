<script>
    import { onMount, onDestroy, createEventDispatcher } from "svelte";
    import CommonHelper from "@/utils/CommonHelper";
    import { Editor } from "novel-svelte";

    // 组件属性
    export let id = "noveleditor_" + CommonHelper.randomString(7);
    export let placeholder = "Click to input content...";
    export let disabled = false;
    export let value = "";
    export let cssClass = "noveleditor-wrapper";
    export let conf = {};

    // 内部状态
    let container;
    let editorRef;
    let lastVal = value;
    let editorElement;

    const dispatch = createEventDispatcher();

    // 初始化编辑器内容
    function prepareInitialContent() {
        try {
            // 尝试解析为JSON，如果失败则作为HTML字符串处理
            return value ? JSON.parse(value) : null;
        } catch (e) {
            return value || "";
        }
    }

    // 响应式更新编辑器内容和状态
    $: {
        try {
            if (editorRef && lastVal !== value) {
                // 更新编辑器内容
                if (typeof editorRef.commands?.setContent === 'function') {
                    editorRef.commands.setContent(value);
                    lastVal = value;
                }
            }
        } catch (err) {
            console.warn("NovelEditor reactive error:", err);
        }
    }

    // 编辑器内容更新处理函数
    function handleEditorUpdate(editor) {
        try {
            if (editor) {
                const newValue = editor.getHTML();
                if (newValue !== lastVal) {
                    lastVal = newValue;
                    value = newValue;
                    dispatch("change", { value });
                }
            }
        } catch (err) {
            console.warn("NovelEditor update error:", err);
        }
    }

    // 编辑器初始化完成处理函数
    function handleEditorInit(editor) {
        editorRef = editor;
        dispatch("init", { editor: editorRef });

        // 设置编辑器只读状态
        if (disabled) {
            try {
                if (typeof editorRef.setEditable === 'function') {
                    editorRef.setEditable(!disabled);
                }
            } catch (err) {
                console.warn("NovelEditor setEditable error:", err);
            }
        }
    }

    onMount(() => {
        // 在组件挂载后，编辑器会自动初始化
    });

    onDestroy(() => {
        try {
            if (editorRef) {
                // 尝试销毁编辑器实例
                if (typeof editorRef.destroy === 'function') {
                    editorRef.destroy();
                }
                editorRef = null;
            }
        } catch (err) {
            console.warn("NovelEditor destroy error:", err);
        }
    });
</script>

<div bind:this={container} class={cssClass}>
    <Editor
        bind:this={editorRef}
        class="{id}"
        defaultValue={prepareInitialContent()}
        onDebouncedUpdate={handleEditorUpdate}
        onUpdate={(e) => handleEditorUpdate(e)}
        bind:element={editorElement}
        {...conf}
    />
</div>

<style>
    :global(.noveleditor-wrapper) {
        width: 100%;
        border: 1px solid #ddd;
        border-radius: 4px;
        overflow: hidden;
        min-height: 440px; /* 约22行文本的高度 */
    }
</style>