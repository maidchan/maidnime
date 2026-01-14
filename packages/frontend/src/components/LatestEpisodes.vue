<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { apiService, type Anime } from '../services/api';
import EpisodeCard from './EpisodeCard.vue';

const latestAnime = ref<Anime[]>([]);
const loading = ref(true);

onMounted(async () => {
    try {
        latestAnime.value = await apiService.getHome();
    } catch (error) {
        console.error('Failed to fetch latest episodes:', error);
    } finally {
        loading.value = false;
    }
});
</script>

<template>
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-8">
        <div class="flex items-center justify-between mb-4">
            <h2 class="text-2xl font-bold text-white tracking-tight flex items-center gap-2">
                <span class="w-1 h-6 bg-primary rounded-full block"></span>
                Episode Terbaru
            </h2>
            <router-link to="/browse" class="text-sm font-medium text-primary hover:text-primary/80 transition-colors">See All</router-link>
        </div>
        
        <div v-if="loading" class="flex gap-4 overflow-x-auto pb-4 no-scrollbar">
            <div v-for="i in 5" :key="i" class="min-w-[280px] aspect-video bg-white/5 animate-pulse rounded-xl"></div>
        </div>

        <div v-else class="relative group/slider">
            <div class="flex overflow-x-auto gap-4 pb-4 no-scrollbar scroll-smooth">
                <EpisodeCard 
                    v-for="anime in latestAnime" 
                    :key="anime.id" 
                    :episode="{
                        id: anime.id,
                        animeTitle: anime.title,
                        title: anime.last_episode || 'Episode Baru',
                        image: anime.image,
                        duration: '24:00',
                        uploadTime: 'Baru saja'
                    }" 
                />
            </div>
        </div>
    </section>
</template>
