<script setup lang="ts">
import EditTodo from '@/components/todo/EditTodo.vue'
</script>

<template>
    <el-container class="max-w-screen-xl h-[65vh] w-96 flex items-center justify-center">
        <div class="w-2/4 mx-16 border rounded-xl shadow-2xl ring ring-orange-100 ring-offset-2 overflow-hidden bg-amber-50">
            <EditTodo />
        </div>
    </el-container>   
</template>

