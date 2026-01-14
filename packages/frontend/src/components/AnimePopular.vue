<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { apiService, type Anime } from '../services/api';
import PopularCard from './PopularCard.vue';

const popularAnime = ref<Anime[]>([]);
const loading = ref(true);

onMounted(async () => {
    try {
        popularAnime.value = await apiService.getPopular();
    } catch (error) {
        console.error('Failed to fetch popular anime:', error);
    } finally {
        loading.value = false;
    }
});
</script>

<template>
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-12">
        <h2 class="text-2xl font-bold text-white tracking-tight flex items-center gap-2 mb-6">
            <span class="w-1 h-6 bg-primary rounded-full block"></span>
            Anime Populer
        </h2>
        
        <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <div v-for="i in 3" :key="i" class="h-32 rounded-xl bg-surface-dark animate-pulse"></div>
        </div>
        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <PopularCard 
                v-for="(anime, index) in popularAnime" 
                :key="anime.id" 
                :anime="{
                    id: anime.id,
                    title: anime.title,
                    image: anime.image,
                    rating: anime.rating || 'N/A',
                    mainGenre: anime.genre && anime.genre.length > 0 ? anime.genre[0] : 'Anime',
                    description: anime.last_episode || 'Recommended',
                    genres: anime.genre || []
                }"
                :rank="index + 1"
            />
        </div>
    </section>
</template>
