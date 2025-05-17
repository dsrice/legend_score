import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import viteTsconfigPaths from 'vite-tsconfig-paths';
import svgr from 'vite-plugin-svgr';
import { resolve } from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    viteTsconfigPaths(),
    svgr(),
  ],
  server: {
    port: 3000, // Default port for Create React App
    open: true, // Open browser on start
    host: true, // Listen on all addresses
  },
  build: {
    outDir: 'build', // Same output directory as Create React App
  },
  publicDir: 'public', // Specify the public directory (same as Create React App)
  resolve: {
    alias: {
      // Add any path aliases if needed
      '@': resolve(__dirname, 'src'),
    },
  },
});