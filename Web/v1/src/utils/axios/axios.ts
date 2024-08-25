// packaging the axios instance
import axios from 'axios'

// Create an axios instance

// For Auth Request
const authClient = axios.create(
    {
        baseURL: import.meta.env.VITE_AUTH_URL, // get the auth url from .env file
        timeout: 10000, // set the timeout to 10s
        headers: {
            'Content-Type': 'application/json', // set the content type to json
        }
    }
)

// For API Request
const apiClient = axios.create({
    baseURL: import.meta.env.VITE_API_URL, // get the api url from .env file
    timeout: 10000, // set the timeout to 10s
})

// Add a interceptor to the axios instance
apiClient.interceptors.request.use(
    (config) => {
        // Dynamically set the Content-Type header
        if (config.data) {
            if (config.data instanceof FormData) {
                config.headers['Content-Type'] = 'multipart/form-data'
            } else if (typeof config.data === 'string') {
                config.headers['Content-Type'] = 'application/x-www-form-urlencoded'
            } else {
                config.headers['Content-Type'] = 'application/json'
            }
        }

        // Auto add the token to the header
        const token = localStorage.getItem('authToken')
        if (!token) {
            console.error('No token found!!!')
        }

        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`
        }
        
        return config
    },
    (error) => Promise.reject(error)
)

// Response interceptor
apiClient.interceptors.response.use(
    (Response) => Response,
    (error) => {
        // handle the error
        console.log('Response Error:', error)
        return Promise.reject(error)
    }
)


// Export the axios instance
export { authClient, apiClient}