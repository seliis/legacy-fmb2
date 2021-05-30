<script>
    import Router from "svelte-spa-router";
    import { wrap } from "svelte-spa-router/wrap"

    // Components
    import Header from "./component/header.svelte";

    // Pages
    import Home from "./routes/home.svelte";

    // Routes
    const routes = {
        "/": Home,
        "/document/:target?": wrap({
            asyncComponent: () => import("./routes/document.svelte")
        })
    }
</script>

<svelte:head>
    <!-- Title -->
    <title>VFA-314</title>
    
    <!-- Font Awesome -->
    <script src="https://kit.fontawesome.com/e0292d2e49.js" crossorigin="anonymous"></script>
    
    <!-- Google Font -->
    <link rel="preconnect" href="https://fonts.gstatic.com">
    <link href="https://fonts.googleapis.com/css2?family=Quicksand:wght@300;500;700&display=swap" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css2?family=Noto+Sans+KR:wght@300;500;700&display=swap" rel="stylesheet">
</svelte:head>

<main>
    <Header/>
    <Router {routes}/>
</main>

<style lang="scss">
    :global(*) {
        margin: 0;
        padding: 0;
    }

    :global(body) {
        font-family: "Noto Sans KR", sans-serif;
        position: relative;
        overflow: hidden;
        height: 100vh;
        width: 100vw;
    }

    :global(section) {
        position: relative;
        height: 100%;
        width: 100%;
    }

    main {
        grid-template-rows: 50px 1fr;
        grid-template-columns: 1fr;
        display: grid;
        height: 100%;
        width: 100%;
    }
</style>