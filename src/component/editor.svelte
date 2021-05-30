<script>
    import CKEditor from "ckeditor5-svelte";
    import Decoupled from "@ckeditor/ckeditor5-build-decoupled-document/build/ckeditor";

    let editor = Decoupled;
    let instance = null;
    let data = "Hello, World!";

    let editorConfig = {
        toolbar: {
            items: [
                "heading",
                "|",
                "fontFamily",
                "fontSize",
                "bold",
                "italic",
                "underline"
            ]
        }
    };

    function OnReady({detail: editor}) {
        instance = editor;
        editor.ui
        .getEditableElement()
        .parentElement.insertBefore(
            editor.ui.view.toolbar.element,
            editor.ui.getEditableElement()
        );
    }
</script>

<div id="editor">
    <CKEditor
        bind:editor
        on:ready={OnReady}
        bind:config={editorConfig}
        bind:value={data}
    />
</div>

<style lang="scss">
    #editor {
        font-weight: 500;
        font-size: 48px;
    }
</style>