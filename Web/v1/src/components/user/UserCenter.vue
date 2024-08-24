<template>
    <div>
        <h1> User Center </h1>
        <hr>
        <h2>
            <span class="text-red-400"> Username: </span>
            {{ currentUser.Username }}
        </h2>
        <hr>
        <p>
            <span class="text-red-400"> Nickname: </span>
            {{ currentUser.Nickname }}
        </p>
        <hr>
        <p>
            <span class="text-red-400"> Email: </span>
            {{ currentUser.Email}}
        </p>


        <span>
            <el-button @click="logout" type="primary" class="w-24 border rounded-md "> Logout </el-button>
        </span>
    </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/userStore'
import { useTodoStore } from '@/stores/todoStore'
import { useRouter } from 'vue-router'


const userStore = useUserStore()
const todoStore = useTodoStore()
const currentUser = userStore.userInfo
const router = useRouter()

function logout() {
    localStorage.removeItem('authToken')
    userStore.logout()
    todoStore.clear()

    router.push('/')
}
</script>