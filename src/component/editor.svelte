<script>
    import CKEditor from "ckeditor5-svelte";
    import Decoupled from "@ckeditor/ckeditor5-build-decoupled-document/build/ckeditor";

    let data = "Hello, World!";
    let editor = Decoupled;
    let instance = null;

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
    };

    async function GetDirs() {
        const resp = await fetch("doc/dirs");
        return await resp.json();
    }

    async function SendData() {
        const payload = {
            directory: document.getElementById("directory").value,
            subject: document.getElementById("subject").value,
            content: data
        };

        // Need to Code the Rule Here

        const resp = await fetch("doc/send", {
            method: "POST",
            headers: {
                "Content-Type": "application/json; charset=utf-8"
            },
            body: JSON.stringify(payload)
        });

        if (resp.status == 200) {
            alert("Sent")
        } else {
            alert("Failed")
        }
    };
</script>

<div id="editor">
    <div id="control">
        <div id="control-left">
            <div id="control-left-directory" class="container">
                <h1>directory</h1>
                <select name="directory" id="directory">
                    {#await GetDirs() then result}
                        {#each result as dir}
                            <option value={dir}>
                                {dir}
                            </option>
                        {/each}
                    {/await}
                </select>
            </div>
            <div id="control-left-subject" class="container">
                <h1>subject</h1>
                <input type="text" id="subject" name="subject"/>
            </div>
        </div>
        <div id="control-right">
            <input type="submit" value="send"
                on:click={SendData}
            />
        </div>
    </div>
    <CKEditor
        bind:editor
        on:ready={OnReady}
        bind:config={editorConfig}
        bind:value={data}
    />
</div>

<style lang="scss">
    #editor {
        grid-template-rows: 50px 50px 1fr;
        background-color: #F1F1F1;
        grid-template-columns: 1fr;
        word-wrap: break-word;
        position: absolute;
        font-weight: 400;
        color: #3B3B3B;
        font-size: 32px;
        display: grid;
        height: 100%;
        width: 100%;
    }

    #control {
        font-family: "Quicksand", sans-serif;
        justify-content: space-between;
        text-transform: capitalize;
        align-items: center;
        display: flex;
        #control-left {
            display: flex;
            .container {
                align-items: center;
                display: flex;
                h1 {
                    font-size: 16px;
                    padding: 0 25px;
                }
                select, input {
                    font-family: "Quicksand", sans-serif;
                    text-transform: none;
                    cursor: pointer;
                    padding: 0 25px;
                    font-size: 16px;
                    height: 40px;
                    width: 500px;
                }
            }
        }
        #control-right {
            padding: 0 25px;
            display: flex;
            input {
                font-family: "Quicksand", sans-serif;
                text-transform: capitalize;
                cursor: pointer;
                height: 40px;
                width: 250px;
            }
        }
    }
</style>