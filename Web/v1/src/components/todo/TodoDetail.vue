<template>
    <div>
        <h1>
            TODO Detail
        </h1>
        <hr>
        <span>
            ID: {{ currentTodo.ID }}
        </span>
        <span>
            {{ currentTodo.title }}
        </span>
    </div>
</template>

<script setup lang="ts">
import { onBeforeMount, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useTodoStore } from '@/stores/todoStore';

const todoStore = useTodoStore();
const currentTodo = ref<any>({});

const route = useRoute();
// const router = useRouter();
const curtodoid = Number(route.params.todoid);

console.log('@@@@@@@@@@curtodoid', curtodoid);

onBeforeMount(() => {
    console.log('TodoDetail onBeforeMount');
    console.log('before@@@@@@@@@@curtodoid', curtodoid);
})

onMounted(() => {
    console.log('TodoDetail onMounted');

    const currentTodoData = todoStore.allTodos.filter(item => item.ID === curtodoid)[0];
    currentTodo.value = currentTodoData;
    console.log('on@@@@@@@@@@currentTodo', currentTodo.value);
})


</script>