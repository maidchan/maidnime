<script setup lang="ts">
interface Props {
    id: string | number;
    title: string;
    image: string;
    rating?: string | number;
    year?: string | number;
    genre?: string;
}

defineProps<Props>();
</script>

<template>
  <router-link :to="`/anime/${id}`" class="group flex flex-col gap-3 cursor-pointer">
    <div class="relative w-full aspect-2/3 rounded-xl overflow-hidden bg-surface-dark shadow-lg">
      <img 
        class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110" 
        :src="image" 
        :alt="title"
        loading="lazy"
      />
      <div v-if="rating" class="absolute top-2 right-2 bg-black/60 backdrop-blur-md px-1.5 py-0.5 rounded flex items-center gap-1 border border-white/10">
        <span class="material-symbols-outlined text-yellow-400 text-[12px] fill-current">star</span>
        <span class="text-xs font-bold text-white">{{ rating }}</span>
      </div>
      <div class="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 transition-opacity duration-300 flex items-center justify-center">
        <div class="size-12 rounded-full bg-primary text-white flex items-center justify-center shadow-[0_0_20px_rgba(127,19,236,0.6)] transform scale-90 group-hover:scale-100 transition-transform">
          <span class="material-symbols-outlined text-[32px] ml-1">play_arrow</span>
        </div>
      </div>
    </div>
    <div class="flex flex-col gap-0.5">
      <h4 class="text-white font-bold text-base truncate group-hover:text-primary transition-colors">{{ title }}</h4>
      <div class="flex items-center gap-2 text-xs text-gray-400">
        <span v-if="year">{{ year }}</span>
        <span v-if="year && genre" class="size-1 rounded-full bg-gray-600"></span>
        <span v-if="genre">{{ genre }}</span>
      </div>
    </div>
  </router-link>
</template>
