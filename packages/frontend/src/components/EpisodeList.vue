<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

interface Episode {
    id: string | number;
    title: string;
    description: string;
    thumbnail: string;
    duration: string;
}

defineProps<{
    episodes: Episode[];
}>();

const router = useRouter();
const showingAll = ref(false);

const goToWatch = (episodeId: string | number) => {
    router.push({ name: 'watch', query: { id: episodeId.toString() } });
};
</script>

<template>
  <section>
    <div class="flex items-center justify-between mb-6">
      <h3 class="text-xl font-bold text-white flex items-center gap-2">
        Episodes
        <span class="text-sm font-normal text-gray-400 bg-white/5 px-2 py-0.5 rounded-full">{{ episodes.length }}</span>
      </h3>
      <!-- Episode Filter/Sort -->
      <div class="flex bg-surface-dark rounded-lg p-1 border border-white/5">
        <button class="p-1.5 md:p-1.5 rounded bg-white/10 text-white active:bg-white/10 min-h-[44px] min-w-[44px] flex items-center justify-center">
          <span class="material-symbols-outlined text-lg md:text-[20px]">grid_view</span>
        </button>
        <button class="p-1.5 md:p-1.5 rounded text-gray-400 hover:text-white active:text-white min-h-[44px] min-w-[44px] flex items-center justify-center">
          <span class="material-symbols-outlined text-lg md:text-[20px]">list</span>
        </button>
      </div>
    </div>

    <div class="space-y-3">
        <div 
            v-for="episode in (showingAll ? episodes : episodes.slice(0, 3))" 
            :key="episode.id"
            @click="goToWatch(episode.id)"
            class="group flex flex-col sm:flex-row items-start sm:items-center gap-3 sm:gap-4 p-3 rounded-lg bg-surface-dark/40 hover:bg-surface-light border border-transparent hover:border-primary/30 transition-all cursor-pointer active:bg-surface-light"
        >
            <!-- Thumbnail -->
            <div class="relative w-full sm:w-32 aspect-video sm:aspect-video rounded-md overflow-hidden shrink-0">
                <img :src="episode.thumbnail" :alt="episode.title" class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" />
                <div class="absolute inset-0 bg-black/40 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
                    <div class="w-8 h-8 md:w-10 md:h-10 rounded-full bg-primary flex items-center justify-center text-white">
                        <span class="material-symbols-outlined text-[20px] md:text-[24px] ml-0.5">play_arrow</span>
                    </div>
                </div>
                <div class="absolute bottom-1 right-1 px-1.5 py-0.5 bg-black/80 text-[10px] font-bold rounded text-white">{{ episode.duration }}</div>
            </div>
            
            <!-- Info -->
            <div class="flex-1 min-w-0 w-full sm:w-auto">
                <div class="flex items-center justify-between mb-1 gap-2">
                    <h4 class="text-white font-bold text-sm md:text-base truncate group-hover:text-primary transition-colors flex-1">{{ episode.title }}</h4>
                    <span class="text-xs text-gray-500 font-mono shrink-0">EP {{ episode.id.toString().padStart(2, '0') }}</span>
                </div>
                <p class="text-xs md:text-sm text-gray-400 line-clamp-2 sm:line-clamp-1">{{ episode.description }}</p>
            </div>
        </div>
    </div>

    <button 
        v-if="episodes.length > 3"
        @click="showingAll = !showingAll"
        class="w-full mt-4 py-3 text-sm md:text-base text-gray-400 hover:text-white bg-surface-dark hover:bg-surface-light rounded-lg transition-colors font-medium flex items-center justify-center gap-2 min-h-[44px] active:bg-surface-light"
    >
        {{ showingAll ? 'Show less' : `Show all ${episodes.length} episodes` }}
        <span class="material-symbols-outlined text-lg" :class="{ 'rotate-180': showingAll }">expand_more</span>
    </button>
  </section>
</template>
