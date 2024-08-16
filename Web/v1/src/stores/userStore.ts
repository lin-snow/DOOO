// Store Current User Information
import { defineStore } from 'pinia'
import { ref } from 'vue'
import axios from 'axios'

export const useUserStore = defineStore('user', () => {
    // State
    const userInfo = ref({
        UserID: 0,
        Username: '',
        Email: '',
        Nickname: '',
        Token: ''
    })

    // Actions
    function setUserInfo(user: {UserID: number; Username: string; Email: string; Nickname: string; Token: string }) {
        userInfo.value = user
    }

    const fetchUserInfo = async () => {
        const authToken = localStorage.getItem('authToken')
        if (!authToken) {
            console.error('No auth token found')
            return
        }

        try {
            const response = await axios.get('http://127.0.0.1:7879/api/getuserinfo', {
                headers: {
                    'Authorization': `Bearer ${ authToken }` // 添加 token 到请求头中
                }  
            })
            // Set UserInfo
            userInfo.value.UserID = response.data.data.userid
            userInfo.value.Username = response.data.data.username
            userInfo.value.Email = response.data.data.email
            userInfo.value.Nickname = response.data.data.nickname
            userInfo.value.Token = authToken

            // console.log('User Info:', userInfo.value)
        } catch (error) {
            console.error(error)
        }
    }

    fetchUserInfo()

    return { userInfo, setUserInfo, fetchUserInfo }
})
