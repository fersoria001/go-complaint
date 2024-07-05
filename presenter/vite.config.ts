import { defineConfig } from "vite";
import { TanStackRouterVite } from "@tanstack/router-plugin/vite";
import react from "@vitejs/plugin-react-swc";
import tailwindcss from "tailwindcss";
// https://vitejs.dev/config/
export default defineConfig({
  base: "https://www.go-complaint.com/",
  plugins: [react(),TanStackRouterVite(),],
  css: {
    postcss: {
      plugins: [tailwindcss()],
    },
  },
})

/*   resolve: {
    alias: [
      { find: '@assets', replacement: '/src/assets' },
      { find: '@components', replacement: '/src/components' },
      { find: '@pages', replacement: '/src/pages' },
    ],
  }, */
