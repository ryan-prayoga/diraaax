import { defineConfig } from "vite";
import { sveltekit } from "@sveltejs/kit/vite";

export default defineConfig({
  plugins: [sveltekit()],
  preview: {
    host: "127.0.0.1",
    port: 4001,
    allowedHosts: ["diraaax.ryannn.net"],
  },
});
