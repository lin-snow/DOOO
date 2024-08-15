<template>
    <div>
        <div class="flex flex-nowrap w-auto h-auto gap-11">
            <SideBar />
            <div class="bg-red-700 h-auto">
                <div>
                    <CategoryComp />
                </div>
                <Content />
            </div>
            <RightSide/>
        </div>
        
    </div>
</template>

<script setup lang="ts">
import SideBar from '@/components/SideBar.vue';
import CategoryComp from '@/components/CategoryComp.vue';
import Content from '@/components/CenterContent.vue';
import RightSide from '@/components/RightSide.vue';

</script>

<style scoped></style>
