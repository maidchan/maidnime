<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { apiService, type Anime } from '../services/api';

const props = defineProps<{
    anime?: Anime;
}>();

const router = useRouter();
const isInList = ref(false);

const title = computed(() => props.anime?.title || 'Maidnime');
const image = computed(() => props.anime?.image || 'https://images2.alphacoders.com/131/1313627.jpg');
const description = computed(() => props.anime?.last_episode ? `Episode Terbaru: ${props.anime.last_episode}` : 'Streaming anime subtitle Indonesia gratis tanpa iklan.');
const genres = computed(() => props.anime?.genre || ['Action', 'Adventure']);

const watchNow = async () => {
    try {
        // If we have a last episode id on the prop, prefer it
        if (props.anime?.last_episode_id) {
            router.push(`/watch?id=${props.anime.last_episode_id}`);
            return;
        }

        // If we have an anime id, fetch latest detail to get the most recent episode id
        if (props.anime?.id) {
            // If the id looks like an episode id, navigate directly
            if (typeof props.anime.id === 'string' && props.anime.id.includes('episode')) {
                router.push(`/watch?id=${props.anime.id}`);
                return;
            }

            // Fetch fresh anime detail to ensure we have the latest update
            try {
                const detail = await apiService.getAnimeDetail(String(props.anime.id));
                if (detail && detail.last_episode_id) {
                    router.push(`/watch?id=${detail.last_episode_id}`);
                    return;
                }
            } catch (err) {
                console.error('Failed to fetch anime detail for watch button:', err);
            }

            // Fallback: navigate to anime page
            router.push(`/anime/${props.anime.id}`);
        }
    } catch (error) {
        console.error('watchNow error:', error);
    }
};

const toggleList = () => {
    if (!props.anime) return;
    isInList.value = !isInList.value;
    const list = JSON.parse(localStorage.getItem('watchlist') || '[]');
    if (isInList.value) {
        if (!list.find((item: any) => item.id === props.anime?.id)) {
            list.push({ id: props.anime.id, title: props.anime.title, image: props.anime.image });
        }
    } else {
        const index = list.findIndex((item: any) => item.id === props.anime?.id);
        if (index > -1) list.splice(index, 1);
    }
    localStorage.setItem('watchlist', JSON.stringify(list));
};

onMounted(() => {
    if (props.anime) {
        const list = JSON.parse(localStorage.getItem('watchlist') || '[]');
        isInList.value = !!list.find((item: any) => item.id === props.anime?.id);
    }
});
</script>

<template>
    <div class="relative w-full h-[500px] md:h-[600px] overflow-hidden group">
        <!-- Background Image -->
        <div 
            class="absolute inset-0 bg-cover bg-center transition-transform duration-700 group-hover:scale-105"
            :style="`background-image: url('${image}');`"
        ></div>

        <!-- Gradient Overlay -->
        <div class="absolute inset-0 bg-linear-to-t from-background-dark via-background-dark/60 to-transparent"></div>
        <div class="absolute inset-0 bg-linear-to-r from-background-dark via-background-dark/40 to-transparent"></div>

        <!-- Content -->
        <div class="relative h-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex items-center">
            <div class="max-w-2xl flex flex-col gap-4 md:gap-6 mt-16 md:mt-20">
                <div class="flex flex-wrap items-center gap-2 md:gap-3">
                    <span 
                        v-for="genre in genres.slice(0, 3)" 
                        :key="genre"
                        class="px-2 py-1 bg-white/10 text-white text-xs font-medium uppercase rounded backdrop-blur-sm border border-white/10"
                    >
                        {{ genre }}
                    </span>
                </div>
                
                <h1 class="text-3xl sm:text-4xl md:text-5xl lg:text-6xl font-extrabold text-white leading-tight tracking-tight text-shadow-sm" v-html="title.replace(':', ':<br/>')">
                </h1>
                
                <p class="text-gray-300 text-base md:text-lg line-clamp-2 md:line-clamp-3 max-w-xl text-shadow-sm">
                    {{ description }}
                </p>
                
                <div class="flex flex-col sm:flex-row items-stretch sm:items-center gap-3 md:gap-4 pt-2">
                    <button @click="watchNow" class="flex items-center justify-center gap-2 bg-primary hover:bg-primary/90 text-white px-6 md:px-8 py-3 rounded-lg font-bold transition-transform active:scale-95 shadow-lg shadow-primary/25 min-h-[44px]">
                        <span class="material-symbols-outlined fill-1">play_arrow</span>
                        Watch Now
                    </button>
                    <button 
                        @click="toggleList"
                        class="flex items-center justify-center gap-2 bg-white/10 hover:bg-white/20 text-white px-5 md:px-6 py-3 rounded-lg font-medium backdrop-blur-sm transition-colors border border-white/10 min-h-[44px]"
                        :class="{ 'text-primary': isInList }"
                    >
                        <span class="material-symbols-outlined" :class="{ 'fill-1': isInList }">{{ isInList ? 'done' : 'add' }}</span>
                        {{ isInList ? 'In My List' : 'My List' }}
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
