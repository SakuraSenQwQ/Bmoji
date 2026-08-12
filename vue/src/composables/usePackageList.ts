import { ref, onMounted, onUnmounted, computed, nextTick } from 'vue'
import { fetchPackageList } from '@/api/emote'
import type { PackageItem } from '@/types'

export function usePackageList() {
  const rawList = ref<PackageItem[]>([])
  const display = ref<PackageItem[]>([])
  const search = ref('')
  const loading = ref(true)
  const error = ref<string | null>(null)

  let dispNum = 0
  let isPushing = false
  let listEl: HTMLElement | null = null

  const isSearching = computed(() => search.value.trim() !== '')

  const displayedCount = computed(() => display.value.length)
  const totalCount = computed(() => rawList.value.length)

  /** 根据屏幕尺寸动态计算一屏能装下多少个卡片 */
  function calcStep(): number {
    const width = window.innerWidth
    const height = window.innerHeight
    let cols = 5
    if (width < 400) cols = 2
    else if (width < 640) cols = 3
    else if (width < 900) cols = 4
    const cardH = 170 // 估计卡片高度
    const viewH = height - 220 // 减去 header + padding
    const rows = Math.max(4, Math.ceil(viewH / cardH) + 2)
    return cols * rows
  }

  /** 初始加载两批：先填满一屏，再默认预取一批更多表情 */
  function initialFill() {
    push(calcStep())
    push(calcStep())
  }

  async function load() {
    loading.value = true
    error.value = null
    try {
      const res = await fetchPackageList()
      rawList.value = res.data.all_packages
      // 列表加载完成后，若已有搜索词（如来自 URL ?search=），则应用搜索
      if (search.value.trim() !== '') {
        doSearch()
      } else {
        initialFill()
      }
    } catch (e) {
      error.value = e instanceof Error ? e.message : '加载失败'
    } finally {
      loading.value = false
    }
  }

  function push(num: number) {
    for (let i = 0; i < num; i++) {
      dispNum++
      const item = rawList.value[dispNum - 1]
      if (item) {
        display.value.push(item)
      }
    }
  }

  /** 统一滚动处理：通过底部哨兵元素判断是否需要加载更多 */
  function handleScroll() {
    if (isSearching.value || isPushing) return
    const sentinel = document.querySelector('.list-sentinel')
    if (!sentinel) return
    const rect = sentinel.getBoundingClientRect()
    // 当哨兵距离视口底部不足 200px 时加载更多
    if (rect.top <= window.innerHeight + 200) {
      isPushing = true
      push(calcStep())
      nextTick(() => { isPushing = false })
    }
  }

  function doSearch() {
    display.value = []
    const q = search.value.trim()
    if (q === '') {
      dispNum = 0
      initialFill()
      return
    }
    const rx = new RegExp(q, 'i')
    for (const item of rawList.value) {
      if (rx.test(item.text)) {
        display.value.push(item)
      }
    }
  }

  function shuffle() {
    const a = [...rawList.value]
    for (let i = a.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1))
      ;[a[i], a[j]] = [a[j]!, a[i]!]
    }
    rawList.value = a
    dispNum = 0
    display.value = []
    initialFill()
  }

  onMounted(() => {
    load()
    // 等 DOM 就绪后同时监听容器和窗口的 scroll 事件，
    // 无论滚动发生在哪个元素上都能捕获
    nextTick(() => {
      listEl = document.querySelector('.list-contents')
      listEl?.addEventListener('scroll', handleScroll, { passive: true })
      window.addEventListener('scroll', handleScroll, { passive: true })
    })
  })

  onUnmounted(() => {
    listEl?.removeEventListener('scroll', handleScroll)
    window.removeEventListener('scroll', handleScroll)
    listEl = null
  })

  return {
    display,
    search,
    loading,
    error,
    isSearching,
    displayedCount,
    totalCount,
    doSearch,
    shuffle,
    reload: load,
  }
}