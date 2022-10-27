<script lang="ts" setup>
import { defineProps, PropType, defineEmits, computed } from 'vue';

const emit = defineEmits(['selected', 'unselect']);
const props = defineProps({
    area: {
        type: String as PropType<string>,
        required: true,
    },
    selectedAreas: {
        type: Array as PropType<Array<string>>,
        required: true,
    }
});
const isSelected = computed(() => {
    let condition = false;
    for(let i = 0; i < props.selectedAreas.length; i++) {
        if(props.selectedAreas[i] === props.area) {
            condition = true;
            break;
        }
    }
    return condition;
});
const toggleArea = () => {
    if(isSelected.value) {
        emit("unselect", props.area);
        return;
    }
    emit("selected", props.area);
}
</script>
<template>
    <button @click="toggleArea();" 
        class="px-3 py-2 rounded mr-2 mb-2 dark:text-gray-200 "
        :class="{'bg-green-500 dark:bg-emerald-500': isSelected, 'dark:bg-slate-500 dark:hover:bg-slate-600': !isSelected}"
        >{{ props.area }}</button>
</template>