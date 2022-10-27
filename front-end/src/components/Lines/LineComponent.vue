<script lang="ts" setup>
import { defineProps, PropType, computed } from 'vue';
import { LineInterface } from '../../store/index';
import { mainStore } from '../../store/index';
import { storeToRefs } from 'pinia';

const store = mainStore();
const { line } = storeToRefs(store);

const props = defineProps({
    line: {
        type: Object as PropType<LineInterface>,
        required: true,
    }
});

const isSelected = computed(() => {
    let selected = false;
    for(let i = 0; i < line.value.length; i++) {
        if(line.value[i].id === props.line.id) {
            selected = true;
            break;
        }
    }

    return selected;
});

const setLine = () => {
    console.log(`settings line`);
    if(!isSelected.value) {
        store.setLine(props.line);
        return;
    }
    store.removeLine(props.line.id)
}
</script>

<template>
    <div class="flex flex-row px-2 py-1 dark:hover:bg-sky-500 cursor-pointer" @click="setLine()">
        <p class="truncate dark:text-gray-200">
            <i class="fa-solid fa-circle-check text-green-500 mr-2" v-if="isSelected"></i>
            <span class="font-bold">{{ props.line.code }}</span>
            - 
            <span class="italic">{{ props.line.description }}</span>
        </p>
    </div>
</template>