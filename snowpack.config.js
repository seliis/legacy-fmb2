module.exports = {
    mount: {
        public: {
            url: "/"
        },
        src: {
            url: "/dist"
        }
    },
    plugins: [
        "@snowpack/plugin-svelte",
        "@snowpack/plugin-sass"
    ],
    devOptions: {
        open: "none",
        hmr: true
    }
}