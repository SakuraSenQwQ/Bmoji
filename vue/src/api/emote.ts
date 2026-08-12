/** 表情包相关 API */

import { request } from './index'
import type { ApiResponse, PackageItem, PackageDetail } from '@/types'

/** 获取表情包列表 */
export function fetchPackageList() {
  return request<ApiResponse<{ all_packages: PackageItem[] }>>('list')
}

/** 获取单个表情包详情 */
export function fetchPackageDetail(id: string | number) {
  return request<PackageDetail>('info?id=' + id)
}