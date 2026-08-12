<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { fetchPackageList } from '@/api/emote'

const router = useRouter()
const totalCount = ref(0)
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await fetchPackageList()
    totalCount.value = res.data.all_packages.length
  } catch {
    totalCount.value = 0
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="home">
    <div class="hero">
      <div class="hero-icon">
        <img src="/img/amy.gif" alt="Bmoji" />
      </div>
      <h1 class="hero-title">Bmoji</h1>
      <p class="hero-subtitle">B站表情包大全</p>
      <p class="hero-desc">在这里，你可以找到大部分的哔哩哔哩表情包</p>
      <p class="hero-count" v-if="!loading">共收录 {{ totalCount }} 个系列</p>
      <p class="hero-count loading" v-else>正在获取列表...</p>
      <div class="hero-actions">
        <button class="btn btn-primary" @click="router.push('/list')">浏览表情包</button>
        <a
          class="btn btn-outline"
          href="https://github.com/SakuraSenQwQ/Bmoji"
          target="_blank"
          rel="noopener"
        >GitHub</a>
      </div>
    </div>
  </div>
</template>

<style scoped>
.home {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: calc(100vh - 56px);
  padding: 40px 0;
}

.hero {
  text-align: center;
  max-width: 520px;
  width: 100%;
  animation: fadeIn 0.6s ease;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(12px); }
  to   { opacity: 1; transform: translateY(0); }
}

.hero-icon {
  margin-bottom: 24px;
}

.hero-icon img {
  width: 96px;
  height: 96px;
  border-radius: 24px;
  object-fit: cover;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.hero-title {
  font-size: 48px;
  font-weight: 800;
  letter-spacing: -1.5px;
  margin-bottom: 8px;
  color: #1d1d1f;
}

.hero-subtitle {
  font-size: 20px;
  font-weight: 600;
  color: #86868b;
  margin-bottom: 16px;
}

.hero-desc {
  font-size: 16px;
  color: #86868b;
  line-height: 1.6;
  margin-bottom: 8px;
}

.hero-count {
  font-size: 14px;
  color: #a1a1a6;
  margin-bottom: 32px;
}

.hero-count.loading {
  animation: pulse 1.5s ease infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50%      { opacity: 0.4; }
}

.hero-actions {
  display: flex;
  gap: 12px;
  justify-content: center;
  flex-wrap: wrap;
}

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 12px 28px;
  font-size: 15px;
  font-weight: 500;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
  border: 1.5px solid transparent;
  font-family: inherit;
}

.btn-primary {
  background: #1d1d1f;
  color: #fff;
  border-color: #1d1d1f;
}

.btn-primary:hover {
  background: #000;
  border-color: #000;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.btn-outline {
  background: transparent;
  color: #1d1d1f;
  border-color: #d2d2d7;
}

.btn-outline:hover {
  background: #f5f5f7;
  border-color: #1d1d1f;
  transform: translateY(-1px);
}

@media (max-width: 640px) {
  .hero-title {
    font-size: 36px;
  }

  .hero-subtitle {
    font-size: 18px;
  }

  .hero-icon img {
    width: 80px;
    height: 80px;
  }

  .btn {
    padding: 10px 24px;
    font-size: 14px;
  }
}
</style>