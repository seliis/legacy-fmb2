<script>
    export let control = false;
    export let changer;
    export let inherit;
    export let name;

    function Controller() {
        control = !control;
    }

    function RemoveExt(str) {
        return str.split(
                "."
            ).slice(
                0, -1
            ).join(
                "."
            )
    }
</script>

{#if inherit.dir == true}
    <span class="folder" on:click={Controller}>
        {#if control == false}
            <i class="far fa-folder"/>
        {:else}
            <i class="far fa-folder-open"/>
        {/if}
        {name}
    </span>
{:else}
    <span class="file"
        on:click={
            () => changer(name)
        }
    >
        <i class="far fa-file-alt"/>
        {RemoveExt(name)}
    </span>
{/if}

{#if control == true}
    <ul>
        {#each inherit.contain as data}
            <li>
                {#if data.dir == true}
                    <svelte:self
                        inherit={data}
                        name={data.name}
                        control={!control}
                        changer={changer}
                    />
                {:else}
                    <span class="file"
                        on:click={
                            () => changer(data.name)
                        }
                    >
                        <i class="far fa-file-alt"/>
                        {RemoveExt(data.name)}
                    </span>
                {/if}
            </li>
        {/each}
    </ul>
{/if}

<style lang="scss">
    %block {
        grid-template-columns: 25px 1fr;
        grid-template-rows: 1fr;
        align-items: center;
        padding: 2.5px 0;
        font-size: 18px;
        cursor: pointer;
        display: grid;
        i {
            padding-right: 5px;
            text-align: center;
        }
        &:hover {
            color: $miho-yellow-a;
        }
    }

    .folder {
        @extend %block;
        i {
            color: $miho-yellow-a;
        }
    }

    .file {
        @extend %block;
        i {
            color: $miho-green-a;
        }
    }
    
    ul {
        border-left: dotted 2.5px $miho-blue-c;
        padding-left: 10px;
        li {
            list-style: none;
        }
    }
</style>