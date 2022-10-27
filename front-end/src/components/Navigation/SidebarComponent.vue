<script lang="ts" setup>
import { computed } from '@vue/reactivity';

import { useRoute } from 'vue-router';

const route = useRoute();

const matched = computed(() => {
    if(route.matched[0]) {
        return route.matched[0].name;
    }
    return false;
});

</script>
<template>
    <div class="flex flex-col w-20 bg-gray-100 shadow border space-y-3 justify-between dark:bg-gray-900 dark:border-gray-800">
        <div class="flex flex-col space-y-10 items-center">
            <img src="../../../public/vite.svg" alt="">
            <router-link to="/" title="Dispatch" :active-class="'dark:text-sky-500 dark:hover:text-sky-600 text-blue-500 hover:text-blue-700'" :class="{'text-gray-500 hover:text-gray-600': matched !== 'Home'}">
                <i class="fa-regular fa-paper-plane fa-2x"></i>
            </router-link>
            <router-link to="/delete" title="Delete dispatch" :active-class="'dark:text-sky-500 dark:hover:text-sky-600 text-blue-500 hover:text-blue-700'" :class="{'text-gray-500 hover:text-gray-600': matched !== 'Delete'}">
                <i class="fa-solid fa-trash-can fa-2x"></i>
            </router-link>
            <router-link to="/descriptions" title="Dispatch description" :active-class="'dark:text-sky-500 dark:hover:text-sky-600 text-blue-500 hover:text-blue-700'" :class="{'text-gray-500 hover:text-gray-600': matched !== 'Delete'}">
                <i class="fa-solid fa-comment fa-2x"></i>
            </router-link>
        </div>
    </div>
</template>