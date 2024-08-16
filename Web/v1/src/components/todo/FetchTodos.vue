<template>
    <div v-if="todoStore.loading" class="text-center">Loading...</div>
    <el-scrollbar v-else class="overflow-hidden h-full" height="500px">
        <div v-for="todo in uncompletedTodos" :key="todo.ID"
            class="my-2 rounded-lg p-4 border-2 shadow-sm border-amber-900">
            <h3 class="text-amber-950">
                {{ todo.title }}
            </h3>
            <p class="text-slate-400 truncate ...">
                {{ todo.description }}
            </p>
        </div>
        <div v-for="todo in completedTodos" :key="todo.ID"
            class="my-2 rounded-lg p-4 border-2 shadow-sm border-amber-900 border-dashed ">
            <h3 class=" line-through text-gray-500 border-dashed ">
                {{ todo.title }}
            </h3>
            <p class="'line-through text-gray-500 ">
                {{ todo.description }}
            </p>
        </div>
    </el-scrollbar>
</template>

<script setup lang="ts">
import { onBeforeMount, computed } from 'vue'
// import axios from 'axios'
import { useTodoStore } from '@/stores/todoStore'

// const todos = ref<Todo[]>([])
const todoStore = useTodoStore()

// 计算出已完成和未完成的待办事项
const uncompletedTodos = computed(() => todoStore.allTodos.filter((todo) => !todo.isCompleted))
const completedTodos = computed(() => todoStore.allTodos.filter((todo) => todo.isCompleted))

// 钩子：组件挂载后获取所有待办事项
onBeforeMount(() => {
    todoStore.fetchAllTodos()
})
</script>

<style scoped></style>
