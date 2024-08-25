<template>
    <div>
        <h1 class="text-center font-serif text-lg subpixel-antialiased text-amber-800 font-bold underline decoration-wavy underline-offset-8 my-2">Edit Todo</h1>
        <div class="m-4">
            <form @submit.prevent="EditTodo">
                <div class="my-2">
                    <span class="font-serif text-md subpixel-antialiased text-black font-bold underline underline-offset-8 block h-8">Title:</span>
                    <textarea v-model="form.Title" ref="titlearea" type="text" class="w-full border rounded-lg border-neutral-400 outline-none bg-amber-50 p-1 shadow-inner my-2 overflow-y-hidden resize-none focus:border-amber-800 focus:border-2 focus:bg-white" placeholder="Title"></textarea>
                </div>
                <div class="mb-2">
                    <span class="font-serif text-md subpixel-antialiased text-black font-bold underline underline-offset-8 block h-8">Description:</span>
                    <textarea v-model="form.Description" ref="desarea" type="text" class="w-full border rounded-lg border-neutral-400 outline-none bg-amber-50 p-1 shadow-inner my-2 overflow-y-hidden resize-none focus:border-amber-800 focus:border-2 focus:bg-white" placeholder="Description"></textarea>
                </div>
                <div class="grid justify-items-center my-2">
                    <button type="submit" class="border rounded-xl border-amber-700 shadow-lg p-2 w-20 text-black font-medium hover:bg-rose-200"> Update </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router'
import { useTodoStore } from '@/stores/todoStore'
import { useUserStore } from '@/stores/userStore'
import { apiClient } from '@/utils/axios/axios'


interface TodoForm {
    ID: number
    UserID: number
    Title: string
    Description: string
    DueDate: string
    IsCompleted: boolean
    CategoryID: number
}


const route = useRoute();
const todoid = ref(Number(route.params.todoid))
const todoStore = useTodoStore()
const userStore = useUserStore()
const currentTodo = todoStore.allTodos.filter(item => item.ID === todoid.value)[0]

const form = ref<TodoForm>({
    ID: todoid.value,
    UserID: userStore.userInfo.UserID,
    Title: currentTodo.title,
    Description: currentTodo.description,
    DueDate: '2024-08-18T08:00:00Z',
    IsCompleted: currentTodo.isCompleted,
    CategoryID: currentTodo.categoryId
})

console.log(todoid.value)

const EditTodo = async () => {
    try {
        const response = await apiClient.put('/updatetodo', 
        form.value)
        
        console.log(response.data)

        // todoStore.fetchAllTodos()
    } catch (error) {
        console.error(error)
    }
}


const titlearea = ref<HTMLTextAreaElement | null>(null);
const desarea = ref<HTMLTextAreaElement | null>(null);

const resizeTextarea = (autosizeTextarea: any) => {
    if (autosizeTextarea.value) {
    autosizeTextarea.value.style.height = '2.3rem'; // 重置高度
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