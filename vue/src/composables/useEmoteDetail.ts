/** 表情包详情 Composable */

import { ref, onMounted } from 'vue'
import { fetchPackageDetail } from '@/api/emote'
import type { PackageDetail } from '@/types'

export function useEmoteDetail(id: string | number) {
  const detail = ref<PackageDetail | null>(null)
  const loading = ref(true)
  const error = ref<string | null>(null)
  const hasGif = ref(false)

  async function load() {
    loading.value = true
    error.value = null
    try {
      const data = await fetchPackageDetail(id)
      detail.value = data
      hasGif.value = data.emote.some((e) => e.gif_url !== '')
    } catch (e) {
      error.value = e instanceof Error ? e.message : '加载失败'
    } finally {
      loading.value = false
    }
  }

  onMounted(load)

  return {
    detail,
    loading,
    error,
    hasGif,
    reload: load,
  }
}