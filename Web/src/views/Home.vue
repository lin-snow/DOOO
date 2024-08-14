<template>
    <div>
        <div class="flex flex-nowrap w-auto h-auto gap-11">
            <SideBar />
            <div class="bg-red-700 h-auto">
                <Header />
                <Content />
            </div>
            <RightSide/>
        </div>
        
    </div>
</template>

<script setup lang="ts">
import SideBar from '@/components/SideBar.vue';
import Header from '@/components/Header.vue';
import Content from '@/components/Content.vue';
import RightSide from '@/components/RightSide.vue';
</script>

<style scoped></style>
