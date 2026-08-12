<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { usePackageList } from '@/composables/usePackageList'
import { getImageUrl } from '@/types'
import PackageOverlay from '@/components/PackageOverlay.vue'

const route = useRoute()
const router = useRouter()

const {
  display,
  search,
  loading,
  error,
  isSearching,
  displayedCount,
  totalCount,
  doSearch,
  shuffle,
  reload,
} = usePackageList()

const selectedPackageId = ref<number | null>(null)

function parseId(val: unknown): number | null {
  const s = Array.isArray(val) ? val[0] : val
  if (typeof s !== 'string') return null
  const n = Number(s)
  return Number.isInteger(n) && n > 0 ? n : null
}

function openDetail(id: number) {
  selectedPackageId.value = id
  // 同步 ?id=xxx 到 URL，便于分享/收藏链接
  router.replace({ query: { ...route.query, id: String(id) } })
}

function closeOverlay() {
  selectedPackageId.value = null
  const query = { ...route.query }
  delete query.id
  router.replace({ query })
}

// URL ?id=xxx → 打开对应表情包详情（进入页面 / 刷新 / 前进后退时自动恢复）
watch(
  () => route.query.id,
  (val) => {
    selectedPackageId.value = parseId(val)
  },
  { immediate: true },
)

// 搜索框输入 → 更新列表并同步 ?search= 到 URL
watch(search, (val) => {
  doSearch()
  const q = val.trim()
  const query = { ...route.query }
  if (q) {
    query.search = q
  } else {
    delete query.search
  }
  router.replace({ query })
})

// URL ?search=xxx → 恢复搜索框（来自外部链接 / 前进后退）
watch(
  () => route.query.search,
  (val) => {
    const q = typeof val === 'string' ? val : ''
    if (search.value !== q) {
      search.value = q
    }
  },
  { immediate: true },
)
</script>

<template>
  <div class="list-page">
    <div class="list-header">
      <div class="list-header-inner">
        <h1 class="list-title">表情包</h1>
        <div class="list-controls">
          <span class="list-count">{{ displayedCount }} / {{ totalCount }}</span>
          <button class="btn-shuffle" @click="shuffle" title="随机打乱">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="16 3 21 3 21 8" />
              <line x1="4" y1="20" x2="21" y2="3" />
              <polyline points="21 16 21 21 16 21" />
              <line x1="15" y1="15" x2="21" y2="21" />
              <line x1="4" y1="4" x2="9" y2="9" />
            </svg>
          </button>
          <div class="search-wrapper">
            <svg class="search-icon" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="11" cy="11" r="8" />
              <line x1="21" y1="21" x2="16.65" y2="16.65" />
            </svg>
            <input
              v-model="search"
              class="search-input"
              type="search"
              placeholder="搜索表情包..."
            />
          </div>
        </div>
      </div>
    </div>

    <div class="list-contents">
      <div class="empty-state" v-if="loading">加载中…</div>
      <div class="empty-state error" v-else-if="error">{{ error }}</div>
      <div class="list-disp" v-else-if="display.length > 0">
        <div
          class="card"
          v-for="item in display"
          :key="item.id"
          @click="openDetail(item.id)"
        >
          <div class="card-image">
            <img
              :src="getImageUrl(item.UrlType, item.url)"
              :alt="item.text"
              referrerpolicy="no-referrer"
              loading="lazy"
            />
          </div>
          <div class="card-info">
            <p class="card-name">{{ item.text }}</p>
          </div>
        </div>
        <div class="list-sentinel"></div>
      </div>
      <div class="empty-state" v-else>
        <p v-if="isSearching">没有找到匹配的表情包</p>
        <p v-else>暂无数据</p>
      </div>
    </div>
  </div>

  <PackageOverlay
    v-if="selectedPackageId"
    :key="selectedPackageId"
    :id="selectedPackageId"
    @close="closeOverlay"
  />
</template>

<style scoped>
.list-page {
  display: flex;
  flex-direction: column;
  height: calc(100vh - 56px);
  padding: 24px 0;
}

.list-header {
  flex-shrink: 0;
  margin-bottom: 24px;
}

.list-header-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  flex-wrap: wrap;
}

.list-title {
  font-size: 28px;
  font-weight: 700;
  letter-spacing: -0.5px;
}

.list-controls {
  display: flex;
  align-items: center;
  gap: 16px;
  flex-wrap: wrap;
}

.list-count {
  font-size: 14px;
  color: #86868b;
  font-weight: 500;
}

.btn-shuffle {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border: 1.5px solid #d2d2d7;
  border-radius: 10px;
  background: #fff;
  cursor: pointer;
  color: #86868b;
  transition: all 0.2s;
}

.btn-shuffle:hover {
  background: #f5f5f7;
  color: #1d1d1f;
  border-color: #1d1d1f;
}

.search-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.search-icon {
  position: absolute;
  left: 12px;
  color: #86868b;
  pointer-events: none;
}

.search-input {
  width: 220px;
  height: 40px;
  padding: 0 16px 0 40px;
  font-size: 14px;
  border: 1.5px solid #d2d2d7;
  border-radius: 12px;
  background: #fff;
  outline: none;
  transition: all 0.2s;
  font-family: inherit;
  color: #1d1d1f;
}

.search-input:focus {
  border-color: #1d1d1f;
  box-shadow: 0 0 0 3px rgba(0, 0, 0, 0.06);
}

.search-input::placeholder {
  color: #a1a1a6;
}

.list-contents {
  flex: 1;
  overflow-y: auto;
  overflow-x: hidden;
  scrollbar-width: thin;
  scrollbar-color: #d2d2d7 transparent;
}

.list-disp {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  gap: 16px;
  padding-bottom: 40px;
}

.list-sentinel {
  grid-column: 1 / -1;
  height: 1px;
}

.card {
  background: #fff;
  border-radius: 14px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.2s ease;
  border: 1px solid rgba(0, 0, 0, 0.04);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.04);
}

.card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
  border-color: rgba(0, 0, 0, 0.08);
}

.card-image {
  width: 100%;
  aspect-ratio: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 16px;
  background: #fafafa;
}

.card-image img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.card-info {
  padding: 10px 12px;
  border-top: 1px solid rgba(0, 0, 0, 0.04);
}

.card-name {
  font-size: 13px;
  font-weight: 500;
  color: #1d1d1f;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  text-align: center;
}

.empty-state {
  display: flex;
  justify-content: center;
  padding: 60px 0;
  color: #86868b;
  font-size: 15px;
}

.empty-state.error {
  color: #ff3b30;
}

@media (max-width: 640px) {
  .list-page {
    padding: 16px 0;
  }

  .list-header-inner {
    flex-direction: column;
    align-items: flex-start;
  }

  .list-controls {
    width: 100%;
  }

  .search-wrapper {
    flex: 1;
  }

  .search-input {
    width: 100%;
  }

  .list-disp {
    grid-template-columns: repeat(auto-fill, minmax(100px, 1fr));
    gap: 12px;
  }

  .list-title {
    font-size: 24px;
  }
}

@media (max-width: 400px) {
  .list-disp {
    grid-template-columns: repeat(2, 1fr);
    gap: 10px;
  }
}
</style>