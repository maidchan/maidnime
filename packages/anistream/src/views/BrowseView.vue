<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';
import { apiService, type Anime } from '../services/api';
import SearchInput from '../components/SearchInput.vue';
import FilterBar from '../components/FilterBar.vue';
import AnimeCard from '../components/AnimeCard.vue';

const route = useRoute();
const searchQuery = ref('');
const animeList = ref<Anime[]>([]);
const loading = ref(false);
const recentSearches = ref(['Naruto Shippuden', 'Attack on Titan', 'Jujutsu Kaisen']);

const handleSearch = async (query: string) => {
    if (!query) return;
    loading.value = true;
    try {
        animeList.value = await apiService.searchAnime(query);
        if (!recentSearches.value.includes(query)) {
            recentSearches.value.unshift(query);
            recentSearches.value = recentSearches.value.slice(0, 5);
        }
    } catch (error) {
        console.error('Search failed:', error);
    } finally {
        loading.value = false;
    }
};

watch(() => route.query.q, (newQuery) => {
    if (newQuery) {
        searchQuery.value = newQuery as string;
        handleSearch(newQuery as string);
    }
}, { immediate: true });

onMounted(() => {
    if (!route.query.q) {
        handleSearch('Naruto'); // Example initial search
    }
});
</script>

<template>
  <main class="flex-1 flex flex-col items-center w-full px-4 py-6 md:px-10 lg:px-40 min-h-screen pt-24">
    <div class="w-full max-w-[1200px] flex flex-col gap-6">
      
      <!-- Search Header -->
      <SearchInput 
        v-model="searchQuery"
        @search="handleSearch"
      />

      <!-- Recent Searches & Filters -->
      <div class="flex flex-col gap-5">
        <!-- Recent Searches -->
        <div class="flex flex-col gap-3">
          <div class="flex items-center justify-between">
            <h3 class="text-white text-lg font-bold leading-tight">Recent Searches</h3>
            <button @click="recentSearches = []" class="text-xs font-medium text-primary hover:text-white transition-colors">Clear All</button>
          </div>
          <div class="flex flex-wrap gap-3">
            <div 
                v-for="search in recentSearches" 
                :key="search"
                class="flex h-8 shrink-0 items-center justify-center gap-x-2 rounded-lg bg-surface-dark border border-white/5 pl-3 pr-2 hover:bg-white/10 cursor-pointer transition-colors group"
                @click="handleSearch(search)"
            >
              <span class="text-white text-sm font-medium leading-normal">{{ search }}</span>
              <span class="material-symbols-outlined text-gray-400 group-hover:text-white text-[16px]" @click.stop="recentSearches = recentSearches.filter(s => s !== search)">close</span>
            </div>
          </div>
        </div>

        <!-- Filters -->
         <FilterBar />
      </div>

      <!-- Results Grid -->
      <div v-if="loading" class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-x-4 gap-y-8 pb-10">
        <div v-for="i in 12" :key="i" class="aspect-2/3 rounded-xl bg-surface-dark animate-pulse"></div>
      </div>
      <div v-else class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-x-4 gap-y-8 pb-10">
        <AnimeCard 
            v-for="anime in animeList" 
            :key="anime.id"
            :id="anime.id"
            :title="anime.title"
            :image="anime.image"
            :rating="anime.rating || '0.0'"
            :year="anime.status || 'N/A'"
            :genre="anime.genre && anime.genre.length ? anime.genre[0] : 'Action'"
        />
      </div>

      <div v-if="!loading && animeList.length === 0" class="flex flex-col items-center justify-center py-20 text-gray-400 italic">
        No results found for "{{ searchQuery }}"
      </div>

    </div>
  </main>
</template>
