<script setup lang="ts">
import { ref } from 'vue'
import { useEmoteDetail } from '@/composables/useEmoteDetail'
import { fetchBlob } from '@/api'
import { saveAs } from 'file-saver'
import JSZip from 'jszip'

const props = defineProps<{ id: number }>()
const emit = defineEmits<{ close: [] }>()

const { detail, loading, error, hasGif, reload } = useEmoteDetail(props.id)

const viewgif = ref(false)
const preview = ref<{ display: boolean; url: string; name: string; alias: string }>({
  display: false,
  url: '',
  name: '',
  alias: '',
})
const todownload = ref(false)

function toggleGif() {
  viewgif.value = !viewgif.value
}

function selectEmote(index: number) {
  const e = detail.value!.emote[index]
  if (e.type === 4) return
  preview.value = {
    display: true,
    url: (viewgif.value && e.gif_url ? e.gif_url : e.url) as string,
    name: e.text,
    alias: e.meta?.alias ?? '',
  }
}

function closePreview() {
  preview.value.display = false
}

function handleBackdropClick(e: MouseEvent) {
  if ((e.target as HTMLElement).classList.contains('overlay-backdrop')) {
    emit('close')
  }
}

function copyText(text: string) {
  navigator.clipboard.writeText(text).then(() => {
    alert('复制成功')
  })
}

function openUrl(url: string) {
  // 使用带 rel="noopener noreferrer" 的 <a> 模拟点击，不发送 Referer 头
  const a = document.createElement('a')
  a.href = url
  a.target = '_blank'
  a.rel = 'noopener noreferrer'
  a.style.display = 'none'
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
}
async function downloadPackage(mode: number) {
  if (!detail.value) return
  const zip = new JSZip()
  const emote = detail.value.emote

  if (mode === 1) {
    // 静态
    await Promise.all(
      emote.map(async (e) => {
        const blob = await fetchBlob(e.url)
        const ext = blob.type.split("/")[1] || "png"
        zip.file(e.text + "." + ext, blob)
      }),
    )
    const content = await zip.generateAsync({ type: "blob" })
    saveAs(content, detail.value.text + ".zip")
  } else if (mode === 2 && hasGif.value) {
    // 动图
    await Promise.all(
      emote.map(async (e) => {
        if (e.gif_url) {
          const blob = await fetchBlob(e.gif_url)
          const ext = blob.type.split("/")[1] || "gif"
          zip.file(e.text + "." + ext, blob)
        }
      }),
    )
    const content = await zip.generateAsync({ type: "blob" })
    saveAs(content, detail.value.text + ".zip")
  } else if (mode === 3) {
    // JSON
    const blob = new Blob([JSON.stringify(emote, null, 2)], { type: "application/json" })
    saveAs(blob, detail.value.text + ".json")
  } else if (mode === 4) {
    // Waline
    const waline: { icon: string; name: string; type: string; prefix: string; items: string[] } = {
      icon: "",
      name: detail.value.text,
      type: "",
      prefix: "",
      items: [],
    }
    const first = emote[0]
    if (first?.type === 4) {
      waline.prefix = ""
      waline.type = "text"
      waline.items = emote.map((e) => e.text)
    } else {
      waline.prefix = ""
      waline.type = "image"
      await Promise.all(
        emote.map(async (e) => {
          const url = e.gif_url || e.url
          const blob = await fetchBlob(url)
          const ext = blob.type.split("/")[1] || "png"
          zip.file(e.text + "." + ext, blob)
          waline.items.push(e.text)
        }),
      )
      waline.icon = waline.items[0] ?? ""
      zip.file("info.json", new Blob([JSON.stringify(waline)], { type: "application/json" }))
    }
    const content = await zip.generateAsync({ type: "blob" })
    saveAs(content, detail.value.text + ".zip")
  }

  todownload.value = false
}
</script>
<template>
  <Teleport to="body">
    <div class="overlay-backdrop" @click="handleBackdropClick">
      <div class="overlay-panel" @click.stop>
        <!-- 头部 -->
        <div class="overlay-header">
          <button class="btn-icon" @click="emit('close')" title="返回">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <line x1="19" y1="12" x2="5" y2="12" />
              <polyline points="12 19 5 12 12 5" />
            </svg>
          </button>
          <div class="overlay-title-group" v-if="!loading && detail">
            <h2 class="overlay-title">{{ detail.text }}</h2>
            <span class="overlay-meta">{{ detail.emote.length }} 个表情</span>
          </div>
          <div class="overlay-title-group" v-else-if="loading">
            <h2 class="overlay-title">加载中…</h2>
          </div>
          <div class="overlay-actions" v-if="!loading && detail">
            <label class="toggle-gif" v-if="hasGif" @click="toggleGif" :title="viewgif ? '切换至静态' : '切换至动图'">
              <span class="toggle-track" :class="{ active: viewgif }">
                <span class="toggle-thumb" />
              </span>
            </label>
            <button class="btn-icon" @click="todownload = true" title="下载">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
                <polyline points="7 10 12 15 17 10" />
                <line x1="12" y1="15" x2="12" y2="3" />
              </svg>
            </button>
          </div>
        </div>

        <!-- 加载态 -->
        <div class="overlay-body" v-if="loading">
          <div class="state-msg">加载中…</div>
        </div>

        <!-- 错误态 -->
        <div class="overlay-body" v-else-if="error">
          <div class="state-msg state-error">{{ error }}</div>
        </div>

        <!-- 表情网格 -->
        <div class="overlay-body" v-else-if="detail">
          <div class="emote-grid">
            <template v-for="(e, i) in detail.emote" :key="i">
              <div class="emote-card" @click="selectEmote(i)" v-if="e.type !== 4">
                <div class="emote-image">
                  <img referrerpolicy="no-referrer" loading="lazy" :src="viewgif && e.gif_url ? e.gif_url : e.url" :alt="e.text" />
                </div>
                <p class="emote-name">{{ e.text }}</p>
              </div>
              <div class="emote-card emote-card-text" v-else>
                <div class="emote-text-content">
                  <span class="emote-text">{{ e.text }}</span>
                </div>
                <p class="emote-name">颜文字</p>
              </div>
            </template>
          </div>
        </div>

        <!-- 预览弹窗 -->
        <Teleport to="body">
          <div v-if="preview.display" class="preview-overlay" @click="closePreview">
            <div class="preview-panel" @click.stop>
              <button class="preview-close" @click="closePreview">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <line x1="18" y1="6" x2="6" y2="18" /><line x1="6" y1="6" x2="18" y2="18" />
                </svg>
              </button>
              <img class="preview-image" :src="preview.url" :alt="preview.name" referrerpolicy="no-referrer" />
              <div class="preview-details">
                <h3>{{ preview.name }}</h3>
                <p v-if="preview.alias">别名: {{ preview.alias }}</p>
                <div class="preview-actions">
                  <button class="btn btn-sm" @click="copyText(preview.url)">复制链接</button>
                  <button class="btn btn-sm" @click="copyText('![' + preview.name + '](' + preview.url + ')')">复制 Markdown</button>
                  <button class="btn btn-sm" @click="openUrl(preview.url)">新标签页打开</button>
                </div>
              </div>
            </div>
          </div>
        </Teleport>

        <!-- 下载弹窗 -->
        <Teleport to="body">
          <div v-if="todownload" class="preview-overlay" @click="todownload = false">
            <div class="preview-panel download-panel" @click.stop>
              <h2 class="download-title">下载</h2>
              <div class="download-options">
                <button class="btn btn-primary" @click="downloadPackage(1)" v-if="detail && detail.emote[0].type !== 4">下载压缩包 (静态)</button>
                <button class="btn btn-primary" @click="downloadPackage(2)" v-if="hasGif">下载压缩包 (动图)</button>
                <button class="btn btn-outline" @click="downloadPackage(3)">下载 JSON</button>
                <button class="btn btn-outline" @click="downloadPackage(4)">保存为 Waline 表情包</button>
              </div>
              <button class="btn btn-ghost" @click="todownload = false">取消</button>
            </div>
          </div>
        </Teleport>
      </div>
    </div>
  </Teleport>
</template>
<style scoped>
@keyframes panelIn {
  from { opacity: 0; transform: translateY(8px) scale(0.98); }
  to   { opacity: 1; transform: translateY(0) scale(1); }
}

.overlay-backdrop {
  position: fixed;
  inset: 0;
  z-index: 1000;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
  backdrop-filter: blur(4px);
}

.overlay-panel {
  background: #fff;
  border-radius: 20px;
  max-width: 640px;
  width: 100%;
  max-height: 90dvh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  box-shadow: 0 24px 64px rgba(0, 0, 0, 0.2);
  animation: panelIn 0.2s ease;
}

.overlay-header {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px 20px;
  border-bottom: 1px solid rgba(0, 0, 0, 0.06);
  flex-shrink: 0;
}

.btn-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border: 1.5px solid #d2d2d7;
  border-radius: 10px;
  background: #fff;
  cursor: pointer;
  color: #1d1d1f;
  transition: all 0.2s;
  flex-shrink: 0;
}

.btn-icon:hover {
  background: #f5f5f7;
  border-color: #1d1d1f;
}

.overlay-title-group {
  flex: 1;
  min-width: 0;
}

.overlay-title {
  font-size: 18px;
  font-weight: 700;
  letter-spacing: -0.3px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.overlay-meta {
  font-size: 13px;
  color: #86868b;
}

.overlay-actions {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-shrink: 0;
}

.toggle-gif {
  cursor: pointer;
  display: flex;
  align-items: center;
}

.toggle-track {
  position: relative;
  width: 36px;
  height: 20px;
  background: #e8e8ed;
  border-radius: 10px;
  transition: background 0.2s;
}

.toggle-track.active {
  background: #1d1d1f;
}

.toggle-thumb {
  position: absolute;
  top: 2px;
  left: 2px;
  width: 16px;
  height: 16px;
  background: #fff;
  border-radius: 50%;
  transition: transform 0.2s;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.15);
}

.toggle-track.active .toggle-thumb {
  transform: translateX(16px);
}

.overlay-body {
  flex: 1;
  overflow-y: auto;
  overflow-x: hidden;
  padding: 16px 20px 20px;
  scrollbar-width: thin;
  scrollbar-color: #d2d2d7 transparent;
}

.state-msg {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 60px 0;
  color: #86868b;
  font-size: 15px;
}

.state-error {
  color: #ff3b30;
}

.emote-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(80px, 1fr));
  gap: 10px;
}

.emote-card {
  background: #fafafa;
  border-radius: 10px;
  padding: 10px;
  cursor: pointer;
  transition: all 0.2s;
  border: 1px solid rgba(0, 0, 0, 0.04);
  text-align: center;
}

.emote-card:hover {
  background: #fff;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08);
  transform: translateY(-2px);
}

.emote-card-text {
  cursor: default;
  background: #f5f5f7;
}

.emote-image {
  width: 100%;
  aspect-ratio: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 6px;
}

.emote-image img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.emote-text-content {
  width: 100%;
  aspect-ratio: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 6px;
  padding: 4px;
  word-break: break-all;
  font-size: 12px;
  color: #86868b;
}

.emote-name {
  font-size: 11px;
  color: #86868b;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
/* ── 预览弹窗 ── */
.preview-overlay {
  position: fixed;
  inset: 0;
  z-index: 1100;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
  backdrop-filter: blur(4px);
}

.preview-panel {
  background: #fff;
  border-radius: 20px;
  padding: 32px;
  max-width: 420px;
  width: 100%;
  position: relative;
  text-align: center;
  box-shadow: 0 24px 64px rgba(0, 0, 0, 0.2);
  animation: panelIn 0.2s ease;
}

.preview-close {
  position: absolute;
  top: 12px;
  right: 12px;
  background: none;
  border: none;
  cursor: pointer;
  color: #86868b;
  padding: 4px;
  border-radius: 8px;
  transition: all 0.2s;
}

.preview-close:hover {
  background: #f5f5f7;
  color: #1d1d1f;
}

.preview-image {
  width: 140px;
  height: 140px;
  object-fit: contain;
  margin-bottom: 16px;
}

.preview-details h3 {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 4px;
}

.preview-details p {
  font-size: 13px;
  color: #86868b;
  margin-bottom: 14px;
}

.preview-actions {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  justify-content: center;
}

/* ── 下载面板 ── */
.download-panel {
  max-width: 320px;
}

.download-title {
  font-size: 18px;
  font-weight: 700;
  margin-bottom: 16px;
}

.download-options {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-bottom: 16px;
}

.download-options .btn {
  width: 100%;
}

/* ── 通用按钮 ── */
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 10px 20px;
  font-size: 14px;
  font-weight: 500;
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.2s ease;
  border: 1.5px solid transparent;
  font-family: inherit;
  white-space: nowrap;
}

.btn-sm {
  padding: 7px 14px;
  font-size: 12px;
}

.btn-primary {
  background: #1d1d1f;
  color: #fff;
  border-color: #1d1d1f;
}

.btn-primary:hover {
  background: #000;
  border-color: #000;
}

.btn-outline {
  background: transparent;
  color: #1d1d1f;
  border-color: #d2d2d7;
}

.btn-outline:hover {
  background: #f5f5f7;
  border-color: #1d1d1f;
}

.btn-ghost {
  background: transparent;
  color: #86868b;
  border-color: transparent;
}

.btn-ghost:hover {
  color: #1d1d1f;
  background: #f5f5f7;
}

/* ── 响应式 ── */
@media (max-width: 640px) {
  .overlay-backdrop {
    padding: 0;
  }

  .overlay-panel {
    max-width: 100%;
    max-height: 100dvh;
    border-radius: 0;
  }

  .overlay-header {
    padding: 12px 16px;
  }

  .overlay-body {
    padding: 12px 16px 16px;
  }

  .emote-grid {
    grid-template-columns: repeat(auto-fill, minmax(70px, 1fr));
    gap: 8px;
  }

  .preview-panel {
    padding: 24px;
    margin: 16px;
  }

  .preview-image {
    width: 120px;
    height: 120px;
  }
}

@media (max-width: 400px) {
  .emote-grid {
    grid-template-columns: repeat(3, 1fr);
    gap: 6px;
  }

  .emote-card {
    padding: 6px;
  }
}
</style>
