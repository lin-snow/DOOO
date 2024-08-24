<script setup lang="ts">
import TodoDetail from '@/components/todo/TodoDetail.vue';
</script>

<template>
        <el-container class="max-w-screen-xl h-[65vh] w-96">
            <div class="border rounded-xl shadow-2xl ring ring-orange-100 ring-offset-2 overflow-hidden">
                <TodoDetail />
            </div>
        </el-container>    
</template>

