<script setup lang="ts">
interface Props {
    anime: {
        id: string | number;
        title: string;
        image: string;
        rating?: string | number;
        mainGenre?: string;
        description?: string;
        genres?: string[];
    };
    rank: number;
}
defineProps<Props>();
</script>

<template>
    <router-link :to="`/anime/${anime.id}`" class="flex gap-4 bg-surface-dark rounded-xl p-3 hover:bg-surface-dark-hover transition-colors group cursor-pointer h-full">
        <div class="w-24 aspect-2/3 rounded-lg bg-gray-800 flex-shrink-0 overflow-hidden">
            <div 
                class="w-full h-full bg-cover bg-center group-hover:scale-110 transition-transform duration-500"
                :style="`background-image: url('${anime.image}');`"
            ></div>
        </div>
        
        <div class="flex flex-col justify-center flex-1 min-w-0">
            <div class="flex items-center gap-2 mb-1">
                <span v-if="anime.rating" class="text-yellow-400 material-symbols-outlined text-sm">star</span>
                <span v-if="anime.rating" class="text-white text-sm font-bold">{{ anime.rating }}</span>
                <span v-if="anime.mainGenre" class="text-gray-500 text-xs">| {{ anime.mainGenre }}</span>
            </div>
            
            <h3 class="text-white font-bold text-lg mb-1 group-hover:text-primary transition-colors truncate">
                {{ anime.title }}
            </h3>
            
            <p v-if="anime.description" class="text-gray-400 text-xs line-clamp-2">
                {{ anime.description }}
            </p>
            
            <div v-if="anime.genres && anime.genres.length" class="mt-2 flex gap-2 flex-wrap">
                <span 
                    v-for="genre in anime.genres" 
                    :key="genre"
                    class="text-[10px] bg-white/5 border border-white/10 text-gray-300 px-2 py-0.5 rounded"
                >
                    {{ genre }}
                </span>
            </div>
        </div>
        
        <div class="flex items-center justify-center pr-2">
            <span class="text-4xl font-black text-white/5">
                {{ rank.toString().padStart(2, '0') }}
            </span>
        </div>
    </router-link>
</template>
