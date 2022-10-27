<script lang="ts" setup>
import DispatchTypeSelectionVue from '../components/Dispatch/DispatchTypeSelection.vue';
import { FetchDispatchtypeCount } from '../requests/fetchDispatchTypeCount';
import { ImportDispatchtypeDescriptions } from '../requests/importDispatchtypeDescriptions';
import { ref } from 'vue';
import { mainStore } from '../store';
import { storeToRefs } from 'pinia';
import LoaderComponent from '../components/LoaderComponent.vue';

const store = mainStore();

const { url } = store;
const { dispatchType } = storeToRefs(store);
const count = ref<number>(0);
const isLoaded = ref<boolean>(false);
const isLoading = ref<boolean>(false);

const fetchCount = async () => {
    const t0 = performance.now();
    isLoaded.value = false;
    isLoading.value = true;
    try {
        if(!dispatchType.value) {
            return
        }
        const request = new FetchDispatchtypeCount(url, dispatchType.value.code);
        const response = await request.req();
        if(response.success) {
            count.value = response.count;
        }
    } catch (exception) {
        if(import.meta.env.DEV) {
            if(exception instanceof Error) {
                throw new Error(exception.message);
            }
        }
    } finally {
        const t1 = performance.now();
        isLoaded.value = true;
        isLoading.value = false;
        console.log(`request took ${t1 - t0} miliseconds`);
    }
}

const importDescriptions = async () => {
    const t0 = performance.now()
    try {
        if(!dispatchType.value) {
            return
        }
        isLoading.value = true;
        const request = new ImportDispatchtypeDescriptions(url, dispatchType.value.code);
        const response = await request.req();
        if (response.success) {
            fetchCount();
        }
    } catch (exception) {
        if(import.meta.env.DEV) {
            if(exception instanceof Error) {
                throw new Error(exception.message);
            }
        }
    } finally {
        const t1 = performance.now();
        isLoading.value = false;
        console.log(`request took ${t1 - t0} miliseconds`);
    }
}
</script>
<template>
    <div class="grid grid-cols-3 gap-3 relative">
        <LoaderComponent v-if="isLoading"></LoaderComponent>
        <div class="flex flex-row w-full">
            <DispatchTypeSelectionVue></DispatchTypeSelectionVue>
            <button @click="fetchCount()" class="bg-blue-500 text-white px-3 py-2 rounded-r hover:bg-blue-700"><i class="fa-solid fa-magnifying-glass"></i></button>
        </div>
        <div v-if="isLoaded" class="flex flex-row bg-gray-100 shadow border px-3 py-2 rounded dark:bg-gray-800 dark:border-gray-700 relative w-full">
            <p class="dark:text-gray-200 font-semibold italic">Aveti {{ count }} numar de descrieri in baza de date</p>
        </div>
        <div v-if="isLoaded" class="w-full" @click="importDescriptions()">
            <button class="bg-red-500 text-white px-3 py-2 rounded w-full hover:bg-red-700">Import</button>
        </div>
    </div>
</template>