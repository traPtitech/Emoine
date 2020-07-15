import { UserConfig } from 'vite'
import path from 'path'

const config: UserConfig = {
  minify: 'esbuild',
  alias: {
    '/@/': path.resolve(__dirname, 'src')
  },
  proxy: {
    '/api': {
      target: 'https://example.com',
      changeOrigin: true
    }
  },
  optimizeDeps: {
    include: ['protobufjs/minimal']
  }
}

export default config
