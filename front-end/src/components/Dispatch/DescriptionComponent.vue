<script lang="ts" setup>
import { computed } from 'vue';
import { mainStore } from '../../store/index';
import { storeToRefs } from 'pinia';

const store = mainStore();
const { description, descriptionIsRandom } = storeToRefs(store);
const isRandom = computed({
    get() {
        return descriptionIsRandom.value;
    },
    set(value:boolean) {
        store.setDescriptionIsRandom(value);
    }
});

const model = computed({
    get() {
        return description.value;
    },
    set(value:string) {
        store.setDescription(value);
    }
})
</script>
<template>
    <div class="flex flex-col col-span-2">
        <textarea v-model="model" class="rounded border dark:bg-gray-800 dark:border-gray-700 dark:text-gray-200 px-3 py-2" placeholder="Descriere..." id="" rows="3"></textarea>
        <div class="flex flex-row px-3 py-2 items-center space-x-3">
            <input type="checkbox" id="isRandom" v-model="isRandom">
            <label for="isRandom" class="dark:text-gray-200 font-semibold italic">Random description</label>
        </div>
    </div>
</template>