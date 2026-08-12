/** 统一 API 请求模块 */

const BASE_URL = "https://api.3mua.cn/api/bmoji/v1/emote/"

export class ApiError extends Error {
  constructor(public status: number, message: string) {
    super(message)
    this.name = 'ApiError'
  }
}

/**
 * 统一请求函数
 * 所有 API 调用都通过此函数，便于统一错误处理、拦截、缓存等
 */
export async function request<T>(endpoint: string, options?: RequestInit): Promise<T> {
  const res = await fetch(BASE_URL + endpoint, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })

  if (!res.ok) {
    throw new ApiError(res.status, `请求失败: ${res.status} ${res.statusText}`)
  }

  return res.json() as Promise<T>
}

/** 下载单个文件并返回 Blob */
export async function fetchBlob(url: string): Promise<Blob> {
  const resp = await fetch(url, {
    method: "GET",
    referrerPolicy: "no-referrer",
  })
  return resp.blob()
}
