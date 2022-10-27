<script lang="ts" setup>
import { onMounted, ref, computed } from 'vue';
import { mainStore, DispatchTypeInterface } from '../../store/index';
import { DispatchTypeRequest } from '../../requests/dispatchtypeRequest';
import { storeToRefs } from 'pinia';
import autoanimate from '@formkit/auto-animate';

const store = mainStore();
const { url } = store;
const { dispatchType } = storeToRefs(store);
const showOptions = ref<boolean>(false);
const dispatchTypes = ref<DispatchTypeInterface[]>([]);
const filter = ref<string>("");
const dropdown = ref();

const fetchtypes = async () => {
    try {
        const request = new DispatchTypeRequest(url);
        const response = await request.request();
        if(response.success) {
            dispatchTypes.value = response.data;
        }
    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.DEV) throw new Error(exception.message);
        }
    }
}

const filtered = computed(() => {
    return dispatchTypes.value.filter((element:DispatchTypeInterface) => {
        const description = element.description.toString().toLowerCase();
        const f = filter.value.toString().toLowerCase();
        const code = element.code.toString().toLowerCase();

        return description.includes(f) || code.includes(f);
    });
})

const setDispatchType = (dispatchtype: DispatchTypeInterface) => {
    try {
        store.setDispatchType(dispatchtype);
        showOptions.value = false;
    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.DEV) throw new Error(exception.message)
        }
    }
}

onMounted(() => {
    fetchtypes();
    autoanimate(dropdown.value);
});
</script>
<template>
    <div class="flex flex-col bg-gray-100 shadow border px-3 py-2 rounded dark:bg-gray-800 dark:border-gray-700 relative w-full" ref="dropdown">
        <div class="flex flex-row items-center justify-between group cursor-pointer relative" >
            <div class="absolute top-0 bottom-0 left-0 right-0" @click="showOptions = !showOptions"></div>
            <h3 class="font-bold dark:text-gray-200">
                <i class="fa-regular fa-paper-plane text-blue-500 dark:text-sky-500"></i>
                {{ dispatchType ? `${dispatchType.code} - ${dispatchType.description}` : 'Alege tipul de dispatch' }}
            </h3>
            <i class="fa-solid fa-sort text-blue-500 dark:text-sky-500 group-hover:text-blue-600 dark:group-hover:text-sky-600"></i>
        </div>
        <div class="absolute left-0 right-0 top-12 z-20" v-if="showOptions">
            <div 
                class="flex flex-col w-full px-3 py-2 bg-white shadow border rounded dark:bg-slate-800 dark:border-gray-700 max-h-96 
                overflow-y-scroll scrollbar-thin dark:scrollbar-track-slate-600 dark:scrollbar-thumb-slate-900
                scrollbar-track-blue-200 scrollbar-thumb-blue-400"
                >
                <input v-model="filter" type="text" placeholder="Cautare..." class="px-3 py-2 rounded border bg-gray-200 dark:bg-gray-600 text-gray-200">
                <div @click="setDispatchType(dispatchtype)" class="flex flex-row px-2 py-1 dark:hover:bg-sky-500 cursor-pointer" v-for="dispatchtype in filtered" :key="dispatchtype.id">
                    <p class="truncate dark:text-gray-200"><span class="font-bold">{{ dispatchtype.code }}</span> - <span class="italic">{{ dispatchtype.description }}</span></p>
                </div>
            </div>
        </div>
    </div>
</template>