<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { ChangeDispatchStatusRequest } from '../requests/changeDispatchStatus';
import LoaderComponent from '../components/LoaderComponent.vue'
import { DispatchSysStatusesInterface } from '../models/dispatchsystemstatuse';
import { DispatchSystemStatusesRequest } from '../requests/dispatchsystemstatuses';


const model = ref<string>("");
const isLoading = ref<boolean>(false);
const isSuccess = ref<boolean>(false)
const dispatchStatuses = ref<Array<DispatchSysStatusesInterface>>([]);
const dispatchStatus = ref<number>();

const fetchdispatchStatuse = async () => {
    try {
        isLoading.value = true;
        const request = new DispatchSystemStatusesRequest();
        const response = await request.req();
        if(response.success) {
            dispatchStatuses.value = response.statuses;
        }
    } catch (exception) {
        if(import.meta.env.DEV) {
            if(exception instanceof Error){
                throw new Error(exception.message);
            }
        }
    } finally {
        isLoading.value = false;
    }
}

const changeDispatchStatus = async () => {
    try {
        isLoading.value = true;
        if(!dispatchStatus.value) {
            return
        }
        const request = new ChangeDispatchStatusRequest(model.value, dispatchStatus.value);
        const response = await request.req();
    } catch (exception) {
        if(import.meta.env.DEV) {
            if(exception instanceof Error) {
                throw new Error(exception.message)
            }
        }
    } finally {
        isLoading.value = false;
    }
}
onMounted(async () => {
    await fetchdispatchStatuse();
})
</script>

<template>
    
    <form @submit.prevent="changeDispatchStatus()" class="grid grid-cols-3 gap-3 relative">
        <LoaderComponent v-if="isLoading"></LoaderComponent>
        <div class="flex flex-col space-y-2">
            <textarea v-model="model" name="" id="" rows="20" class="px-3 py-2 rounded dark:bg-gray-600 dark:text-gray-200"></textarea>
            <div class="px-3 py-2 bg-green-200 rounded flex flex-row items-center justify-center" v-if="isSuccess">
                <p class="text-green-500 font-bold">
                    <i class="fa-solid fa-circle-check"></i>
                    Dispatch-urile au fost inchise
                </p>
            </div>
            <button type="submit" class="bg-red-500 px-3 py-2 text-white hover:bg-red-600">
                Schimba
            </button>
        </div>
        <div class="flex flex-col max-h-max">
            <label for="dispatch_status" class="text-sm font-bold dark:text-gray-200">Dispatch status</label>
            <select v-model="dispatchStatus" id="dispatch_status" class="px-3 py-2 rounded border" required>
                <option :value="status.id" v-for="status in dispatchStatuses" :key="status.id">{{ status.description }}</option>
            </select>
        </div>
    </form>
    
</template>