<script setup lang="ts">
import { ref, onMounted } from 'vue';

const props = defineProps<{
  id: string;
  title: string;
  poster: string;
  banner: string;
  description: string;
  rating: number;
  year: number;
  episodes: number;
  status: string;
  genre: string[];
  trailer?: string;
  latestEpisodeId?: string;
}>();

const isInList = ref(false);

const toggleList = () => {
    isInList.value = !isInList.value;
    // Simple local storage mock for demonstration
    const list = JSON.parse(localStorage.getItem('watchlist') || '[]');
    if (isInList.value) {
        if (!list.find((item: any) => item.id === props.id)) {
            list.push({ id: props.id, title: props.title, image: props.poster });
        }
    } else {
        const index = list.findIndex((item: any) => item.id === props.id);
        if (index > -1) list.splice(index, 1);
    }
    localStorage.setItem('watchlist', JSON.stringify(list));
};

onMounted(() => {
    const list = JSON.parse(localStorage.getItem('watchlist') || '[]');
    isInList.value = !!list.find((item: any) => item.id === props.id);
});
</script>

<template>
  <div class="relative w-full h-[600px] lg:h-[700px]">
    <!-- Background Banner -->
    <div class="absolute inset-0 w-full h-full">
      <div 
        class="w-full h-full bg-cover bg-center" 
        :style="{ backgroundImage: `url('${banner}')` }"
      ></div>
      <!-- Gradient Overlays -->
      <div class="absolute inset-0 bg-linear-to-t from-background-dark via-background-dark/80 to-transparent"></div>
      <div class="absolute inset-0 bg-linear-to-r from-background-dark/90 via-background-dark/40 to-transparent"></div>
    </div>

    <!-- Hero Content -->
    <div class="relative h-full max-w-[1280px] mx-auto px-4 sm:px-6 lg:px-8 flex flex-col justify-end pb-8 md:pb-12 pt-20 md:pt-24">
      <div class="flex flex-col md:flex-row gap-6 md:gap-8 items-end">
        <!-- Poster Card (Floating) -->
        <div class="hidden md:block shrink-0 relative group">
          <div class="w-56 lg:w-64 aspect-2/3 rounded-xl overflow-hidden shadow-2xl border-2 border-white/10 relative z-10 transform transition-transform group-hover:scale-[1.02]">
            <img :src="poster" :alt="title" class="w-full h-full object-cover" />
          </div>
          <!-- Glow effect behind poster -->
          <div class="absolute -inset-4 bg-primary/30 blur-2xl -z-10 rounded-full opacity-0 group-hover:opacity-100 transition-opacity duration-500"></div>
        </div>

        <!-- Info -->
        <div class="flex-1 space-y-4 md:space-y-6 w-full">
          <!-- Badges -->
          <div class="flex flex-wrap gap-2">
            <span v-for="g in genre" :key="g" class="px-2 md:px-3 py-1 rounded-full bg-primary/20 border border-primary/30 text-primary-200 text-xs font-bold uppercase tracking-wider">{{ g }}</span>
          </div>

          <!-- Title -->
          <h1 class="text-3xl sm:text-4xl md:text-5xl lg:text-6xl font-extrabold text-white leading-tight tracking-tight drop-shadow-lg">
            {{ title }}
          </h1>

          <!-- Meta Data -->
          <div class="flex flex-wrap items-center gap-2 md:gap-4 text-xs sm:text-sm md:text-base text-gray-300 font-medium">
            <div class="flex items-center text-yellow-400 gap-1">
              <span class="material-symbols-outlined text-base md:text-lg fill-current">star</span>
              <span class="font-bold text-white">{{ rating }}</span>
            </div>
            <span class="w-1 h-1 rounded-full bg-gray-500 hidden sm:block"></span>
            <span>{{ year }}</span>
            <span class="w-1 h-1 rounded-full bg-gray-500 hidden sm:block"></span>
            <span>{{ episodes }} Episodes</span>
            <span class="w-1 h-1 rounded-full bg-gray-500 hidden sm:block"></span>
            <span class="text-green-400">{{ status }}</span>
            <span class="w-1 h-1 rounded-full bg-gray-500 hidden sm:block"></span>
            <span class="border border-white/20 px-1.5 rounded text-xs">TV-MA</span>
          </div>

          <!-- Actions -->
          <div class="flex flex-col sm:flex-row items-stretch sm:items-center gap-3 md:gap-4 pt-2">
            <router-link :to="latestEpisodeId ? `/watch?id=${latestEpisodeId}` : `/anime/${id}`" class="h-12 px-6 md:px-8 bg-primary hover:bg-primary-hover text-white rounded-lg font-bold flex items-center justify-center gap-2 transition-all shadow-lg shadow-primary/25 cursor-pointer min-h-[44px]">
              <span class="material-symbols-outlined fill">play_arrow</span>
              Watch Now
            </router-link>
            <a v-if="trailer" :href="trailer" target="_blank" class="h-12 px-6 md:px-8 bg-white/10 hover:bg-white/20 text-white rounded-lg font-bold flex items-center justify-center gap-2 backdrop-blur-sm transition-all border border-white/10 min-h-[44px]">
              <span class="material-symbols-outlined">movie</span>
              Trailer
            </a>
            <button 
                @click="toggleList"
                class="h-12 px-5 md:px-6 bg-white/10 hover:bg-white/20 text-white rounded-lg font-bold flex items-center justify-center gap-2 backdrop-blur-sm transition-all border border-white/10 min-h-[44px]"
                :class="{ 'text-primary': isInList }"
            >
              <span class="material-symbols-outlined" :class="{ 'fill-1': isInList }">{{ isInList ? 'done' : 'add' }}</span>
              {{ isInList ? 'In My List' : 'My List' }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
