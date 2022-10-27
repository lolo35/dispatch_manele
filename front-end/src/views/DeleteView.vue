<script lang="ts" setup>
import { ref } from 'vue';
import { mainStore } from '../store/index';
import { DeleteDispatchRequest } from '../requests/deleteDispatchRequest';
import LoaderComponent from '../components/LoaderComponent.vue'

const store = mainStore();
const { url } = store;

const model = ref<string>("");
const isLoading = ref<boolean>(false);
const isSuccess = ref<boolean>(false)

const deleteDispatches = async () => {
    try {
        isLoading.value = true;
        const deleteDispatch = new DeleteDispatchRequest(url, model.value);
        const req = await deleteDispatch.request();

    } catch (exception) {
        if(exception instanceof Error) {
            if(import.meta.env.DEV) throw new Error(exception.message)
        }
    } finally {
        isLoading.value = false;
    }
}
</script>

<template>
    <div class="grid grid-cols-3">
        <form @submit.prevent="deleteDispatches()" class="flex flex-col space-y-2 relative">
            <LoaderComponent v-if="isLoading"></LoaderComponent>
            <textarea v-model="model" name="" id="" rows="20" class="px-3 py-2 rounded dark:bg-gray-600 dark:text-gray-200"></textarea>
            <div class="px-3 py-2 bg-green-200 rounded flex flex-row items-center justify-center" v-if="isSuccess">
                <p class="text-green-500 font-bold">
                    <i class="fa-solid fa-circle-check"></i>
                    Dispatch-urile au fost sterse
                </p>
            </div>
            <button type="submit" class="bg-red-500 px-3 py-2 text-white hover:bg-red-600">
                <i class="fa-solid fa-trash-can"></i>
                Sterge
            </button>
        </form>
    </div>
</template>