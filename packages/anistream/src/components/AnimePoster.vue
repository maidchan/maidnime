<script setup lang="ts">
import { useRouter } from 'vue-router';

const props = defineProps<{
    anime: {
        id: string | number;
        title: string;
        image: string;
        currentEpisode?: string | number;
        type?: string;
    }
}>();

const router = useRouter();

const goToDetail = () => {
    router.push(`/anime/${props.anime.id}`);
};
</script>

<template>
    <div class="relative group cursor-pointer" @click="goToDetail">
        <div class="aspect-2/3 rounded-lg overflow-hidden bg-gray-800 relative">
            <div 
                class="w-full h-full bg-cover bg-center group-hover:scale-110 transition-transform duration-500"
                :style="`background-image: url('${anime.image}');`"
            ></div>
            
            <div class="absolute inset-0 bg-gradient-to-t from-black via-transparent to-transparent opacity-80"></div>
            
            <div class="absolute bottom-0 left-0 p-3 w-full">
                <span class="block text-xs text-primary font-bold mb-1">
                    Ep {{ anime.currentEpisode }} out
                </span>
                <h3 class="text-white font-bold text-sm line-clamp-2 leading-tight group-hover:text-primary transition-colors">
                    {{ anime.title }}
                </h3>
            </div>
            
            <div class="absolute top-2 right-2 bg-primary text-white text-[10px] font-bold px-2 py-0.5 rounded">
                {{ anime.type }}
            </div>
        </div>
    </div>
</template>
