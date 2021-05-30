<script>
    import { slide } from "svelte/transition";

    const menu = [
        "aviators",
        "training",
        "document",
        "campaign",
        "database"
    ]

    let slave = false;
    function SlaveControl() {
        slave = !slave;
    }

    function ChangePage(t) {
        window.location.href = "#/" + t;
        SlaveControl();
    }

    function LoadHome() {
        window.location.href = "#";
        if (slave == true) {
            SlaveControl();
        }
    }
</script>

<header>
    <div id="master">
        <h1 on:click={LoadHome}>
            VFA-314 FULL METAL BITCHES
        </h1>
        <i class="fas fa-bars"
            on:click={SlaveControl}
        />
    </div>
    {#if slave == true}
        <ul id="slave" transition:slide>
            {#each menu as m}
                <li on:click={
                    () => ChangePage(m)
                }>
                    {m}
                </li>
            {/each}
        </ul>
    {/if}
</header>

<style lang="scss">
    header {
        font-family: "Quicksand", sans-serif;
        background-color: $miho-black-a;
        color: $miho-white-a;
        height: 100%;
        z-index: 100;
    }

    #master {
        justify-content: space-between;
        align-items: center;
        padding: 0 25px;
        display: flex;
        height: 100%;
        h1, i {
            font-size: 16px;
            cursor: pointer;
        }
    }

    #slave {
        background-color: $miho-black-a;
        position: relative;
        padding: 0 25px;
        li {
            text-transform: capitalize;
            list-style: none;
            font-size: 16px;
            cursor: pointer;
            padding: 8px 0;
            &:hover {
                color: $miho-red-a;
            }
        }
    }
</style>