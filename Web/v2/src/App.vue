<template>
  <div class="flex items-center justify-center h-screen bg-gray-200">
    <div class="bg-white shadow-lg rounded-lg overflow-hidden" style="width: 800px; height: 600px">
      <div class="flex h-full">
        <SideBar />
        <div class="flex flex-col w-max w-auto">
          <CategoryComp />
          <CenterContent />
        </div>
        <RightSide />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import SideBar from './components/SideBar.vue'
import CategoryComp from './components/CategoryComp.vue'
import CenterContent from './components/CenterContent.vue'
import RightSide from './components/RightSide.vue'
</script>

<style scoped>
/* 全局样式可在这里添加 */
</style>

<!-- 
<template>
  <div class="flex h-screen bg-gray-200">
    
    <aside class="w-1/5 bg-gray-300 p-4 flex flex-col items-center space-y-6">
      <div class="w-20 h-20 bg-red-300 rounded-full flex items-center justify-center">
        Logo
      </div>
      <p class="text-lg text-red-700">Nickname</p>
      <button class="bg-red-400 py-2 px-6 rounded-full text-white">ADD</button>
      <nav class="w-full mt-6">
        <ul class="space-y-4 text-center">
          <li class="bg-black text-white py-2 rounded">Home</li>
          <li class="text-pink-400 py-2">Default</li>
          <li class="text-pink-400 py-2">Cate1</li>
          <li class="text-pink-400 py-2">Cate2</li>
          <li class="text-pink-400 py-2">Cate3</li>
          <li class="text-pink-400 py-2">Cate4</li>
        </ul>
      </nav>
    </aside>
    
    
    <main class="flex-1 flex flex-col p-6 bg-gray-500">
      <h1 class="text-lg text-gray-100 mb-4">Category</h1>
      <div class="grid grid-cols-3 gap-4">
        <div class="bg-yellow-300 h-24 flex items-center justify-center">todo1</div>
        <div class="bg-yellow-300 h-24 flex items-center justify-center">todo2</div>
        <div class="bg-yellow-300 h-24 flex items-center justify-center">todo3</div>
        <div class="bg-yellow-300 h-24 flex items-center justify-center">todo4</div>
      </div>
    </main>

    
    <aside class="w-1/5 bg-gray-400 p-4 flex flex-col items-center space-y-6">
      <button class="bg-gray-600 text-white py-2 px-6 rounded-full">Login</button>
      <div class="bg-orange-400 w-28 h-28 flex items-center justify-center text-red-800">Picture</div>
    </aside>
  </div>
</template>

<script setup lang="ts">
// 可以在此添加任何逻辑或状态
</script>

<style scoped>
/* 可以在这里编写任何特定的样式 */
</style> -->
