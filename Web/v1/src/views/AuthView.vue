<template>
    <el-container class="max-w-screen-xl h-[65vh] w-96 flex items-center justify-center">
        <div class="w-3/12 h-1/5 mx-16 border rounded-xl shadow-2xl ring ring-orange-100 ring-offset-2 overflow-hidden bg-amber-50 flex items-center justify-center">
            <!-- Register -->
            <router-link to="/register">
                <button  class="border rounded-xl border-amber-700 shadow-lg p-2 mx-2 w-20 text-black font-medium hover:bg-sky-200"> Register </button>
            </router-link>
            
            <!-- Login -->
            <router-link to="/login">
                <button class="border rounded-xl border-amber-700 shadow-lg p-2 mx-2 w-20 text-black font-medium hover:bg-sky-200"> Login </button>
            </router-link>           
        </div>
    </el-container> 
</template>

<script setup lang="ts">

</script>
