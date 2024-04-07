import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { VitePWA } from "vite-plugin-pwa";

// Manifest configuration
const manifest = {
    name: "Tasks",
    short_name: "Tasks",
    start_url: "/",
    display: "standalone",
    background_color: "#18181b",
    theme_color: "#18181b",
    description: "Tasks - ttgowebddv2",
    icons: [
        {
            src: "icon/lowres.webp",
            sizes: "48x48",
            type: "image/webp",
        },
    ],
};

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react(), VitePWA({ manifest })],
});
