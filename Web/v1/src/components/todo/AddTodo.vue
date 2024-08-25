<template>
    <div>
        <h1 class="text-center font-mono text-lg subpixel-antialiased text-amber-800 font-bold underline decoration-wavy underline-offset-8">Add Todo</h1>   
        <form @submit.prevent="AddTodo" class="font-mono">
            <div class="my-2">
                <textarea v-model="form.Title" ref="titlearea" type="text" class="border rounded-lg border-neutral-400 w-44 p-2 shadow-inner my-2 overflow-y-hidden resize-none focus:border-amber-800 outline-none focus:border-2" placeholder="Title"></textarea>
            </div>
            <div class="mb-2">
                <textarea v-model="form.Description" ref="desarea" type="text" class="border rounded-lg border-neutral-400 w-44  p-2 shadow-inner my-2 overflow-y-hidden resize-none focus:border-amber-800 outline-none focus:border-2" placeholder="Description"></textarea>
            </div>
            <div class="grid justify-items-center">
                <button type="submit" class="border-2 rounded-xl border-amber-700 shadow-lg p-2 w-24 text-black font-medium hover:bg-rose-300"> Add </button>
            </div>
            
        </form>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { apiClient } from '@/utils/axios/axios'
import { useUserStore } from '@/stores/userStore';
import { useRouter } from 'vue-router'


const userStore = useUserStore()
const router = useRouter()

interface TodoForm {
    UserID: number
    Title: string
    Description: string
    DueDate: string
    IsCompleted: boolean
    CategoryID: number
}

const form = ref<TodoForm>({
    UserID: userStore.userInfo.UserID,
    Title: '',
    Description: '',
    DueDate: '2024-08-18T08:00:00Z',
    IsCompleted: false,
    CategoryID: 0
})

const AddTodo = async () => {
    try {
        const response = await apiClient.post(
            '/add',
            form.value // 传递表单数据
    )

        console.log('Response Data:', response.data)
        
        router.push('/')
    } catch (error) {
        console.error('Error adding todo:', error)
    }
}

const titlearea = ref<HTMLTextAreaElement | null>(null);
const desarea = ref<HTMLTextAreaElement | null>(null);

const resizeTextarea = (autosizeTextarea: any) => {
    if (autosizeTextarea.value) {
    autosizeTextarea.value.style.height = '2.5rem'; // 重置高度
    autosizeTextarea.value.style.height = autosizeTextarea.value.scrollHeight + 'px'; // 根据内容设置高度
    }
};


onMounted(() => {
  // 初始调整高度
    resizeTextarea(titlearea);
    resizeTextarea(desarea);
});

// 监听表单数据变化，以便在数据变化时调整高度
watch(() => form.value.Title, () => {
    resizeTextarea(titlearea);
});

watch(() => form.value.Description, () => {
    resizeTextarea(desarea);
});
</script>
<style scoped></style>