<template>
    <div class="p-2">
        <h1 class="text-center font-mono text-lg subpixel-antialiased text-amber-800 font-bold underline decoration-wavy underline-offset-8 my-4"> Register </h1>
        <form @submit.prevent="register" class="flex flex-col">
            <input type="text" class="border rounded-lg border-neutral-400 w-44 p-2 shadow-inner my-2 overflow-y-hidden focus:border-amber-800 outline-none focus:border-2" placeholder="Username" v-model="UserInfo.username" />
            <input type="email" class="border rounded-lg border-neutral-400 w-44 p-2 shadow-inner my-2 overflow-y-hidden focus:border-amber-800 outline-none focus:border-2" placeholder="Email" v-model="UserInfo.email" />
            <input type="password" class="border rounded-lg border-neutral-400 w-44 p-2 shadow-inner my-2 overflow-y-hidden focus:border-amber-800 outline-none focus:border-2" placeholder="Password" v-model="UserInfo.password" autocomplete="new-password"/>
            <input type="text" class="border rounded-lg border-neutral-400 w-44 p-2 shadow-inner my-2 overflow-y-hidden focus:border-amber-800 outline-none focus:border-2" placeholder="Nickname" v-model="UserInfo.nickname" />
            <div class="flex items-center justify-center my-2">
                <button type="submit" class="border rounded-xl border-amber-700 shadow-lg p-2 w-20 text-black font-medium hover:bg-sky-100">SignUp</button>
            </div>  
        </form>
    </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'


const router = useRouter();

const UserInfo = reactive({
    username: '',
    password: '',
    email: '',
    nickname: ''
})

const register = async () => {
    try {
        const response = await axios.post('http://127.0.0.1:7879/signup', UserInfo);
        // 假设后端返回token
        localStorage.setItem('authToken', response.data.token);

        // 重定向到登录页面或主页
        router.push('/auth');
    } catch (error) {
        console.error(error);
    }
};

defineExpose({
    register
});
</script>

<style scoped></style>
