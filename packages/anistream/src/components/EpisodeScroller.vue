<script setup lang="ts">
interface Episode {
    id: number;
    title: string;
    thumbnail: string;
    isCurrent?: boolean;
}

defineProps<{
    episodes: Episode[];
}>();
</script>

<template>
    <div class="mt-8 mb-12">
        <div class="flex items-center justify-between mb-4 px-1">
            <h3 class="text-white text-lg font-bold">Episode Lainnya</h3>
            <a href="#" class="text-primary text-sm font-semibold hover:text-white transition-colors">View All</a>
        </div>
        
        <div class="relative w-full group/scroll">
            <!-- Scroll Buttons (Visual only for now) -->
            <button class="absolute left-0 top-1/2 -translate-y-1/2 -translate-x-1/2 z-10 size-10 bg-[#362348] rounded-full text-white shadow-xl flex items-center justify-center opacity-0 group-hover/scroll:opacity-100 hover:bg-primary transition-all disabled:opacity-0 hidden lg:flex">
                <span class="material-symbols-outlined">chevron_left</span>
            </button>
            <button class="absolute right-0 top-1/2 -translate-y-1/2 translate-x-1/2 z-10 size-10 bg-[#362348] rounded-full text-white shadow-xl flex items-center justify-center opacity-0 group-hover/scroll:opacity-100 hover:bg-primary transition-all hidden lg:flex">
                <span class="material-symbols-outlined">chevron_right</span>
            </button>

            <!-- Horizontal Scroll Container -->
            <div class="flex gap-4 overflow-x-auto pb-4 scrollbar-hide snap-x snap-mandatory">
                <div v-for="ep in episodes" :key="ep.id" 
                    class="flex-none w-60 group snap-start cursor-pointer"
                    :class="{ 'relative': ep.isCurrent }"
                >
                    <!-- Current Episode Style -->
                    <div v-if="ep.isCurrent" class="relative aspect-video rounded-lg overflow-hidden mb-2 border-2 border-primary shadow-[0_0_15px_rgba(127,19,236,0.3)]">
                        <div class="absolute inset-0 bg-cover bg-center" :style="{ backgroundImage: `url('${ep.thumbnail}')` }"></div>
                        <div class="absolute inset-0 bg-black/20"></div>
                        <div class="absolute inset-0 flex items-center justify-center">
                            <div class="px-3 py-1 bg-black/60 backdrop-blur rounded text-white text-xs font-bold flex items-center gap-1">
                                <span class="size-2 bg-primary rounded-full animate-pulse"></span>
                                Playing
                            </div>
                        </div>
                        <div class="absolute bottom-0 inset-x-0 h-1 bg-white/30">
                            <div class="h-full w-1/3 bg-primary"></div>
                        </div>
                    </div>

                    <!-- Other Episode Style -->
                    <div v-else class="relative aspect-video rounded-lg overflow-hidden mb-2 border border-[#362348] group-hover:border-primary/50 transition-colors">
                        <div class="absolute inset-0 bg-cover bg-center transition-transform duration-500 group-hover:scale-110" :style="{ backgroundImage: `url('${ep.thumbnail}')` }"></div>
                        <div class="absolute inset-0 bg-black/40 group-hover:bg-black/20 transition-colors"></div>
                        <div class="absolute bottom-2 left-2 px-2 py-0.5 bg-black/60 backdrop-blur rounded text-white text-xs font-bold">Ep {{ ep.id }}</div>
                        <div class="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
                            <div class="size-10 bg-primary/90 rounded-full flex items-center justify-center shadow-lg">
                                <span class="material-symbols-outlined text-white text-[24px]">play_arrow</span>
                            </div>
                        </div>
                    </div>

                    <h4 :class="ep.isCurrent ? 'text-primary font-bold' : 'text-white font-medium group-hover:text-primary transition-colors'" class="text-sm truncate">
                        {{ ep.title }}
                    </h4>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.scrollbar-hide::-webkit-scrollbar {
    display: none;
}
.scrollbar-hide {
    -ms-overflow-style: none;
    scrollbar-width: none;
}
</style>
