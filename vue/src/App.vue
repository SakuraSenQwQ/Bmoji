<script setup lang="ts">
</script>

<template>
  <div class="app-container">
    <header class="app-header">
      <nav class="app-nav">
        <router-link to="/" class="logo">Bmoji</router-link>
        <div class="nav-links">
          <router-link to="/">首页</router-link>
          <router-link to="/list">表情包</router-link>
        </div>
      </nav>
    </header>
    <main class="app-main">
      <RouterView />
    </main>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC',
    'Hiragino Sans GB', 'Microsoft YaHei', sans-serif;
  background-color: #f5f5f7;
  color: #1d1d1f;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  min-height: 100vh;
}

a {
  text-decoration: none;
  color: inherit;
}

.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.app-header {
  position: sticky;
  top: 0;
  z-index: 100;
  background: rgba(255, 255, 255, 0.78);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-bottom: 1px solid rgba(0, 0, 0, 0.06);
}

.app-nav {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px;
  height: 56px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.logo {
  font-size: 20px;
  font-weight: 700;
  letter-spacing: -0.5px;
  color: #1d1d1f;
}

.nav-links {
  display: flex;
  gap: 24px;
  align-items: center;
}

.nav-links a {
  font-size: 14px;
  color: #86868b;
  transition: color 0.2s;
  padding: 8px 0;
}

.nav-links a:hover,
.nav-links a.router-link-active {
  color: #1d1d1f;
}

.app-main {
  flex: 1;
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px;
  width: 100%;
}

@media (max-width: 640px) {
  .app-nav {
    padding: 0 16px;
  }

  .app-main {
    padding: 0 16px;
  }
}
</style>
