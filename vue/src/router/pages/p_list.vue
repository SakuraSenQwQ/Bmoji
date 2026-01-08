<script lang="ts" setup>
import { onMounted, ref, watch, type PropType } from 'vue';
const display = ref<Listobj[]>([])
type Listobj = {
  "id": number,
  "text": string,
  "type": number,
  "url": string,
  "UrlType": number
}
const prop = defineProps({
  list: {
    type: Array as PropType<Listobj[]>,
    required: true,
  }
})
const rawlist = ref<Listobj[]>(prop.list)

const search = ref({
  input: ""
})
let dispNum = 1

function Push(num: number) {
  console.log("Push")
  for (let i = 0; i < num; i++) {
    dispNum++
    display.value.push(rawlist.value[dispNum - 1]!)
  }
}
watch(rawlist.value, () => {
  if (rawlist.value.length !== 0) {
    Start()
  }
})
onMounted(() => {
  Start()
})
function update() {
  if (searchis.value) {
    return
  }
  const dipC = document.getElementById("disp")!.clientHeight
  const contentC = document.getElementsByClassName("contents")[0]!.clientHeight
  const contentS = document.getElementsByClassName("contents")[0]!.scrollTop
  if (contentS + contentC >= dipC - 100) {
    Push(20)
  }
}
const searchis = ref(false)
watch(search.value, () => {
  display.value = []
  if (search.value.input !== "") {
    searchis.value = true
    const rx = RegExp(search.value.input)
    for (let i = 0; i < rawlist.value.length; i++) {
      if (rx.test(rawlist.value[i]!.text)) {
        display.value.push(rawlist.value[i]!)
      }
    }
  } else {
    dispNum = 0
    searchis.value = false
    Push(100)
  }
})
function Start() {
  setTimeout(() => {
    if (rawlist.value.length <= 0) {
      return
    }
    Push(100)
  }, 0);
  document.getElementsByClassName("contents")[0]!.addEventListener('scroll', () => {
    update()
  })
}
function open(id: number) {
  window.open("http://" + window.location.host + "/info?" + id)
}
function rd() {
  const a = rawlist.value
  for (let i = a.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));

    [a[i], a[j]] = [a[j]!, a[i]!];
  }
  dispNum = 0
  display.value = []
}

const imglist = ["https://i0.hdslb.com/bfs/emote/", "https://i0.hdslb.com/bfs/emote/", "https://i0.hdslb.com/bfs/garb/", "https://i0.hdslb.com/bfs/garb/"]
</script>

<template>
  <div class="list">
    <div class="topbar">
      <h1>表情包</h1>

      <div class="sl">
        <span>{{ display.length + "/" + list.length }}</span>
        <span @click="rd">随机打乱</span>
        <input class="search" v-model="search.input" :placeholder="'共' + list.length + '个表情，点击搜索'" type="text">

      </div>

    </div>
    <div class="contents">
      <div class="disp" id="disp">
        <div class="box" v-for="v in display" :key="v.id" @click="open(v.id)">
          <img referrerpolicy="no-referrer" loading="lazy" :src="(imglist[v.UrlType] ?? '') + v.url" :alt="v.text">
          <div class="info">
            <p :title="v.text">{{ v.text }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
* {
  scrollbar-width: none;
}

.disp {
  box-sizing: border-box;
  margin-top: 1rem;
  gap: 1rem;
  justify-content: center;
  justify-self: center;
  width: fit-content;
  display: flex;
  flex-wrap: wrap;
  height: fit-content;
}

.info {

  align-self: flex-end;
  margin-top: auto;
  height: 1.8rem;
  width: 100%;
  display: flex;
  background-color: #fdff92;
  border-radius: 10px;
  justify-content: center;
  align-items: center;
  padding: 0 0.4rem;
  box-sizing: border-box;
}

.box {
  border-radius: 12px;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  width: 8rem;
  height: 10rem;
  background-color: #fff;

  img {
    width: 100%;
  }

  p {
    margin: 0;
    height: 100%;
    text-wrap-mode: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }
}

.contents {
  overflow-x: hidden;
  overflow-y: scroll;
  justify-content: center;
  align-items: start;
  display: flex;
  flex-wrap: wrap;
  width: 80%;
  margin-top: 2rem;
  height: 85%;
  border-radius: 12px;
  background-color: #afafaf;
}

.sl {
  gap: 1rem;
  justify-content: start;
  align-items: center;
  display: flex;

  p {
    padding: 0 0.5rem;
    transition: 0.5s;
    color: #00000073;
    position: absolute;
    pointer-events: none;
  }
}
@media (width < 480px) {
  .sl{
    flex-direction: column;
    gap: 0;
  }
}
.sl:has(.search:focus-within) {
  p {
    transition: 0.3s;
    opacity: 0;
  }

  .search {
    color: #000000;
    transition: 0.5s;
    background-color: #ffb1ff;
  }
}

.topbar {
  padding: 0 1rem;
  justify-content: space-between;
  width: min(30rem, 90%);
  align-items: center;
  display: flex;
  background-color: #88dfff;
  border-radius: 0 0 20px 20px;
}

.search {
  transition: 0.5s;
  outline: none;
  border: none;
  padding: 0 0.5rem;
  caret-color: transparent;
  background-color: #2ec7ff;
  width: 12rem;
  height: 2rem;
  border-radius: 12px;
}

.list {
  flex-direction: column;
  align-items: center;
  width: 100dvw;
  height: 100dvh;
  overflow-x: hidden;
  display: flex;
}
</style>
