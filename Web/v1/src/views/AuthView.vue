<template>
    <div>
        <el-container class="max-w-screen-xl h-[65vh] border rounded-xl shadow-2xl ring ring-orange-100 ring-offset-2 overflow-hidden flex items-center justify-center">
                <aside class="mx-16">
                    <h2>Register</h2>
                    <UserRegister  />
                </aside>
                <aside class="mx-16">
                    <h2>Login</h2>
                    <UserLogin />
                </aside>
        </el-container>
    </div>
</template>

<script setup lang="ts">
import UserRegister from '@/components/user/UserRegister.vue';
import UserLogin from '@/components/user/UserLogin.vue';

</script>
<style scoped></style>