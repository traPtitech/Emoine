interface UrlObject {
  query: Record<string, string>
  search: string
  url: URL
}

const BASE = 'https://example.com'
const baseReg = new RegExp(`^${BASE}`)

export const parse = (urlStr: string, parseQueryString: true): UrlObject => {
  const url = new URL(urlStr, BASE)
  const query = parseQueryString
    ? Object.fromEntries(url.searchParams.entries())
    : {}

  return {
    query,
    search: url.search,
    url
  }
}

export const format = (urlObject: UrlObject): string => {
  const url = urlObject.url
  url.search = ''
  Object.entries(urlObject.query).forEach(([k, v]) => {
    url.searchParams.set(k, v)
  })
  return url.href.replace(baseReg, '')
}
