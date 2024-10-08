// Store Current User's All Todos
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { apiClient } from '@/utils/axios/axios'


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
    


    // Actions
    const fetchAllTodos = async () => {
        try {
            const response = await apiClient.get('/querytodo', {
                params: {
                    pageSize: '-1',
                    pageNum: '-1'
                }
            })

            allTodos.value = response.data.data.list
            if (response.data.data.list.length > 0) {
                console.log('Get All Todos Done:', allTodos.value)
                loading.value = false
            }
        } catch (error) {
            console.error(error)
        }
    }

    fetchAllTodos()

    const markAsCompleted = (todoid: number) => {
        allTodos.value.filter(item => {
            if (item.ID === todoid) {
                item.isCompleted = true
            }
        })

        // fetchAllTodos()
    }

    const unMarkAsCompleted = (todoid: number) => {
        allTodos.value.filter(item => {
            if (item.ID === todoid) {
                item.isCompleted = false
            }
        })

        // fetchAllTodos()
    }

    const clear = () => {
        allTodos.value = []
        loading.value = true
    }


    return { allTodos, fetchAllTodos, loading, markAsCompleted, unMarkAsCompleted, clear }
},

{
    persist: {
        key: 'alltododata',
        storage: localStorage,
        paths: ['allTodos'],   
    }
})