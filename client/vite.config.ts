import { UserConfig } from 'vite'
import VuePlugin from '@vitejs/plugin-vue'
// @ts-expect-error prevent installing @types/node
import path from 'path'

declare const __dirname: string

const config: UserConfig = {
  alias: {
    '/@': path.resolve(__dirname, 'src')
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:3050',
        changeOrigin: true
      }
    }
  },
  optimizeDeps: {
    include: ['protobufjs/minimal']
  },
  plugins: [VuePlugin()]
}

export default config
