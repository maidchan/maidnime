<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { apiService, type AnimeDetail } from '../services/api';
import DetailHero from '../components/DetailHero.vue';
import EpisodeList from '../components/EpisodeList.vue';
import InfoSidebar from '../components/InfoSidebar.vue';
import RelatedAnime from '../components/RelatedAnime.vue';

const route = useRoute();
const id = route.params.id as string;

const animeDetail = ref<AnimeDetail | null>(null);
const relatedAnime = ref<AnimeDetail['genre'] extends any ? any[] : any[]>([]); // Use Anime[] from api.ts
const loading = ref(true);

onMounted(async () => {
    try {
        const [detail, homeData] = await Promise.all([
            apiService.getAnimeDetail(id),
            apiService.getHome()
        ]);
        animeDetail.value = detail;
        relatedAnime.value = homeData;
    } catch (error) {
        console.error('Failed to fetch anime detail:', error);
    } finally {
        loading.value = false;
    }
});

const animeData = computed(() => {
    if (!animeDetail.value) return null;
    return {
        id: animeDetail.value.id,
        title: animeDetail.value.title,
        poster: animeDetail.value.image,
        banner: animeDetail.value.image,
        description: animeDetail.value.synopsis,
        rating: animeDetail.value.rating ? parseFloat(animeDetail.value.rating.split('/')[0]) || 0 : 0,
        year: 2024,
        episodes: animeDetail.value.episodes.length,
        status: animeDetail.value.status || 'Ongoing',
        genre: animeDetail.value.genre || [],
        latestEpisodeId: animeDetail.value.episodes.length > 0 ? animeDetail.value.episodes[0].id : undefined,
        studio: 'Ufotable',
        aired: 'N/A',
        season: 'N/A',
        duration: '24 min',
        quality: 'HD',
        characters: []
    };
});

const episodes = computed(() => {
    if (!animeDetail.value) return [];
    return animeDetail.value.episodes.map(extEp => ({
        id: extEp.id,
        title: extEp.title,
        description: '',
        thumbnail: animeDetail.value?.image || '',
        duration: extEp.date
    }));
});
</script>

<template>
  <main class="w-full">
    <div v-if="loading" class="min-h-screen flex items-center justify-center">
        <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
    </div>
    
    <template v-else-if="animeData">
        <!-- Hero Section -->
        <DetailHero v-bind="animeData" />

        <!-- Content Section -->
        <div class="max-w-[1280px] mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
            <!-- Main Column (Left) -->
            <div class="lg:col-span-8 space-y-12">
                <!-- Mobile Poster -->
                <div class="md:hidden w-full flex gap-4 mb-8">
                    <div class="w-32 aspect-2/3 rounded-lg overflow-hidden shrink-0 shadow-lg border border-white/10">
                        <img :src="animeData.poster" :alt="animeData.title" class="w-full h-full object-cover" />
                    </div>
                    <div class="flex flex-col justify-center gap-2">
                        <div class="text-sm text-gray-400">Studio {{ animeData.studio }}</div>
                        <div class="text-sm text-gray-400">{{ animeData.aired }}</div>
                    </div>
                </div>

                <!-- Synopsis -->
                <section>
                    <h3 class="text-xl font-bold text-white mb-4 flex items-center gap-2">Synopsis</h3>
                    <div class="bg-surface-dark/50 p-6 rounded-xl border border-white/5">
                        <p class="text-gray-300 leading-relaxed">{{ animeData.description }}</p>
                    </div>
                </section>

                <!-- Episodes List -->
                <EpisodeList :episodes="episodes" />
            </div>

            <!-- Sidebar (Right) -->
            <div class="lg:col-span-4 space-y-8">
                <InfoSidebar v-bind="animeData" />
            </div>
        </div>

        <!-- Related Anime -->
        <RelatedAnime :anime-list="relatedAnime" />
        </div>
    </template>

    <div v-else class="min-h-screen flex items-center justify-center text-white">
        Anime not found.
    </div>
  </main>
</template>
