import { Apis, Configuration } from './generated'

const params = new URLSearchParams(location.search)
const token = params.get('token')

const apis = new Apis(
  new Configuration({
    basePath: '/api',
    accessToken: token ?? undefined
  })
)

export default apis
export * from './generated'
