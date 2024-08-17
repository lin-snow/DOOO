<script setup lang="ts">


import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useTodoStore } from '@/stores/todoStore';
import EditIcon from '@/assets/icon/EditIcon.vue';

// const todoStore = useTodoStore();
const route = useRoute();
const router = useRouter();
const curtodoid = Number(route.params.todoid);
const todoStore = useTodoStore();
const alltodos = todoStore.allTodos;
const currentTodo = ref(alltodos.filter(item => item.ID === curtodoid)[0]);

const EditTodo = (todoid: number) => {
    router.push({ name: 'EditTodo', params: { todoid } })
}


onMounted(() => {
    console.log('TodoView onMounted');
});

</script>

<template>
    <div>
        <el-container class="max-w-screen-xl h-[65vh] border rounded-xl shadow-2xl ring ring-orange-100 ring-offset-2 overflow-hidden">   
            <!-- TodoDetail View -->
            <div>
        <h1>Todo Detail</h1>
        <div>
            <div>
                <div>
                    <span>Todo ID: {{ currentTodo.ID }}</span>
                </div>
                <div>
                    <span>Todo Title: {{ currentTodo.title }}</span>
                </div>
                <div>
                    <span>Todo Description: {{ currentTodo.description }}</span>
                </div>
                <div>
                    <span>Todo Status: {{ currentTodo.isCompleted }}</span>
                </div>
            </div>
        </div>
        <div>
            <button @click="EditTodo(curtodoid)" class="w-6 h-6 border-2 rounded-full border-rose-900"> 
                <EditIcon />
            </button>
        </div>
    </div>
        </el-container>
    </div>
</template>

