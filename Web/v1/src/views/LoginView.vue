<script setup lang="ts">
import UserLogin from '@/components/user/UserLogin.vue'
</script>

<template>
    <el-container class="max-w-screen-xl h-[65vh] w-96 flex items-center justify-center">
        <div class="w-2/4 h-6/12 mx-16 border rounded-xl shadow-2xl ring ring-orange-100 ring-offset-2 overflow-hidden bg-amber-50 flex item-center justify-center">
            <UserLogin />
        </div>
    </el-container>       
</template>