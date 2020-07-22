import { UserConfig } from 'vite'
import path from 'path'

const config: UserConfig = {
  alias: {
    '/@/': path.resolve(__dirname, 'src')
  },
  proxy: {
    '/api': {
      target: 'http://localhost:3050',
      changeOrigin: true
    }
  },
  optimizeDeps: {
    include: ['protobufjs/minimal']
  }
}

export default config
