<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import SidebarComponent from './components/Navigation/SidebarComponent.vue';
import { mainStore } from './store/index';
import { onMounted } from 'vue';

const store = mainStore();

onMounted(() => {
  store.setUrl("http://tro-app11:8004/api/v1/")
  if(import.meta.env.DEV) {
    store.setUrl("http://localhost:8000/api/v1/")
    document.body.classList.add("debug-screens")
  }
})
</script>

<template>
  <main class="flex w-full h-screen">
    <SidebarComponent></SidebarComponent>
    <section class="w-full max-h-full overflow-y-scroll relative">
      <router-view class="px-3 py-2"></router-view>
    </section>
  </main>
</template>

<style scoped>
</style>
