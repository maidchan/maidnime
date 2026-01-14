<script setup lang="ts">
import AnimeCard from './AnimeCard.vue';

import type { Anime } from '../services/api';

defineProps<{
    animeList: Anime[];
}>();
</script>

<template>
  <section class="mt-16 border-t border-white/5 pt-12">
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-2xl font-bold text-white">Anime Terkait</h3>
      <div class="flex gap-2">
        <button class="p-2 rounded-full border border-white/10 hover:bg-white/10 text-white transition-colors">
          <span class="material-symbols-outlined text-sm">arrow_back</span>
        </button>
        <button class="p-2 rounded-full border border-white/10 hover:bg-white/10 text-white transition-colors">
          <span class="material-symbols-outlined text-sm">arrow_forward</span>
        </button>
      </div>
    </div>
    <div class="flex gap-6 overflow-x-auto pb-8 scrollbar-hide scroll-smooth">
        <div v-for="anime in animeList" :key="anime.id" class="w-[160px] md:w-[200px] shrink-0">
            <AnimeCard 
                :id="anime.id"
                :title="anime.title"
                :image="anime.image"
                :rating="anime.rating"
                :year="anime.status"
                :genre="anime.genre && anime.genre.length ? anime.genre[0] : ''"
            />
        </div>
    </div>
  </section>
</template>

<style scoped>
.scrollbar-hide::-webkit-scrollbar {
    display: none;
}
.scrollbar-hide {
    -ms-overflow-style: none;
    scrollbar-width: none;
}
</style>
