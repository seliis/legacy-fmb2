<script>
    import TreeMaker from "../component/treemaker.svelte";
    import Editor from "../component/editor.svelte";
    import { slide } from "svelte/transition";

    export let params = {};

    const controlMenu = [
        "create"
    ]

    let slave = false;
    function SlaveControl() {
        slave = !slave;
    }

    function MasterControl(target) {
        window.location.href = "#/document/" + target;
        if (slave == true) {
            slave = !slave;
        }
    }

    async function GetDocs() {
        const recv = await fetch("docs");
        return await recv.json();
    }

    async function ChangeContent(name) {
        console.log(name);
        slave = !slave;
    }
</script>

<section>
    <nav>
        <div id="master">
            <div id="tree" on:click={SlaveControl}>
                {#if slave == false}
                    <i class="fas fa-folder"/>
                {:else}
                    <i class="fas fa-folder-open"/>
                {/if}
                <h1>document browser</h1>
            </div>
            <div id="control">
                {#each controlMenu as menuName}
                    <div on:click={
                        () => MasterControl(menuName)
                    }>
                        {menuName}
                    </div>
                {/each}
            </div>
        </div>
        {#if slave == true}
            <div id="slave" transition:slide>
                {#await GetDocs() then result}
                    {#each result as data}
                        <TreeMaker
                            inherit={data}
                            name={data.name}
                            control={!slave}
                            changer={ChangeContent}
                        />
                    {/each}
                {/await}
            </div>
        {/if}
    </nav>
    <div id="content">
        {#if params.target == "create"}
            <Editor/>
        {/if}
    </div>
</section>

<style lang="scss">
    section {
        grid-template-rows: 50px 1fr;
        grid-template-columns: 1fr;
        display: grid;
    }

    nav {
        font-family: "Quicksand", sans-serif;
        background-color: $miho-blue-a;
        color: $miho-white-a;
        z-index: 99;
        #master {
            justify-content: space-between;
            align-items: center;
            padding: 0 25px;
            display: flex;
            height: 100%;
            #tree {
                grid-template-columns: 30px 1fr;
                grid-template-rows: 1fr;
                cursor: pointer;
                display: grid;
                i {
                    text-align: left;
                    font-size: 20px;
                }
                h1 {
                    text-transform: capitalize;
                    font-weight: 700;
                    font-size: 16px;
                }
                &:hover {
                    color: $miho-yellow-a;
                }
            }
            #control {
                display: flex;
                div {
                    background-color: $miho-blue-c;
                    text-transform: capitalize;
                    border-radius: 5px;
                    padding: 8px 16px;
                    font-weight: 500;
                    font-size: 16px;
                    cursor: pointer;
                    &:hover {
                        background-color: $miho-yellow-a;
                        color: $miho-black-a;
                    }
                }
            }
        }
        #slave {
            background-color: $miho-blue-a;
            height: calc(100% - 50px);
            flex-direction: column;
            position: absolute;
            padding: 0 25px;
            display: flex;
            width: 300px;
        }
    }

    #content {
        position: relative;
        height: 100%;
        width: 100%;
    }
</style>