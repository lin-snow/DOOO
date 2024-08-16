<template>
    <div>
        <!-- Register -->
        <div class="flex flex-col">
            <input type="text" class="border border-gray-300 p-2" placeholder="Username" v-model="UserInfo.username" />
            <input type="email" class="border border-gray-300 p-2" placeholder="Email" v-model="UserInfo.email" />
            <input type="password" class="border border-gray-300 p-2" placeholder="NewPassword" v-model="UserInfo.password" autocomplete="new-password"/>
            <input type="text" class="border border-gray-300 p-2" placeholder="Nickname" v-model="UserInfo.nickname" />
            <button class="bg-blue-500 text-white p-2" @click="register">SignUp</button>
        </div>
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
