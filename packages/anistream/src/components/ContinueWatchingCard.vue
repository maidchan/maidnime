<script setup lang="ts">
defineProps<{
    title: string;
    episode: string;
    progress: number;
    timeLeft: string;
    image: string;
    season: string;
}>();
</script>

<template>
  <div class="group flex flex-col sm:flex-row items-stretch gap-4 rounded-xl bg-surface-dark p-3 hover:bg-[#322142] transition-colors border border-transparent hover:border-white/5 cursor-pointer">
    <div class="relative w-full sm:w-48 aspect-video sm:aspect-auto rounded-lg overflow-hidden shrink-0">
      <div 
        class="w-full h-full bg-cover bg-center transition-transform duration-700 group-hover:scale-110" 
        :style="{ backgroundImage: `url('${image}')` }"
      ></div>
      <div class="absolute inset-0 bg-black/30 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
        <span class="material-symbols-outlined text-white text-4xl drop-shadow-lg">play_circle</span>
      </div>
    </div>
    
    <div class="flex flex-col justify-center flex-1 py-1 pr-2">
      <div class="flex justify-between items-start">
        <h4 class="text-white text-base font-bold line-clamp-1">{{ title }}</h4>
        <span class="text-xs font-bold text-primary bg-primary/10 px-2 py-0.5 rounded">{{ season }}</span>
      </div>
      <p class="text-[#ad92c9] text-sm mt-1">{{ episode }}</p>
      
      <div class="mt-auto pt-3">
        <div class="flex justify-between text-xs text-[#ad92c9] mb-1.5">
          <span>{{ timeLeft }} left</span>
          <span>{{ progress }}%</span>
        </div>
        <div class="w-full bg-black/40 rounded-full h-1.5 overflow-hidden">
          <div 
            class="bg-primary h-full rounded-full shadow-[0_0_10px_rgba(127,19,236,0.5)]"
            :style="{ width: `${progress}%` }"
          ></div>
        </div>
      </div>
    </div>
  </div>
</template>
