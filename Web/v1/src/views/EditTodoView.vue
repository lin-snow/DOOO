<template>
    <div>
        <el-container class="max-w-screen-xl h-[65vh] border rounded-xl shadow-2xl ring ring-orange-100 ring-offset-2 overflow-hidden bg-amber-50">   
            <!-- EditTodo View -->
            <EditTodo />
        </el-container>
    </div>
</template>

<script setup lang="ts">
import EditTodo from '@/components/todo/EditTodo.vue'
</script>