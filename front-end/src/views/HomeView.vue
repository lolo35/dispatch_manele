<script lang="ts" setup>
import DispatchTypeSelection from '../components/Dispatch/DispatchTypeSelection.vue';
import SelectLine from '../components/Lines/SelectLine.vue';
import ResourseComponent from '../components/Dispatch/ResourseComponent.vue';
import DescriptionComponent from '../components/Dispatch/DescriptionComponent.vue';
import TradecodeComponent from '../components/Dispatch/TradecodeComponent.vue';
import { AddDispatchRequest } from '../requests/addDispatchRequest';
import LoaderComponent from '../components/LoaderComponent.vue';
import RandomInputs from '../components/Dispatch/RandomInputs.vue';
import { ref } from 'vue';

import { mainStore } from '../store/index';
import { storeToRefs } from 'pinia';

const store = mainStore();
const { url } = store;
const isLoading = ref<boolean>(false);

const { dispatchType, line, description, tradecode, resourseid, randStart, randEnd, descriptionIsRandom } = storeToRefs(store);

const addDispatch = async () => {
    try {
        isLoading.value = true;
        if(dispatchType.value) {
            const request = new AddDispatchRequest(url, dispatchType.value.code, description.value, tradecode.value, line.value, resourseid.value, randStart.value, randEnd.value, descriptionIsRandom.value);

            const response = await request.request();
            if(response.success) {
                console.log(`success`);
            }
        }
    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.DEV) throw new Error(exception.message);
        }
    } finally {
        isLoading.value = false;
    }
}
</script>
<template>
    <div class="grid grid-cols-3 gap-2 mt-2 relative">
        <LoaderComponent v-if="isLoading"></LoaderComponent>
        <DispatchTypeSelection></DispatchTypeSelection>
        <SelectLine></SelectLine>
        <ResourseComponent></ResourseComponent>
        <DescriptionComponent></DescriptionComponent>
        <!-- <div id="placeholder"></div> -->
        <TradecodeComponent></TradecodeComponent>
        <RandomInputs></RandomInputs>
        <button @click="addDispatch()" class="col-span-3 bg-blue-500 text-white dark:text-gray-200 px-3 py-2 rounded hover:bg-blue-600 dark:bg-sky-500 dark:hover:bg-sky-600">Submit</button>
    </div>
</template>