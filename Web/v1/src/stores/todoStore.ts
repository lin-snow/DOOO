// Store Current User's All Todos
import { defineStore } from 'pinia'
import { ref } from 'vue'
import axios from 'axios'


interface TodoModel {
    ID: number
    title: string
    description: string
    isCompleted: boolean

    userId: number
    categoryId: number
}

export const useTodoStore = defineStore('allTodos', () => {
    // State
    const allTodos = ref<TodoModel[]>([])
    const loading = ref(true)
    const token = localStorage.getItem('authToken')


    // Actions
    const fetchAllTodos = async () => {
        try {
            const response = await axios.get('http://127.0.0.1:7879/api/querytodo', {
                headers: {
                    'Authorization': `Bearer ${ token }` // 添加 token 到请求头中
                },
                params: {
                    pageSize: '-1',
                    pageNum: '-1'
                }
            })

            allTodos.value = response.data.data.list
            if (response.data.data.list.length > 0) {
                console.log('All Todos:', allTodos.value)
                loading.value = false
            }
        } catch (error) {
            console.error(error)
        }
    }

    // fetchAllTodos()

    return { allTodos, fetchAllTodos, loading }
})