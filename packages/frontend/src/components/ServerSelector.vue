<script setup lang="ts">
import { ref, watch } from 'vue';

interface Server {
    quality: string;
    url: string;
}

const props = defineProps<{
  animeTitle: string;
  episodeNumber?: number | string;
  episodeTitle: string;
  servers: Server[];
}>();

const emit = defineEmits<{
    (e: 'select-server', url: string): void
}>();

const selectedServer = ref(props.servers[0]?.url || '');

watch(() => props.servers, (newServers) => {
    if (newServers.length > 0 && !selectedServer.value) {
        selectedServer.value = newServers[0].url;
    }
}, { immediate: true });

watch(selectedServer, (newUrl) => {
    emit('select-server', newUrl);
});
</script>

<template>
  <div class="flex flex-col gap-6">
    <!-- Title Section -->
    <div class="flex flex-col gap-2">
      <h1 class="text-white text-2xl md:text-3xl font-bold leading-tight tracking-tight">{{ animeTitle }}</h1>
      <p class="text-[#ad92c9] text-base md:text-lg font-medium">
        Episode {{ episodeNumber }} <span class="text-white/40 mx-2">•</span> <span class="text-white/90">{{ episodeTitle }}</span>
      </p>
    </div>

    <!-- Server Selection -->
    <div class="bg-[#241830] p-4 rounded-xl border border-[#362348]">
      <div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
        <div class="flex items-center gap-2 text-white/80 text-sm font-medium">
          <span class="material-symbols-outlined text-primary">dns</span>
          <span>Stream Source:</span>
        </div>
        
        <!-- Segmented Buttons -->
        <div class="flex flex-wrap gap-2 w-full sm:w-auto">
          <label 
            v-for="server in servers" 
            :key="server.url"
            class="cursor-pointer group"
          >
            <input 
                type="radio" 
                name="server" 
                class="peer sr-only" 
                :value="server.url"
                v-model="selectedServer"
            >
            <div class="px-4 py-2 rounded-lg bg-[#362348] text-[#ad92c9] text-sm font-medium transition-all peer-checked:bg-primary peer-checked:text-white peer-checked:shadow-lg peer-checked:shadow-primary/30 group-hover:bg-[#462d5d]">
               {{ server.quality }}
            </div>
          </label>
        </div>
      </div>
    </div>

    <!-- Main Navigation Controls (Prev/Next) -->
    <div class="flex items-center justify-between gap-4 mt-2">
      <button class="flex items-center gap-2 px-6 py-3 rounded-lg bg-[#362348] hover:bg-[#462d5d] text-white transition-all w-full sm:w-auto justify-center group">
        <span class="material-symbols-outlined group-hover:-translate-x-1 transition-transform">arrow_back</span>
        <span class="font-medium">Prev Ep</span>
      </button>
      
      <div class="hidden sm:flex items-center gap-2 text-[#ad92c9] text-sm">
        <span class="material-symbols-outlined text-[20px]">list</span>
        <span>Episode List</span>
      </div>
      
      <button class="flex items-center gap-2 px-6 py-3 rounded-lg bg-white text-primary hover:bg-white/90 transition-all w-full sm:w-auto justify-center font-bold shadow-lg shadow-white/10 group">
        <span>Next Ep</span>
        <span class="material-symbols-outlined group-hover:translate-x-1 transition-transform">arrow_forward</span>
      </button>
    </div>
  </div>
</template>
