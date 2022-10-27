<script lang="ts" setup>
import { ref, onMounted, computed, watch } from 'vue';
import { LineInterface, mainStore } from '../../store/index';
import { LinesRequest } from '../../requests/linesRequest';
import { storeToRefs } from 'pinia';
import LineComponent from './LineComponent.vue';
import AreasFilter from './AreasFilter.vue';
import autoanimate from '@formkit/auto-animate';

const showOptions = ref<boolean>(false);
const filter = ref<string>("");

const store = mainStore();
const { url } = store;
const { line } = storeToRefs(store);
const lines = ref<LineInterface[]>([]);
const dropdown = ref();
const areasDropdown = ref();
const showAreas = ref<boolean>(false);

const selectedAreas = ref<Array<string>>([]);

watch(showAreas, () => {
    if(showAreas.value) {
        autoanimate(areasDropdown.value)
    }
});

const fetchLines = async () => {
    try {
        const request = new LinesRequest(url);
        const response = await request.request();
        if(response.success) {
            lines.value = response.data;
        }
    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.DEV) throw new Error(exception.message);
        }
    }
}

const setSelected = (payload:string) => {
    selectedAreas.value.push(payload);
}

const unsetSelected = (payload:string) => {
    for(let i = 0; i < selectedAreas.value.length; i++) {
        if(selectedAreas.value[i] === payload) {
            selectedAreas.value.splice(i, 1);
            return;
        }
    }
}

const filtered = computed(() => {
    if(selectedAreas.value.length === 0) {
        return lines.value.filter((element:LineInterface) => {
            const description = element.description.toString().toLowerCase();
            const code = element.code.toString().toLocaleLowerCase();
            const f = filter.value.toString().toLowerCase();

            return description.includes(f) || code.includes(f);
        });
    }
    return lines.value.filter((element:LineInterface) => {
        return selectedAreas.value.some((f:string) => {
            return f === element.areacode;
        });
    });
});

const areas = computed(() => {
    let areas = [];
    for(let i = 0; i < lines.value.length; i++) {
        areas.push(lines.value[i].areacode);
    }

    areas = [...new Set(areas)];
    return areas;
})

const selectedLines = computed(() => {
    let lines = "";
    for(let i = 0; i < line.value.length; i++) {
        lines += `${line.value[i].code},`;
    }
    return lines;
});
const selectAll = () => {
    store.selectAllLines(filtered.value);
}

onMounted(() => {
    fetchLines();
    autoanimate(dropdown.value);
});
</script>
<template>
    <div class="flex flex-col bg-gray-100 shadow border px-3 py-2 rounded dark:bg-gray-800 dark:border-gray-700 relative" ref="dropdown">
        <div class="flex flex-row items-center justify-between group cursor-pointer relative">
            <div class="absolute top-0 bottom-0 left-0 right-0" @click="showOptions = !showOptions"></div>
            <h3 class="font-bold dark:text-gray-200 truncate">
                <i class="fa-solid fa-gears text-blue-500 dark:text-sky-500"></i>
                {{ line.length > 0 ? selectedLines : 'Alege linie' }}
            </h3>
            <i class="fa-solid fa-sort text-blue-500 dark:text-sky-500 group-hover:text-blue-600 dark:group-hover:text-sky-600"></i>
        </div>
        <div class="absolute left-0 right-0 top-12 z-20" v-if="showOptions">
            <div 
                class="flex flex-col w-full px-3 py-2 space-y-2 bg-white shadow border rounded dark:bg-slate-800 dark:border-gray-700 max-h-96 
                overflow-y-scroll scrollbar-thin dark:scrollbar-track-slate-600 dark:scrollbar-thumb-slate-900
                scrollbar-track-blue-200 scrollbar-thumb-blue-400"
                >
                <button @click="store.resetLines()" class="bg-gray-200 px-3 py-2 rounded hover:bg-gray-300 dark:bg-slate-600 dark:text-gray-200 dark:hover:bg-slate-700">
                    <i class="fa-solid fa-xmark"></i>
                    Clear all
                </button>
                <div class="cursor-pointer group" ref="areasDropdown">
                    <div class="flex flex-row items-center justify-between">
                        <h3 class="dark:text-gray-200 font-bold group-hover:dark:text-gray-300 group-hover:text-gray-600">
                            Areas <span class="italic">{{ selectedAreas.join(",") }}</span>
                            <button @click="showAreas = !showAreas"><i class="fa-solid fa-angles-down text-blue-500 group-hover:text-blue-700 dark:text-sky-500 group-hover:dark:text-sky-600"></i></button>
                        </h3>
                        <button class="text-xs font-bold italic dark:text-gray-200 dark:hover:text-gray-200 hover:text-red-500" @click="selectedAreas = []">Clear</button>
                    </div>
                    <div class="flex flex-row flex-wrap">
                        <AreasFilter @selected="setSelected" @unselect="unsetSelected" v-if="showAreas" v-for="area in areas" :area="area" :selectedAreas="selectedAreas"></AreasFilter>
                    </div>
                </div>
                <input v-model="filter" type="text" placeholder="Cautare..." class="px-3 py-2 rounded border bg-gray-200 dark:bg-gray-600 text-gray-200">
                <button class="text-sm font-bold dark:text-gray-200 dark:hover:text-gray-300 hover:text-green-500" @click="selectAll()">
                    <i class="fa-solid fa-check-double text-green-500 dark:text-emerald-500"></i>
                    Select all
                </button>
                <div class="flex flex-col max-h-52 overflow-auto scrollbar-thin dark:scrollbar-track-slate-600 dark:scrollbar-thumb-slate-900
                scrollbar-track-blue-200 scrollbar-thumb-blue-400">
                    <LineComponent v-for="line in filtered" :key="line.id" :line="line"></LineComponent>
                </div>
            </div>
        </div>
    </div>
</template>