/** 表情包系列项 */
export interface PackageItem {
  id: number
  text: string
  type: number
  url: string
  UrlType: number
}

/** 单个表情 */
export interface EmoteItem {
  text: string
  url: string
  type: number
  gif_url: string
  meta: {
    alias: string
  }
}

/** 表情包详情 */
export interface PackageDetail {
  id: number
  text: string
  emote: EmoteItem[]
}

/** 通用 API 响应包装 */
export interface ApiResponse<T> {
  data: T
  code?: number
  message?: string
}

/** 图片 URL 前缀表 */
export const IMG_BASE_URLS = [
  "https://i0.hdslb.com/bfs/emote/",
  "https://i0.hdslb.com/bfs/emote/",
  "https://i0.hdslb.com/bfs/garb/",
  "https://i0.hdslb.com/bfs/garb/",
] as const

/** 根据 UrlType 拼接完整图片 URL */
export function getImageUrl(urlType: number, path: string): string {
  // 部分接口返回的 url 已经是完整链接，无需再拼接
  if (path.startsWith('http://') || path.startsWith('https://')) {
    return path
  }
  return (IMG_BASE_URLS[urlType] ?? IMG_BASE_URLS[0]) + path
}