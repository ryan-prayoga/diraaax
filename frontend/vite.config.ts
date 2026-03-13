import { defineConfig } from "vite";
import { sveltekit } from "@sveltejs/kit/vite";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
  plugins: [sveltekit(), tailwindcss()],
  preview: {
    host: "127.0.0.1",
    port: 4001,
    allowedHosts: ["diraaax.ryannn.net"],
  },
});
