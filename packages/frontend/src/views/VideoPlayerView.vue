<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { apiService, type StreamResponse } from '../services/api';
import VideoPlayer from '../components/VideoPlayer.vue';
import ServerSelector from '../components/ServerSelector.vue';

const route = useRoute();
const id = ref(route.query.id as string);

const streamData = ref<StreamResponse | null>(null);
const currentVideoSrc = ref('');
const loading = ref(true);
const isLiked = ref(false);
const isSaved = ref(false);

const handleShare = () => {
    if (navigator.share) {
        navigator.share({
            title: streamData.value?.title || 'Anistream',
            url: window.location.href
        }).catch(() => {
            alert('Sharing failed');
        });
    } else {
        navigator.clipboard.writeText(window.location.href);
        alert('Link copied to clipboard!');
    }
};

const handleReport = () => {
    alert('Thank you for reporting. Our moderators will review this episode.');
};

const fetchStream = async (episodeId: string) => {
    loading.value = true;
    try {
        streamData.value = await apiService.getEpisodeDetail(episodeId);
        if (streamData.value && streamData.value.link) {
            currentVideoSrc.value = apiService.normalizeUrl(streamData.value.link);
        } else if (streamData.value && streamData.value.qualities.length > 0) {
            currentVideoSrc.value = apiService.normalizeUrl(streamData.value.qualities[0].url);
        }
    } catch (error) {
        console.error('Failed to fetch stream detail:', error);
    } finally {
        loading.value = false;
    }
};

const handleServerSelect = (url: string) => {
    currentVideoSrc.value = apiService.normalizeUrl(url);
};

onMounted(() => {
    if (id.value) {
        fetchStream(id.value);
    }
});

const animeTitle = computed(() => streamData.value?.title || 'Loading...');

const qualities = computed(() => {
    if (!streamData.value) return [];
    return streamData.value.qualities.map(q => q.quality);
});
</script>

<template>
  <main class="flex-1 flex justify-center py-6 px-4 md:px-8 lg:px-20 pt-24">
    <div v-if="loading" class="max-w-[1280px] w-full flex items-center justify-center min-h-[400px]">
        <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-primary"></div>
    </div>

    <div v-else-if="streamData" class="max-w-[1280px] w-full flex flex-col gap-6">
      
      <!-- Video Player Section -->
      <VideoPlayer 
        :src="currentVideoSrc" 
        :title="streamData.title"
      />

      <!-- Metadata & Actions Grid -->
      <div class="grid grid-cols-1 xl:grid-cols-3 gap-6 lg:gap-10">
        
        <!-- Left Column: Title & Server Selection -->
        <div class="xl:col-span-2">
            <ServerSelector 
                :anime-title="animeTitle"
                :episode-title="streamData.title"
                :servers="streamData.qualities"
                @select-server="handleServerSelect"
            />
        </div>

        <!-- Right Column: Actions & Info -->
        <div class="xl:col-span-1 flex flex-col gap-6">
            <!-- Actions Bar -->
            <div class="bg-[#241830] rounded-xl p-4 border border-[#362348]">
                <div class="grid grid-cols-4 gap-2">
                    <button @click="isLiked = !isLiked" class="flex flex-col items-center gap-2 p-2 rounded-lg hover:bg-[#362348] transition-colors group">
                        <div class="size-10 rounded-full bg-[#362348] group-hover:bg-primary/20 flex items-center justify-center transition-colors" :class="{ 'bg-primary/20': isLiked }">
                            <span class="material-symbols-outlined text-white text-[20px] group-hover:text-primary" :class="{ 'text-primary fill-1': isLiked }">thumb_up</span>
                        </div>
                        <span class="text-xs text-[#ad92c9] font-medium">{{ isLiked ? 'Liked' : 'Like' }}</span>
                    </button>
                    <button @click="isSaved = !isSaved" class="flex flex-col items-center gap-2 p-2 rounded-lg hover:bg-[#362348] transition-colors group">
                        <div class="size-10 rounded-full bg-[#362348] group-hover:bg-primary/20 flex items-center justify-center transition-colors" :class="{ 'bg-primary/20': isSaved }">
                            <span class="material-symbols-outlined text-white text-[20px] group-hover:text-primary" :class="{ 'text-primary fill-1': isSaved }">bookmark</span>
                        </div>
                        <span class="text-xs text-[#ad92c9] font-medium">{{ isSaved ? 'Saved' : 'Save' }}</span>
                    </button>
                    <button @click="handleShare" class="flex flex-col items-center gap-2 p-2 rounded-lg hover:bg-[#362348] transition-colors group">
                        <div class="size-10 rounded-full bg-[#362348] group-hover:bg-primary/20 flex items-center justify-center transition-colors">
                            <span class="material-symbols-outlined text-white text-[20px] group-hover:text-primary">share</span>
                        </div>
                        <span class="text-xs text-[#ad92c9] font-medium">Share</span>
                    </button>
                    <button @click="handleReport" class="flex flex-col items-center gap-2 p-2 rounded-lg hover:bg-[#362348] transition-colors group">
                        <div class="size-10 rounded-full bg-[#362348] group-hover:bg-red-500/20 flex items-center justify-center transition-colors">
                            <span class="material-symbols-outlined text-white text-[20px] group-hover:text-red-400">flag</span>
                        </div>
                        <span class="text-xs text-[#ad92c9] font-medium">Report</span>
                    </button>
                </div>
            </div>

            <!-- Quality Selector -->
            <div class="flex flex-col gap-2">
                <label class="text-[#ad92c9] text-xs font-bold uppercase tracking-wider ml-1">Video Quality</label>
                <div class="relative">
                    <select class="w-full appearance-none bg-[#241830] border border-[#362348] text-white py-3 px-4 pr-8 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary/50 cursor-pointer font-medium">
                        <option v-for="q in qualities" :key="q">{{ q }}</option>
                    </select>
                    <div class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-4 text-white">
                        <span class="material-symbols-outlined">expand_more</span>
                    </div>
                </div>
            </div>
        </div>
      </div>

      <!-- Other Episodes -->
      <!-- <EpisodeScroller :episodes="episodes" /> -->

    </div>
  </main>
</template>
