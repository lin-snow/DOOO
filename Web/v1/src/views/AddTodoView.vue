<template>
    <div>
        <el-container class="max-w-screen-xl h-[65vh] border rounded-xl shadow-2xl ring ring-orange-100 ring-offset-2 overflow-hidden">
            <hr>
                <AddTodo class="mx-16" />
        </el-container>
    </div>
</template>

<script setup lang="ts">
import AddTodo from '@/components/todo/AddTodo.vue';
</script>

<style scoped></style>