<template>
    <div>
        <h1>Add Todo</h1>   

        <form @submit.prevent="AddTodo">
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
            <div>
                <label class="flex items-center">
                    <input v-model="form.IsCompleted" type="checkbox" class="mr-2" />
                    <span>Is Completed</span>
                </label>
            </div>
            <button type="submit" class="bg-blue-500 text-white p-2">Add Todo</button>
        </form>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { useUserStore } from '@/stores/userStore';

const userStore = useUserStore()

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
    const authToken = userStore.userInfo.Token
    if (!authToken) {
        console.error('No auth token found')
        return
    }

    console.log(authToken)

    try {
        const response = await axios.post(
            'http://127.0.0.1:7879/api/add',
            form.value, // 传递表单数据
            {
                headers: {
                    'Authorization': `Bearer ${authToken}` // 添加 token 到请求头中
                }  
            }
        )

        console.log('Response Data:', response.data)
    } catch (error) {
        console.error('Error adding todo:', error)
    }
}

</script>

<style scoped></style>