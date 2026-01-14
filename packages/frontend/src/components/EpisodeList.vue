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
        <button class="p-1.5 rounded bg-white/10 text-white">
          <span class="material-symbols-outlined text-[20px]">grid_view</span>
        </button>
        <button class="p-1.5 rounded text-gray-400 hover:text-white">
          <span class="material-symbols-outlined text-[20px]">list</span>
        </button>
      </div>
    </div>

    <div class="space-y-3">
        <div 
            v-for="episode in (showingAll ? episodes : episodes.slice(0, 3))" 
            :key="episode.id"
            @click="goToWatch(episode.id)"
            class="group flex items-center gap-4 p-3 rounded-lg bg-surface-dark/40 hover:bg-surface-light border border-transparent hover:border-primary/30 transition-all cursor-pointer"
        >
            <!-- Thumbnail -->
            <div class="relative w-32 aspect-video rounded-md overflow-hidden shrink-0">
                <img :src="episode.thumbnail" :alt="episode.title" class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" />
                <div class="absolute inset-0 bg-black/40 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
                    <div class="w-8 h-8 rounded-full bg-primary flex items-center justify-center text-white">
                        <span class="material-symbols-outlined text-[20px] ml-0.5">play_arrow</span>
                    </div>
                </div>
                <div class="absolute bottom-1 right-1 px-1.5 py-0.5 bg-black/80 text-[10px] font-bold rounded text-white">{{ episode.duration }}</div>
            </div>
            
            <!-- Info -->
            <div class="flex-1 min-w-0">
                <div class="flex items-center justify-between mb-1">
                    <h4 class="text-white font-bold truncate group-hover:text-primary transition-colors">{{ episode.title }}</h4>
                    <span class="text-xs text-gray-500 font-mono">EP {{ episode.id.toString().padStart(2, '0') }}</span>
                </div>
                <p class="text-sm text-gray-400 line-clamp-1">{{ episode.description }}</p>
            </div>
        </div>
    </div>

    <button 
        v-if="episodes.length > 3"
        @click="showingAll = !showingAll"
        class="w-full mt-4 py-3 text-sm text-gray-400 hover:text-white bg-surface-dark hover:bg-surface-light rounded-lg transition-colors font-medium flex items-center justify-center gap-2"
    >
        {{ showingAll ? 'Show less' : `Show all ${episodes.length} episodes` }}
        <span class="material-symbols-outlined text-lg" :class="{ 'rotate-180': showingAll }">expand_more</span>
    </button>
  </section>
</template>
