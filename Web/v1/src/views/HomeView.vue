<script setup lang="ts">
import LeftSide from '@/components/LeftSide.vue'
import MainContent from '@/components/MainContent.vue'
import RightSide from '@/components/RightSide.vue'
</script>

<template>
  <el-container class="max-w-screen-xl h-[65vh] border rounded-xl shadow-2xl ring ring-orange-100 ring-offset-2 overflow-hidden">
    <el-aside width="200px" class="bg-orange-50 ">
        <!-- Left Side -->
        <LeftSide />
      </el-aside>
      <el-container>
        <el-main class="bg-amber-50">
          <!-- Main Content -->
          <MainContent />
        </el-main>
      </el-container>
      <el-aside width="200px" class="bg-amber-50  border-dashed border-l-rose-200 border-l-2">
        <!-- Right Side -->
        <RightSide />
      </el-aside>
  </el-container>  
</template>