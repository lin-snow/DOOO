<template>
    <div>
        <el-container class="w-lvm max-w-screen-xl h-[65vh] border rounded-xl shadow-2xl ring ring-orange-100 ring-offset-2 overflow-hidden">
            <div class="w-2/4 mx-16">
                <AddTodo/>
            </div>
        </el-container>
    </div>
</template>

<script setup lang="ts">
import AddTodo from '@/components/todo/AddTodo.vue';
</script>

<style scoped></style>