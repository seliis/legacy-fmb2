const preProcessor = require("svelte-preprocess");

module.exports = {
    preprocess: preProcessor({
        defaults: {
            style: "scss"
        },
        scss: {
            prependData: `@import "./src/main.scss";`
        }
    })
}