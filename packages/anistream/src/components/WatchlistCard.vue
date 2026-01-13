<script setup lang="ts">
defineProps<{
    title: string;
    image: string;
    genre: string;
    badge?: string;
    badgeColor?: string;
}>();
</script>

<template>
  <div class="flex flex-col gap-3 group cursor-pointer">
    <div class="aspect-2/3 w-full rounded-xl bg-surface-dark relative overflow-hidden ring-1 ring-white/5 group-hover:ring-primary/50 transition-all shadow-lg">
      <div 
        class="w-full h-full bg-cover bg-center transition-transform duration-500 group-hover:scale-105" 
        :style="{ backgroundImage: `url('${image}')` }"
      ></div>
      
      <div class="absolute inset-0 bg-linear-to-t from-black/80 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity flex flex-col justify-end p-3">
        <router-link to="/watch" class="w-full py-2 bg-white text-black font-bold text-xs rounded shadow-lg flex items-center justify-center gap-1 hover:bg-gray-200">
          <span class="material-symbols-outlined text-sm">play_arrow</span> Watch
        </router-link>
      </div>
      
      <div v-if="badge" class="absolute top-2 right-2 bg-black/60 backdrop-blur-md px-1.5 py-0.5 rounded text-[10px] font-bold text-white">
        {{ badge }}
      </div>
      <div v-if="badgeColor === 'primary'" class="absolute top-2 left-2 bg-primary text-white px-1.5 py-0.5 rounded text-[10px] font-bold">NEW</div>
    </div>
    
    <div>
      <h4 class="text-white font-bold truncate group-hover:text-primary transition-colors">{{ title }}</h4>
      <p class="text-[#ad92c9] text-xs">{{ genre }}</p>
    </div>
  </div>
</template>
