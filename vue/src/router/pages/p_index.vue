<script lang="ts" setup>
import { ref } from 'vue';
import router from '..';
const prop = defineProps({
  list: {
    type: Array,
    required: true,
  }
})
const list = ref(prop.list)
</script>

<template>
  <div class="index">
    <div class="info">
      <h1>Bmoji - B站表情包大全</h1>
      <p>在这里，你可以找到大部分的哔哩哔哩表情包</p>
      <p>本项目基于<a href="">哔哩哔哩API</a></p>
      <p v-if="list.length !== 0">共计{{ list.length }}个系列</p>
      <p v-if="list.length === 0">正在获取列表...</p>
      <img src="/img/amy.gif" alt="">
      <div class="bts">
        <span><a href="https://github.com/SakuraSenQwQ/Bmoji" target="_blank">GitHub</a></span>
        <span @click="router.push('list')">去查找</span>
      </div>
    </div>
  </div>
</template>


<style scoped>



.info {
  width: fit-content;
  background-color: #FB7299;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: #18191C 0 0 12px;
  img{
    height: 128px;
    border-radius: 12px;
    margin: 1rem 0;
  }
}

.index {
  display: flex;
  flex-direction: column;
  width: 100%;

  justify-content: center;
  height: 100%;

  margin-top: 30dvh;
  align-items: center;
}
@media (width < 450px) {
  .index{
    justify-self: center;
    margin-top: 10dvh;
    width: 80%;
  }
}
</style>
