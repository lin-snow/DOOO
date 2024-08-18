<template>
    <div v-if="todoStore.loading" class="text-center">Loading...</div>
    <el-scrollbar v-else>
        <div v-for="todo in uncompletedTodos" :key="todo.ID" class="h-18">
            <el-row class="alignment-container"> 
                <el-col :span="2" class="mt-4">
                    <!-- Mark as completed -->
                    <button @click="markAsCompleted(todo.ID)" class="w-6 h-6 text-amber-900 border-2 rounded-full border-amber-900 ">
                        <CheckIcon class="opacity-0 hover:opacity-100 transition-opacity duration-200"/>
                    </button>
                </el-col>
                <!-- Show Todo -->
                <el-col :span="20">
                    <el-row class="group">
                        <el-col :span="21" class="my-2 mx-2 rounded-lg p-2 border-2 shadow-sm border-amber-900 ">
                            <div>
                                <h3 class="text-amber-950 font-sans subpixel-antialiased truncate ... underline decoration-wavy underline-offset-3 h-auto">
                                    {{ todo.title }}
                                </h3>
                                <p class="text-slate-400 font-light truncate ...">
                                {{ todo.description }}
                                </p>
                            </div>
                        </el-col>
                        <el-col :span="1" class="mr">            
                                <button @click="TodoDetail(todo.ID)" class="w-6 h-6 mt-4 border-2 rounded-full border-amber-900 opacity-0 group-hover:opacity-100 transition-opacity duration-200">
                                    <ExpandIcon />
                                </button>
                        </el-col>
                    </el-row>
                </el-col>
            </el-row>
        </div>

        <hr class="border-2 border-dotted">

        <div v-for="todo in completedTodos" :key="todo.ID" class="h-18">
            <el-row class="group alignment-container"> 
                <el-col :span="2" class="mt-4">
                    <!-- Withdraw mark -->
                    <button @click="withDrawMark(todo.ID)" class="w-6 h-6 text-amber-900 border-2 rounded-full border-amber-900 ">
                        <WithdrawIcon class="opacity-0 hover:opacity-100 transition-opacity duration-200"/>
                    </button>
                </el-col>
                <!-- Show Todo -->
                <el-col :span="20">
                    <el-row>
                        <el-col :span="21" class="my-2 mx-2 rounded-lg p-2 border-2 shadow-sm border-amber-800 border-dashed">
                            <div>
                                <h3 class="text-gray-400 line-through truncate ... font-sans subpixel-antialiased h-8">
                                    {{ todo.title }}
                                </h3>
                                <p class="text-slate-400 line-through truncate ... font-light">
                                {{ todo.description }}
                                </p>
                            </div>
                        </el-col>
                        <el-col :span="1" class="mr">
                            <button @click="deleteTodo(todo.ID)" class="w-6 h-6 mt-4 border-2 rounded-full border-rose-900 opacity-0 group-hover:opacity-100 transition-opacity duration-200">
                                <DeleteIcon />
                            </button>
                        </el-col>
                    </el-row>
                </el-col>
            </el-row>
        </div>

    </el-scrollbar>
</template>

<script setup lang="ts">
import { onBeforeMount, computed } from 'vue'
import { useTodoStore } from '@/stores/todoStore'
import { useUserStore } from '@/stores/userStore'
import { useRouter } from 'vue-router'
import axios from 'axios'
import CheckIcon from '@/assets/icon/CheckIcon.vue'
import DeleteIcon from '@/assets/icon/DeleteIcon.vue'
import WithdrawIcon from '@/assets/icon/WithdrawIcon.vue'
import ExpandIcon from '@/assets/icon/ExpandIcon.vue'
// import { RouterLink } from 'vue-router'

const todoStore = useTodoStore()
const userStore = useUserStore()

const router = useRouter()


// 计算出已完成和未完成的待办事项
const uncompletedTodos = computed(() => todoStore.allTodos.filter((todo) => !todo.isCompleted))
const completedTodos = computed(() => todoStore.allTodos.filter((todo) => todo.isCompleted))

// Mark a todo as completed
const markAsCompleted = async (todoid: number) => {
    try {
        await axios.put(`http://localhost:7879/api/markedtodo/${todoid}`,null, {
            headers: {
                Authorization: `Bearer ${ userStore.userInfo.Token }`
            }
        })
    } catch (error) {
        console.error(error)
    }
    todoStore.markAsCompleted(todoid)
}

// Withdraw a mark
const withDrawMark = async (todoid: number) => {
    try {
        const response = await axios.put(`http://localhost:7879/api/unmarkedtodo/${todoid}`,null, {
            headers: {
                Authorization: `Bearer ${ userStore.userInfo.Token }`
            }
        })
        console.log(response.data)
    } catch (error) {
        console.error(error)
    }
    
    todoStore.unMarkAsCompleted(todoid)
}

const deleteTodo = async (todoid: number) => {
    try {
        const response = await axios.delete(`http://localhost:7879/api/deletetodo/${todoid}`, {
            headers: {
                Authorization: `Bearer ${ userStore.userInfo.Token }`
            }
        })
        console.log(response.data)
    } catch (error) {
        console.error(error)
    }

    todoStore.fetchAllTodos()
}

const TodoDetail = (todoid: number) => {
    router.push({ name: 'TodoView', params: { todoid } })
}

// 钩子：组件挂载后获取所有待办事项
onBeforeMount(() => {
    todoStore.fetchAllTodos()
})
</script>

<style scoped>
</style>
