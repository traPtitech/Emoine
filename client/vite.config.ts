import { defineConfig } from 'vite'
import VuePlugin from '@vitejs/plugin-vue'
import path from 'path'

const devHost = 'localhost:3050'
const prodHost = 'emoine.trap.jp'

export default defineConfig({
  resolve: {
    alias: {
      '/@': path.resolve(__dirname, 'src')
    }
  },
  define: {
    __HOST__: `"${process.env.NODE_ENV === 'production' ? prodHost : devHost}"`
  },
  server: {
    proxy: {
      '/api': {
        target: `http://${devHost}`,
        changeOrigin: true
      }
    }
  },
  plugins: [VuePlugin()]
})
