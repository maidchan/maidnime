<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { apiService, type Anime } from '../services/api';
import AnimePoster from './AnimePoster.vue';

const ongoingAnime = ref<Anime[]>([]);
const loading = ref(true);

onMounted(async () => {
    try {
        ongoingAnime.value = await apiService.getOngoing();
    } catch (error) {
        console.error('Failed to fetch ongoing anime:', error);
    } finally {
        loading.value = false;
    }
});
</script>

<template>
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-12">
        <div class="flex items-center justify-between mb-4">
            <h2 class="text-2xl font-bold text-white tracking-tight flex items-center gap-2">
                <span class="w-1 h-6 bg-primary rounded-full block"></span>
                Anime Ongoing
            </h2>
            <div class="flex gap-2">
                <button class="size-8 rounded-full bg-surface-dark hover:bg-surface-dark-hover flex items-center justify-center text-white transition-colors">
                    <span class="material-symbols-outlined text-lg">chevron_left</span>
                </button>
                <button class="size-8 rounded-full bg-surface-dark hover:bg-surface-dark-hover flex items-center justify-center text-white transition-colors">
                    <span class="material-symbols-outlined text-lg">chevron_right</span>
                </button>
            </div>
        </div>
        
        <div v-if="loading" class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
            <div v-for="i in 6" :key="i" class="aspect-2/3 rounded-lg bg-surface-dark animate-pulse"></div>
        </div>
        <div v-else class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
            <AnimePoster 
                v-for="anime in ongoingAnime" 
                :key="anime.id" 
                :anime="{
                    id: anime.id,
                    title: anime.title,
                    image: anime.image,
                    currentEpisode: anime.last_episode ? parseInt(anime.last_episode.replace(/\D/g, '')) || 0 : 0,
                    type: 'Series'
                }" 
            />
        </div>
    </section>
</template>
