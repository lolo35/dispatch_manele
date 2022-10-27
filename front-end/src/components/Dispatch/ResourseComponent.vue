<script lang="ts" setup>
import { computed } from 'vue';
import { mainStore } from '../../store/index';
import { storeToRefs } from 'pinia';

const store = mainStore();
const { resourseid } = storeToRefs(store);

const model = computed({
    get() {
        return resourseid.value;
    },
    set(value:string) {
        store.setResourseid(value);
    }
})
</script>
<template>
    <div class="flex flex-col bg-gray-100 shadow border rounded dark:bg-gray-800 dark:border-gray-700 relative">
        <input v-model="model" type="text" class="px-3 py-2 rounded border dark:bg-gray-800 dark:border-gray-700 dark:text-gray-200" placeholder="Resourse login id ex: TRO07513">
    </div>
</template>