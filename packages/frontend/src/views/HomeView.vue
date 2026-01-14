<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { apiService, type Anime } from '../services/api';
import HeroSection from '../components/HeroSection.vue';
import LatestEpisodes from '../components/LatestEpisodes.vue';
import AnimeOngoing from '../components/AnimeOngoing.vue';
import AnimePopular from '../components/AnimePopular.vue';
import GenreSection from '../components/GenreSection.vue';

const featuredAnime = ref<Anime | undefined>(undefined);

onMounted(async () => {
    try {
        const homeData = await apiService.getHome();
        if (homeData.length > 0) {
            featuredAnime.value = homeData[Math.floor(Math.random() * Math.min(homeData.length, 5))];
        }
    } catch (error) {
        console.error('Failed to fetch home data for hero:', error);
    }
});
</script>

<template>
  <main class="pt-0">
    <HeroSection :anime="featuredAnime" />
    <LatestEpisodes />
    <AnimeOngoing />
    <AnimePopular />
    <GenreSection />
  </main>
</template>
