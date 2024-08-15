<template>
    <div class="left-side py-3 px-3">
        <!-- Logo -->
        <div class="grid place-items-center my-3">
            <el-col :span="12">
                <div class="demo-basic--circle">
                    <div class="block">
                        <el-avatar :size="100" :src="circleUrl" class="shadow-lg" />
                    </div>
                </div>
            </el-col>
        </div>

        <!-- NickName -->
        <div class="mb-4">
            <Span class="font-mono text-lg subpixel-antialiased italic font-medium text-slate-900">
                🫡 Nickname
            </Span>
        </div>

        <!-- Add and Home -->
        <div class="grid place-items-center mb-4">
            <el-button round size="large" color="#f78989" class="shadow-xl bg-red-400 w-20">
                Add
            </el-button>
        </div>
        <!-- All Categories -->
        <div>
            <el-col :span="27">
                <el-menu default-active="2" class="border rounded-lg overflow-hidden shadow-xl" @open="handleOpen" @close="handleClose">
                    <el-sub-menu index="1">
                        <template #title>
                            <el-icon>
                                <Files />
                            </el-icon>
                            <span> Categories </span>
                        </template>
                        <div>
                            <el-menu-item index="1-1">Study</el-menu-item>
                            <el-menu-item index="1-2">Working</el-menu-item>
                        </div>
                        
                    </el-sub-menu>
                </el-menu>
            </el-col>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive, toRefs } from 'vue'
import { Files } from '@element-plus/icons-vue'

const state = reactive({
    circleUrl: 'https://linsnow.cn/img/avatar.png'
})

const { circleUrl } = toRefs(state)

const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
</script>

<style>
/* Left Side Logo */

</style>
