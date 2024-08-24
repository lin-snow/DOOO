<template>
    <div>
        <h1 class="text-center font-serif text-lg subpixel-antialiased text-amber-800 font-bold underline decoration-wavy underline-offset-8 my-2"> Todo Details </h1>
        
        <div class="m-4 text-pretty break-all">
            <h1 class="my-2"> 
                <span class="font-serif text-md subpixel-antialiased text-black font-bold underline underline-offset-8 block h-10">
                    Title:
                </span>
                <p class="p-1 border rounded-md border-amber-700 shadow-lg text-amber-900">
                    {{ currentTodo.title }}
                </p> 
            </h1>

            <h1 class="my-2"> 
                <span class="font-serif text-md subpixel-antialiased text-black font-bold underline underline-offset-8 block h-10">
                    Description:
                </span>
                <p class="p-1 border rounded-md border-amber-700 shadow-lg text-amber-900">
                    {{ currentTodo.description }}
                </p> 
            </h1>
        </div>
    </div>
    
    
    <!-- Edit Todo -->
    <div class="grid justify-items-center my-2">
        <button @click="EditTodo(currentTodo.ID)" class="border rounded-xl border-amber-700 shadow-lg p-2 w-20 text-black font-medium hover:bg-sky-100"> Edit </button>
    </div>
</template>

<script setup lang="ts">
import { useTodoStore } from '@/stores/todoStore';
import { useRoute, useRouter } from 'vue-router';
// import { Edit } from '@element-plus/icons-vue';

const todoStore = useTodoStore();
const route = useRoute();
const router = useRouter();
const todoid = Number(route.params.todoid);
const currentTodo = todoStore.allTodos.filter(item => item.ID === todoid)[0]

function EditTodo (todoid: number) {
    router.push(`/todo/edit/${todoid}`);
}

</script>