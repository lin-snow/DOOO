<template>
    <div class="right-side py-3 px-3 ">
        <div>
            <div class="grid place-items-center mb-4">
                <button @click="UserAct()" class="rounded-full w-24 h-9 bg-sky-200 hover:shadow-md">
                    <span class="text-blue-900"> 
                        <el-icon>
                            <User />
                        </el-icon>
                            User
                    </span>
                </button>
            </div>
        </div>

        <div>
            <el-image :src="src" class="w-28 border rounded-lg">
                <template #placeholder>
                    <div class="image-slot preview-src-list">Loading<span class="dot">...</span></div>
                </template>
            </el-image>
        </div>
    </div>
</template>

<script setup lang="ts">
import { User } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'

const src ='/vangoghmuseum-s0164V1962-800.jpg'
const router = useRouter()
function UserAct(){
    if (localStorage.getItem('authToken')) {
        router.push('/user')
    } else {
        router.push('/auth')
    }
}

</script>

<style scoped>
.demo-image__placeholder .block {
    padding: 30px 0;
    text-align: center;
    display: inline-block;
    width: 49%;
    box-sizing: border-box;
    vertical-align: top;
}

.demo-image__placeholder .demonstration {
    display: block;
    color: var(--el-text-color-secondary);
    font-size: 14px;
    margin-bottom: 20px;
}

.demo-image__placeholder .el-image {
    padding: 0 5px;
    max-width: 300px;
    max-height: 200px;
}

.demo-image__placeholder.image-slot {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
    background: var(--el-fill-color-light);
    color: var(--el-text-color-secondary);
    font-size: 14px;
}

.demo-image__placeholder .dot {
    animation: dot 2s infinite steps(3, start);
    overflow: hidden;
}
</style>
