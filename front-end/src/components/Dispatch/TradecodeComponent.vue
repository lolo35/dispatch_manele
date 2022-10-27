<script lang="ts" setup>
import { TradecodeRequest } from '../../requests/tradecodesRequest';
import { mainStore } from '../../store/index';
import { onMounted, ref, computed } from 'vue';
import autoanimate from '@formkit/auto-animate';
import { storeToRefs } from 'pinia';

interface Tradecodes {
    id: number,
    code: string,
    description: string,
}

const store = mainStore();
const { url } = store;
const { tradecode } = storeToRefs(store);

const tradecodes = ref<Tradecodes[]>([]);
const showOptions = ref<boolean>(false);
const filter = ref<string>("");
const dropdown = ref();

const fetchTradecodes = async () => {
    try {
        const request = new TradecodeRequest(url);
        const response = await request.request();

        if(response.success) {
            tradecodes.value = response.data;
        }
    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.DEV) throw new Error(exception.message);
        }
    }
}

const filtered = computed(() => {
    return tradecodes.value.filter((element:Tradecodes) => {
        const code = element.code.toString().toLocaleLowerCase();
        const description = element.description.toString().toLowerCase();
        const f = filter.value.toString().toLocaleLowerCase();

        return code.includes(f) || description.includes(f);
    })
})

const setTradecode = (code:string) => {
    store.setTradecode(code);
    showOptions.value = false;
}

onMounted(() => {
    fetchTradecodes();
    autoanimate(dropdown.value);
})
</script>
<template>
    <div class="flex flex-col bg-gray-100 shadow border rounded dark:bg-gray-800 dark:border-gray-700 relative px-3 py-2 max-h-10" ref="dropdown">
        <div class="flex flex-row items-center justify-between group cursor-pointer relative" >
            <div class="absolute top-0 bottom-0 left-0 right-0" @click="showOptions = !showOptions"></div>
            <h3 class="font-bold dark:text-gray-200">
                <i class="fa-solid fa-globe text-blue-500 dark:text-sky-500"></i>
                {{ tradecode.length > 0 ? tradecode : 'Alege tradecode' }}
            </h3>
            <i class="fa-solid fa-sort text-blue-500 dark:text-sky-500 group-hover:text-blue-600 dark:group-hover:text-sky-600"></i>
        </div>
        <div class="absolute left-0 right-0 top-12" v-if="showOptions">
            <div 
                class="flex flex-col w-full px-3 py-2 bg-white shadow border rounded dark:bg-slate-800 dark:border-gray-700 max-h-96 
                overflow-y-scroll scrollbar-thin dark:scrollbar-track-slate-600 dark:scrollbar-thumb-slate-900
                scrollbar-track-blue-200 scrollbar-thumb-blue-400"
                >
                <input v-model="filter" type="text" placeholder="Cautare..." class="px-3 py-2 rounded border bg-gray-200 dark:bg-gray-600 text-gray-200">
                <div @click="setTradecode(tradecode.code)" class="flex flex-row px-2 py-1 dark:hover:bg-sky-500 cursor-pointer" v-for="tradecode in filtered" :key="tradecode.id">
                    <p class="truncate dark:text-gray-200"><span class="font-bold">{{ tradecode.code }}</span> - <span class="italic">{{ tradecode.description }}</span></p>
                </div>
            </div>
        </div>
    </div>
</template>