<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { apiService } from '../services/api';

const genres = ref<string[]>([]);
const loading = ref(true);
const router = useRouter();

onMounted(async () => {
    try {
        genres.value = await apiService.getGenres();
    } catch (error) {
        console.error('Failed to fetch genres:', error);
    } finally {
        loading.value = false;
    }
});

const searchByGenre = (genre: string) => {
    router.push(`/browse?genre=${encodeURIComponent(genre)}`);
};
</script>

<template>
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-16 mb-8">
        <h2 class="text-2xl font-bold text-white tracking-tight mb-6">Browse by Genre</h2>
        
        <div v-if="loading" class="flex flex-wrap gap-3">
            <div v-for="i in 10" :key="i" class="w-24 h-10 rounded-full bg-surface-dark animate-pulse"></div>
        </div>
        <div v-else class="flex flex-wrap gap-3">
            <button 
                v-for="genre in genres" 
                :key="genre"
                @click="searchByGenre(genre)"
                class="px-6 py-2.5 rounded-full bg-surface-dark border border-white/5 text-gray-300 hover:bg-primary hover:text-white hover:border-primary transition-all text-sm font-medium"
            >
                {{ genre }}
            </button>
        </div>
    </section>
</template>
