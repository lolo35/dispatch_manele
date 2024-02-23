<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { DispatchTypeRequest } from '../requests/dispatchtypeRequest';
import LoaderComponent from '../components/LoaderComponent.vue';
import { mainStore } from '../store';
import { storeToRefs } from 'pinia';
import { DispatchTypeInterface } from '../models/dispatchType';
import { ChangeDispatchTypeRequest } from '../requests/changeDispatchType'

const store = mainStore();
const { url } = storeToRefs(store);

const isLoading = ref<boolean>(false);
const isSuccess = ref<boolean>(false);

const dispatchTypes = ref<Array<DispatchTypeInterface>>([]);
const dispatchType = ref<number>(0);

const model = ref<string>("");
const isClosed = ref<boolean>(false);
const fetchDispatchTypes = async () => {
    try {
        isLoading.value = true;
        const request = new DispatchTypeRequest(url.value);
        const response = await request.request();
        if (response.success) {
            dispatchTypes.value = response.data;
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

const changeDispatchType = async () => {
    try {
        isLoading.value = true;
        if(dispatchType.value === null) {
            throw new Error("Nu ai selectat un dispatch type");
        }
        const request = new ChangeDispatchTypeRequest(url.value, model.value, dispatchType.value, isClosed.value);
        const response = await request.send();
        if (response.success) {
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
onMounted(async () => {
    await fetchDispatchTypes();
})
</script>

<template>
    <form @submit.prevent="changeDispatchType()" class="grid grid-cols-3 gap-3 relative">
        <LoaderComponent v-if="isLoading"></LoaderComponent>
        <div class="flex flex-col space-y-2">
            <textarea v-model="model" name="" id="" rows="20" class="px-3 py-2 rounded dark:bg-gray-600 dark:text-gray-200"></textarea>
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
        <div class="flex flex-col max-h-max">
            <label for="dispatch_status" class="text-sm font-bold dark:text-gray-200">Dispatch type</label>
            <select v-model="dispatchType" id="dispatch_status" class="px-3 py-2 rounded border" required>
                <option :value="dispatch_type.id" v-for="dispatch_type in dispatchTypes" :key="dispatch_type.id">{{ dispatch_type.description }}</option>
            </select>
            <div class="flex flex-row space-x-3 items-center">
                <label for="close" class="dark:text-gray-200">Si inchide?</label>
                <input type="checkbox" id="close" v-model="isClosed">
            </div>
        </div>
    </form>
</template>