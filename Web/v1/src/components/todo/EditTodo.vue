<template>
    <div class="w-full">
        <h1>Edit Todo</h1>

        <form @submit.prevent="EditTodo">
            <div class="my-2">
                <input v-model="form.Title" type="text" class="border border-gray-300 p-2" placeholder="Title" />
            </div>
            <div class="mb-2">
                <input v-model="form.Description" type="text" class="border border-gray-300 p-2" placeholder="Description" />
            </div>
            <!-- <div class="mb-2">
                <input v-model="form.Category" type="text" class="border border-gray-300 p-2" placeholder="Category" />
            </div> -->
            <!-- <div class="mb-2">
                <input v-model="form.DueDate" type="date" class="border border-gray-300 p-2" placeholder="Due Date" />
            </div> -->

            <button type="submit" class="bg-blue-500 text-white p-2">Update Todo</button>
        </form>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
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

</script>