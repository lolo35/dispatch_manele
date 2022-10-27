<script lang="ts" setup>
import { computed } from 'vue';
import {mainStore} from '../../store/index';
import { storeToRefs } from 'pinia';

const store = mainStore();
const { randStart, randEnd } = storeToRefs(store);

const sStart = computed({
    get() {
        return randStart.value;
    },
    set(value:number) {
        store.setRandStart(value)
    }
})

const sEnd = computed({
    get() {
        return randEnd.value;
    },
    set(value:number) {
        store.setRandEnd(value);
    }
})

</script>
<template>
    <div class="flex flex-row bg-gray-100 shadow border rounded dark:bg-gray-800 dark:border-gray-700 relative px-3 py-2 col-span-2 space-x-3">
        <div class="flex flex-col w-full">
            <label for="randstart" class="dark:text-gray-200 text-sm font-bold">Random start</label>
            <input v-model="sStart" type="number" class="px-3 py-2 rounded w-full" placeholder="Random start">
        </div>
        <div class="flex flex-col w-full">
            <label for="randend" class="dark:text-gray-200 text-sm font-bold">Random end</label>
            <input v-model="sEnd" type="number" class="px-3 py-2 rounded w-full">
        </div>
    </div>
</template>