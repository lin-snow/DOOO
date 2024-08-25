<template>
    <div>
        <h1 class="text-center font-mono text-lg subpixel-antialiased text-amber-800 font-bold underline decoration-wavy underline-offset-8"> Login </h1>
        <div class="p-4 flex items-center justify-center">
            <form @submit.prevent="login" class="flex flex-col ">
                <!-- Username -->
                <div class="my-2">
                    <!-- <span class="font-serif text-md subpixel-antialiased text-black font-bold mx-4"> Username: </span> -->
                    <input type="text" class="border rounded-lg border-neutral-400 w-44 p-2 shadow-inner my-2 overflow-y-hidden focus:border-amber-800 outline-none focus:border-2" placeholder="Username" v-model="loginInfo.username" autocomplete="username"/>
                </div>
                <!-- Password -->
                <div class="my-2">
                    <!-- <span class="font-serif text-md subpixel-antialiased text-black font-bold mx-5"> Password: </span> -->
                    <input type="password" class="border rounded-lg border-neutral-400 w-44 p-2 shadow-inner my-2 overflow-y-hidden focus:border-amber-800 outline-none focus:border-2" placeholder="Password" v-model="loginInfo.password" autocomplete="current-password"/>
                </div>
                <!-- Submit -->
                <div class="flex items-center justify-center my-2">
                    <button type="submit" class="border rounded-xl border-amber-700 shadow-lg p-2 w-20 text-black font-medium hover:bg-sky-100">Login</button>
                </div>           
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { authClient } from '@/utils/axios/axios'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/userStore';

const router = useRouter();

const loginInfo = reactive({
    username: '',
    password: ''
})

const login = async () => {
    try {
        const response = await authClient.post('/login', loginInfo);

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