<script setup lang="ts">

defineProps<{
  src?: string;
  poster?: string;
  title: string;
}>();
</script>

<template>
  <section class="relative group/player w-full">
    <!-- Cinematic Glow -->
    <div class="absolute -inset-1 bg-primary/20 blur-2xl rounded-xl opacity-50 group-hover/player:opacity-75 transition-opacity duration-500"></div>
    
    <div 
        class="relative bg-black rounded-xl overflow-hidden shadow-2xl aspect-video w-full group cursor-pointer"
    >
        <!-- Iframe for streaming -->
        <iframe 
            v-if="src"
            :src="src" 
            class="w-full h-full border-none"
            allowfullscreen
        ></iframe>
        
        <div v-else :style="{ backgroundImage: `url('${poster}')`, backgroundSize: 'cover', backgroundPosition: 'center' }" class="w-full h-full">
            <div class="absolute inset-0 bg-black/40 flex items-center justify-center">
                <div class="size-20 bg-primary/90 rounded-full flex items-center justify-center shadow-[0_0_30px_rgba(127,19,236,0.6)] backdrop-blur-sm group-hover:scale-110 transition-transform duration-300">
                    <span class="material-symbols-outlined text-white text-[48px] ml-1 fill-1">play_arrow</span>
                </div>
            </div>
        </div>

        <!-- Top Controls (Title in player) -->
        <div class="absolute top-0 left-0 right-0 p-6 opacity-0 group-hover:opacity-100 transition-opacity duration-300 pointer-events-none">
            <h3 class="text-white/90 font-medium text-lg drop-shadow-md">{{ title }}</h3>
        </div>
    </div>
  </section>
</template>
