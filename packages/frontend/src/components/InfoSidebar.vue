<script setup lang="ts">
interface Character {
    name: string;
    role: string;
    image: string;
}

defineProps<{
    studio: string;
    aired: string;
    season: string;
    duration: string;
    quality: string;
    characters: Character[];
}>();
</script>

<template>
  <div class="space-y-8">
    <!-- Info Grid -->
    <div class="bg-surface-dark/30 p-6 rounded-xl border border-white/5">
        <h4 class="text-white font-bold mb-4">Information</h4>
        <div class="space-y-4">
            <div class="flex justify-between items-center pb-3 border-b border-white/5">
                <span class="text-sm text-gray-400">Studio</span>
                <span class="text-sm text-primary font-medium">{{ studio }}</span>
            </div>
            <div class="flex justify-between items-center pb-3 border-b border-white/5">
                <span class="text-sm text-gray-400">Aired</span>
                <span class="text-sm text-white">{{ aired }}</span>
            </div>
            <div class="flex justify-between items-center pb-3 border-b border-white/5">
                <span class="text-sm text-gray-400">Season</span>
                <span class="text-sm text-white">{{ season }}</span>
            </div>
            <div class="flex justify-between items-center pb-3 border-b border-white/5">
                <span class="text-sm text-gray-400">Duration</span>
                <span class="text-sm text-white">{{ duration }}</span>
            </div>
            <div class="flex justify-between items-center">
                <span class="text-sm text-gray-400">Quality</span>
                <span class="text-xs font-bold text-background-dark bg-white px-2 py-0.5 rounded">{{ quality }}</span>
            </div>
        </div>
    </div>

    <!-- Main Characters -->
    <div>
        <h4 class="text-white font-bold mb-4">Main Characters</h4>
        <div class="grid grid-cols-1 gap-3">
            <div v-for="char in characters" :key="char.name" class="flex items-center gap-3 p-2 rounded-lg hover:bg-white/5 transition-colors">
                <div class="w-10 h-10 rounded-full bg-gray-700 overflow-hidden">
                    <img :src="char.image" :alt="char.name" class="w-full h-full object-cover" />
                </div>
                <div>
                    <div class="text-sm font-bold text-white">{{ char.name }}</div>
                    <div class="text-xs text-gray-400">{{ char.role }}</div>
                </div>
            </div>
        </div>
    </div>

    <!-- Promo / Ad Placeholder -->
    <div class="rounded-xl overflow-hidden relative h-40 group cursor-pointer">
        <div class="absolute inset-0 bg-gradient-to-r from-primary to-purple-800"></div>
        <div class="absolute inset-0 flex flex-col items-center justify-center p-4 text-center">
            <span class="material-symbols-outlined text-white/50 text-4xl mb-2">diamond</span>
            <p class="text-white font-bold">Go Premium</p>
            <p class="text-white/70 text-xs mt-1">Watch without ads in 4K</p>
        </div>
    </div>
  </div>
</template>
