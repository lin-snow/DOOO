<template>
    <h1 class="text-center font-serif text-lg subpixel-antialiased text-amber-800 font-bold underline decoration-wavy underline-offset-8 my-2"> User Center </h1>
    <div class="p-4">
        <h1 class="my-2"> 
        <span class="font-serif text-md subpixel-antialiased text-sky-300 font-bold underline underline-offset-8 block h-10">
            UserName:
        </span>
        <p class="p-1 border rounded-md border-amber-700 shadow-lg text-amber-900">
            {{ currentUser.Username }}
        </p> 
        </h1>
        <h1 class="my-2"> 
            <span class="font-serif text-md subpixel-antialiased text-sky-300 font-bold underline underline-offset-8 block h-10">
                NickName:
            </span>
            <p class="p-1 border rounded-md border-amber-700 shadow-lg text-amber-900">
                {{ currentUser.Nickname }}
            </p> 
        </h1>
        <h1 class="my-2"> 
            <span class="font-serif text-md subpixel-antialiased text-sky-300 font-bold underline underline-offset-8 block h-10">
                Email:
            </span>
            <p class="p-1 border rounded-md border-amber-700 shadow-lg text-amber-900">
                {{ currentUser.Email }}
            </p> 
        </h1>
    
    
        <!-- Logout -->
        <div class="grid justify-items-center my-2">
            <button @click="logout" class="border rounded-xl border-amber-700 shadow-lg p-2 w-20 text-black font-medium hover:bg-sky-100"> Logout </button>
        </div>
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