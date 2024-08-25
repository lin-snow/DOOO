import { createRouter, createWebHistory } from 'vue-router'

// import.meta.env.BASE_URL
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        { path: '/', name: 'Home', component: () => import('@/views/HomeView.vue')},
        { path: '/user', name: 'User', component: () => import('@/views/UserView.vue')},
        { path: '/auth', name: 'Auth', component: () => import('@/views/AuthView.vue')},
        { path: '/login', name: 'Login', component: () => import('@/views/LoginView.vue')},
        { path: '/register', name: 'Register', component: () => import('@/views/RegisterView.vue')},
        { path: '/about', name: 'About', component: () => import('@/views/AboutView.vue')},
        { path: '/addtodo', name: 'AddTodo', component: () => import('@/views/AddTodoView.vue')},
        { path: '/todo/:todoid', name: 'TodoView', component: () => import('@/views/TodoView.vue')},
        { path: '/todo/edit/:todoid', name: 'EditTodo', component: () => import('@/views/EditTodoView.vue')},
    ]
})

export default router
