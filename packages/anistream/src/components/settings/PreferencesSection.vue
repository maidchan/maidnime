<script setup lang="ts">
import { ref } from 'vue';

const autoPlay = ref(true);
const darkMode = ref(true);
const defaultQuality = ref('1080p');
const subtitleSize = ref('medium');
</script>

<template>
  <section aria-label="App Preferences">
    <h3 class="text-white text-lg font-bold leading-tight tracking-[-0.015em] pb-4 px-1">Preferences</h3>
    
    <div class="flex flex-col gap-px bg-white/5 rounded-xl overflow-hidden shadow-sm">
      <!-- Default Quality -->
      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 p-5 bg-surface-dark">
        <div class="flex flex-col gap-1">
          <label class="text-base font-medium text-white" for="quality-select">Default Quality</label>
          <span class="text-sm text-[#ad92c9]">Set your preferred streaming resolution</span>
        </div>
        <div class="relative min-w-[160px]">
          <select 
            v-model="defaultQuality"
            class="custom-select w-full appearance-none rounded-lg bg-[#1a1122] border border-white/10 text-white py-2.5 pl-4 pr-10 focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all text-sm font-medium" 
            id="quality-select"
          >
            <option value="4k">4K Ultra HD</option>
            <option value="1080p">1080p Full HD</option>
            <option value="720p">720p HD</option>
            <option value="480p">480p SD</option>
          </select>
        </div>
      </div>

      <!-- Auto-play Next -->
      <div 
        @click="autoPlay = !autoPlay"
        class="flex items-center justify-between p-5 bg-surface-dark hover:bg-white/[0.02] transition-colors cursor-pointer"
      >
        <div class="flex items-center gap-3">
          <div class="p-2 rounded-lg bg-white/5 text-[#ad92c9]">
            <span class="material-symbols-outlined text-[20px]">autoplay</span>
          </div>
          <div class="flex flex-col">
            <span class="text-base font-medium text-white">Auto-play Next</span>
            <span class="text-sm text-[#ad92c9]">Automatically start the next episode</span>
          </div>
        </div>
        
        <label class="relative inline-flex items-center cursor-pointer pointer-events-none">
          <input type="checkbox" class="sr-only peer" :checked="autoPlay">
          <div class="w-11 h-6 bg-[#1a1122] peer-focus:outline-none peer-focus:ring-2 peer-focus:ring-primary/50 rounded-full peer peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-primary border border-white/10"></div>
        </label>
      </div>

      <!-- Dark Mode -->
      <div class="flex items-center justify-between p-5 bg-surface-dark hover:bg-white/[0.02] transition-colors">
        <div class="flex items-center gap-3">
          <div class="p-2 rounded-lg bg-white/5 text-[#ad92c9]">
            <span class="material-symbols-outlined text-[20px]">dark_mode</span>
          </div>
          <div class="flex flex-col">
            <span class="text-base font-medium text-white">Dark Mode</span>
            <span class="text-sm text-[#ad92c9]">Use the dark theme for the interface</span>
          </div>
        </div>
        
        <label class="relative inline-flex items-center cursor-not-allowed opacity-80">
          <input type="checkbox" class="sr-only peer" disabled :checked="darkMode">
          <div class="w-11 h-6 bg-[#1a1122] rounded-full peer peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:rounded-full after:h-5 after:w-5 peer-checked:bg-primary border border-white/10"></div>
        </label>
      </div>

      <!-- Subtitle Size -->
      <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4 p-5 bg-surface-dark">
        <div class="flex flex-col gap-1">
          <label class="text-base font-medium text-white" for="subtitle-select">Subtitle Size</label>
          <span class="text-sm text-[#ad92c9]">Adjust text size for subtitles</span>
        </div>
        <div class="relative min-w-[160px]">
          <select 
            v-model="subtitleSize"
            class="custom-select w-full appearance-none rounded-lg bg-[#1a1122] border border-white/10 text-white py-2.5 pl-4 pr-10 focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all text-sm font-medium" 
            id="subtitle-select"
          >
            <option value="small">Small</option>
            <option value="medium">Medium</option>
            <option value="large">Large</option>
            <option value="huge">Huge</option>
          </select>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.custom-select {
    background-image: url(https://lh3.googleusercontent.com/aida-public/AB6AXuD2T7PLCpeXWQfb6UmbA1jLJTcz8E11vdTo3-axHL7pFitFlqDC9iWamkILBjdscTGnXNk92iK8nW5rTDnW7QMfxRwvbqUu9eyv9fYnwt2jNOiyUi-bY8yrS6bOptg2lGlnF3FjQcSpkbWT1i0Hsnt7prZBPR3Dz1Aa-8nDgWIHlYYPON0kE6zK8sBZspccT9KalKZdJQGeCstf-F7vzvd5vw_Sxhz_y2VBvwg8Zl7bCLFpArA4f6Aq70Q_5Lsmhh9uW216vdqZ);
    background-position: right 1rem center;
    background-repeat: no-repeat;
    background-size: 1.5em 1.5em
}
</style>
