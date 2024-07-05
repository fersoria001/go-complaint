/* eslint-disable @typescript-eslint/no-unused-vars */
import { defineConfig, loadEnv } from "vite";
import { TanStackRouterVite } from "@tanstack/router-plugin/vite";
import react from "@vitejs/plugin-react-swc";
import tailwindcss from "tailwindcss";
// https://vitejs.dev/config/
export default defineConfig(() => {
  return {
    base: "/",
    plugins: [react(), TanStackRouterVite()],
    css: {
      postcss: {
        plugins: [tailwindcss()],
      },
    },
  };
});
