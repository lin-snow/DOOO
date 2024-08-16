import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/HomeView.vue'
import Auth from '@/views/AuthView.vue'
import About from '@/views/AboutView.vue'
import AddTodo from '@/components/todo/AddTodo.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        { path: '/', name: 'Home', component: Home},
        { path: '/auth', name: 'Auth', component: Auth },
        { path: '/about', name: 'About', component: About},
        { path: '/addtodo', name: 'AddTodo', component: AddTodo}
    ]
})

export default router
