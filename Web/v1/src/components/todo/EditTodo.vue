<template>
    <div class="w-full">
        <h1>Edit Todo</h1>

        <form @submit.prevent="EditTodo">
            <div class="my-2">
                <span>Title:</span>
                <br>
                <textarea v-model="form.Title" ref="titlearea" type="text" class="border rounded-lg border-neutral-400  p-2 shadow-inner my-2 overflow-y-hidden resize-none focus:border-amber-800 outline-none focus:border-2" placeholder="Title"></textarea>
            </div>
            <div class="mb-2">
                <span>Description:</span>
                <br>
                <textarea v-model="form.Description" ref="desarea" type="text" class="border rounded-lg border-neutral-400  p-2 shadow-inner my-2 overflow-y-hidden resize-none focus:border-amber-800 outline-none focus:border-2" placeholder="Description"></textarea>
            </div>
            <button type="submit" class="border-2 rounded-xl border-amber-700 shadow-lg p-2 w-20 text-black font-medium hover:bg-rose-200"> Update </button>
        </form>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router'
import { useTodoStore } from '@/stores/todoStore'
import { useUserStore } from '@/stores/userStore'
import axios from 'axios'


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
        const response = await axios.put('http://localhost:7879/api/updatetodo', 
        form.value,
        {
            headers: {
                Authorization: `Bearer ${ userStore.userInfo.Token }`
            }
        })
        // console.log(form.value)
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