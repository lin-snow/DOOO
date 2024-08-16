<template>
    <div>
        <h1>Login</h1>
        <!-- Login Form -->
        <form @submit.prevent="login" class="flex flex-col">
            <input type="text" class="border border-gray-300 p-2" placeholder="Username" v-model="loginInfo.username" autocomplete="username"/>
            <input type="password" class="border border-gray-300 p-2" placeholder="Password"
                v-model="loginInfo.password" autocomplete="current-password"/>
            <button type="submit" class="bg-blue-500 text-white p-2">Login</button>
        </form>
    </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/userStore';

const router = useRouter();

const loginInfo = reactive({
    username: '',
    password: ''
})

const login = async () => {
    try {
        const response = await axios.post('http://127.0.0.1:7879/login', loginInfo);

        // Debug: Print response to check if token is received
        console.log('Response Status:', response.status);
        console.log('Response Data:', response.data);

        const token = response.data.data.token;
        console.log('Token before storing:', token);

        const UserInfo = {
            UserID: response.data.data.userid,
            Username: response.data.data.username,
            Email: response.data.data.email,
            Nickname: response.data.data.nickname,
            Token: response.data.data.token
        }

        if (response.status === 200 && token) {
            localStorage.setItem('authToken', token);
            console.log('Token stored:', localStorage.getItem('authToken'));

            // Using Pinia Store
            const userStore = useUserStore();
            userStore.setUserInfo(UserInfo);

            router.push('/');
        } else {
            alert(response.data.Message);
        }
    } catch (error) {
        console.error(error);
    }
};


</script>

<style scoped></style>