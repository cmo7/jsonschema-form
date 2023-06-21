import react from '@vitejs/plugin-react-swc';
import { defineConfig } from 'vite';
import mkcert from 'vite-plugin-mkcert';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), mkcert()],
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          react: ['react', 'react-dom'],
          reactRouter: ['react-router', 'react-router-dom'],
          reactQuery: ['react-query', 'react-query/devtools'],
          chakraui: ['@chakra-ui/react', '@chakra-ui/icons'],
          rjsf: ['@rjsf/core', '@rjsf/chakra-ui'],
        },
      },
    },
  },
});
