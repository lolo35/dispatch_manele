<script lang="ts" setup>
import { ref } from 'vue';
import updateTargetQty from '../requests/postUpdateTargetQuantity';
import { mainStore } from '../store';
import { storeToRefs } from 'pinia';
import LoaderComponent from '../components/LoaderComponent.vue';

const store = mainStore();
const { url } = storeToRefs(store);

const isLoading = ref<boolean>(false);
const model = ref<string>("");
const isSuccess = ref<boolean>(false);
const start = ref<string>("");
const end = ref<string>("");

const updateTargetQtyRequest = async () => {
    try {
        isLoading.value = true;
        const request = await updateTargetQty(url.value, start.value, end.value, model.value,);
        if (request.success) {
            isLoading.value = false;
            isSuccess.value = true;
        }
        
    } catch (exception) {
        if(import.meta.env.DEV) {
            if(exception instanceof Error) {
                throw new Error(exception.message);
            }
        }
    } finally {
        isLoading.value = false;
    }
}
</script>

<template>
    <form @submit.prevent="updateTargetQtyRequest()" class="grid grid-cols-3 gap-3 relative">
        <LoaderComponent v-if="isLoading"></LoaderComponent>
        <div class="flex flex-col space-y-2">
            <textarea v-model="model" rows="20" class="px-3 py-2 rounded dark:bg-gray-600 dark:text-gray-200"></textarea>
            <div class="px-3 py-2 bg-green-200 rounded flex flex-row items-center justify-center" v-if="isSuccess">
                <p class="text-green-500 font-bold">
                    <i class="fa-solid fa-circle-check"></i>
                    Dispatch-urile au fost modificate
                </p>
            </div>
            <button type="submit" class="bg-red-500 px-3 py-2 text-white hover:bg-red-600">
                Schimba
            </button>
        </div>
        
        <div class="flex flex-col">
            <label for="dateStart" class="dark:text-gray-200">Start date</label>
            <input v-model="start" type="datetime-local" id="dateStart" class="px-3 py-2 rounded dark:bg-gray-600 dark:text-gray-200" required>
        </div>
        <div class="flex flex-col">
            <label for="dateEnd" class="dark:text-gray-200">End date</label>
            <input v-model="end" type="datetime-local" id="dateEnd" class="px-3 py-2 rounded dark:bg-gray-600 dark:text-gray-200" required>
        </div>
        
    </form>
</template>