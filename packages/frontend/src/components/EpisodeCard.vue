<script setup lang="ts">
import { useRouter } from 'vue-router';

interface EpisodeProps {
    id: string | number;
    animeTitle: string;
    episodeNumber?: number | string;
    title: string;
    image: string;
    duration?: string;
    uploadTime?: string;
}

const props = defineProps<{
    episode: EpisodeProps
}>();

const router = useRouter();

const handleClick = () => {
    // If it's an anime ID (from home), go to detail. If it's an episode ID, go to watch.
    // For now, let's assume if it contains 'episode' or matches episode pattern, go to watch.
    // Actually, Otakudesu episode IDs usually look like 'anime-name-episode-X'
    if (typeof props.episode.id === 'string' && props.episode.id.includes('episode')) {
        router.push({ name: 'watch', query: { id: props.episode.id } });
    } else {
        router.push({ name: 'anime-detail', params: { id: props.episode.id } });
    }
};
</script>

<template>
    <div 
        @click="handleClick"
        class="min-w-[280px] md:min-w-[320px] rounded-xl bg-surface-dark overflow-hidden group hover:ring-2 hover:ring-primary/50 transition-all duration-300 cursor-pointer"
    >
        <div class="relative aspect-video bg-gray-800 overflow-hidden">
            <div 
                class="w-full h-full bg-cover bg-center group-hover:scale-105 transition-transform duration-500"
                :style="`background-image: url('${episode.image}');`"
            ></div>
            
            <div class="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
                <div class="bg-primary rounded-full p-3 text-white shadow-lg transform scale-0 group-hover:scale-100 transition-transform duration-300">
                    <span class="material-symbols-outlined text-3xl">play_arrow</span>
                </div>
            </div>
            
            <span v-if="episode.duration" class="absolute bottom-2 right-2 bg-black/80 text-white text-xs px-2 py-1 rounded font-medium">
                {{ episode.duration }}
            </span>
        </div>
        
        <div class="p-4">
            <div class="flex items-center justify-between mb-1">
                <span v-if="episode.episodeNumber" class="text-xs font-semibold text-primary">Ep {{ episode.episodeNumber }}</span>
                <span v-else-if="episode.title.includes('Episode')" class="text-xs font-semibold text-primary">{{ episode.title.split(' ').pop() }}</span>
                <span v-else class="text-xs font-semibold text-primary">NEW</span>
                <span class="text-xs text-gray-400">{{ episode.uploadTime || 'Baru' }}</span>
            </div>
            <h3 class="text-white font-semibold truncate">{{ episode.title }}</h3>
            <p class="text-gray-400 text-sm truncate">{{ episode.animeTitle }}</p>
        </div>
    </div>
</template>
