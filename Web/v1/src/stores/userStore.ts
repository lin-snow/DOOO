// Store Current User Information
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { apiClient } from '@/utils/axios/axios'

export const useUserStore = defineStore('user', () => {
    // State
    const userInfo = ref({
        UserID: 0,
        Username: '',
        Email: '',
        Nickname: '',
        Token: ''
    })

    const loginStatus = ref(true)

    // Actions
    function setUserInfo(user: {UserID: number; Username: string; Email: string; Nickname: string; Token: string }) {
        userInfo.value = user
    }

    const fetchUserInfo = async () => {
        try {
            const response = await apiClient.get('/getuserinfo')
            // Set UserInfo
            userInfo.value.UserID = response.data.data.userid
            userInfo.value.Username = response.data.data.username
            userInfo.value.Email = response.data.data.email
            userInfo.value.Nickname = response.data.data.nickname
            userInfo.value.Token = localStorage.getItem('authToken') || ''

            loginStatus.value = false
            // console.log('User Info:', userInfo.value)
        } catch (error) {
            console.error(error)
        }
    }

    fetchUserInfo()

    const logout = () => {
        userInfo.value = {
            UserID: 0,
            Username: '',
            Email: '',
            Nickname: '',
            Token: ''
        }
        loginStatus.value = true
    }

    return { userInfo, loginStatus, setUserInfo, fetchUserInfo, logout }
},
{
    persist: {
        key: 'userdata',
        storage: localStorage,
        paths: ['userInfo'],   
    }
})
