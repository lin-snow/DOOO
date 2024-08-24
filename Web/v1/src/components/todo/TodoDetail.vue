<template>
    <div>
        <h1> {{ currentTodo.title }} </h1>
        <hr>
        <p> {{ currentTodo.description }} </p>
        <hr>
        <span> {{ currentTodo.isCompleted }} </span>
    </div>

    <div>
        <!-- Edit Todo -->
        <el-button @click="EditTodo(currentTodo.ID)" type="primary" :icon="Edit" class="w-24 border rounded-md "> Edit Todo </el-button>
    </div>

</template>

<script setup lang="ts">
import { useTodoStore } from '@/stores/todoStore';
import { useRoute, useRouter } from 'vue-router';
import { Edit } from '@element-plus/icons-vue';

const todoStore = useTodoStore();
const route = useRoute();
const router = useRouter();
const todoid = Number(route.params.todoid);
const currentTodo = todoStore.allTodos.filter(item => item.ID === todoid)[0]

function EditTodo (todoid: number) {
    router.push(`/todo/edit/${todoid}`);
}

</script>