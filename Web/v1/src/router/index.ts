import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/HomeView.vue'
import Auth from '@/views/AuthView.vue'
import About from '@/views/AboutView.vue'
import AddTodo from '@/views/AddTodoView.vue'
import EditTodo from '@/views/EditTodoView.vue'
import TodoView from '@/views/TodoView.vue'

// import.meta.env.BASE_URL
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL || '/'),
    routes: [
        { path: '/', name: 'Home', component: Home},
        { path: '/auth', name: 'Auth', component: Auth },
        { path: '/about', name: 'About', component: About},
        { path: '/addtodo', name: 'AddTodo', component: AddTodo},
        { path: '/edittodo/:todoid', name: 'EditTodo', component: EditTodo},
        { path: '/todo/:todoid', name: 'TodoView', component: TodoView }
    ]
})

export default router
