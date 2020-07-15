import { Apis } from './generated'

const apis = new Apis({ basePath: '/api' })

export default apis
export * from './generated'
