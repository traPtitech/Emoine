import { UserConfig } from 'vite'
import VuePlugin from '@vitejs/plugin-vue'
// @ts-expect-error prevent installing @types/node
import path from 'path'

declare const __dirname: string
declare const process: { env: Record<string, string> }

const devHost = 'localhost:3050'
const prodHost = 'emoine20.trap.jp'

const config: UserConfig = {
  alias: {
    '/@': path.resolve(__dirname, 'src')
  },
  define: {
    __HOST__: process.env.NODE_ENV === 'production' ? prodHost : devHost
  },
  server: {
    proxy: {
      '/api': {
        target: `http://${devHost}`,
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
