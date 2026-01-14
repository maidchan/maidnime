<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
  isPlaying: boolean;
  currentTime: number;
  duration: number;
  buffered: number; // 0 to 1 scale
}>();

const emit = defineEmits<{
  (e: 'togglePlay'): void;
  (e: 'toggleMute'): void;
  (e: 'toggleFullscreen'): void;
}>();

const formatTime = (seconds: number) => {
  const m = Math.floor(seconds / 60);
  const s = Math.floor(seconds % 60);
  return `${m.toString().padStart(2, '0')}:${s.toString().padStart(2, '0')}`;
};

const progressPercent = computed(() => {
  if (!props.duration) return 0;
  return (props.currentTime / props.duration) * 100;
});
</script>

<template>
  <div class="absolute bottom-0 left-0 right-0 p-4 pb-5 opacity-100 flex flex-col gap-2 transition-opacity duration-300 bg-gradient-to-t from-black/90 to-transparent">
    <!-- Progress Bar -->
    <div class="group/progress relative h-1.5 w-full bg-white/20 rounded-full cursor-pointer hover:h-2 transition-all">
      <div 
        class="absolute top-0 left-0 h-full bg-primary rounded-full relative"
        :style="{ width: `${progressPercent}%` }"
      >
        <div class="absolute right-0 top-1/2 -translate-y-1/2 translate-x-1/2 size-3 bg-white rounded-full scale-0 group-hover/progress:scale-100 transition-transform shadow"></div>
      </div>
      <!-- Buffer -->
      <div 
        class="absolute top-0 left-0 h-full bg-white/30 rounded-full -z-10"
        :style="{ width: `${buffered * 100}%` }"
      ></div> 
    </div>

    <!-- Buttons Row -->
    <div class="flex items-center justify-between mt-1">
      <div class="flex items-center gap-4">
        <button 
          @click="emit('togglePlay')"
          class="text-white hover:text-primary transition-colors"
        >
          <span class="material-symbols-outlined text-[28px] fill-1">{{ isPlaying ? 'pause' : 'play_arrow' }}</span>
        </button>
        
        <button class="text-white hover:text-primary transition-colors group/vol relative flex items-center gap-2">
          <span class="material-symbols-outlined text-[24px]">volume_up</span>
          <div class="w-0 overflow-hidden group-hover/vol:w-20 transition-all duration-300">
            <div class="h-1 bg-white/30 rounded-full w-20 ml-1">
              <div class="h-full w-2/3 bg-white rounded-full"></div>
            </div>
          </div>
        </button>
        
        <div class="text-white/80 text-sm font-medium font-mono">
            {{ formatTime(currentTime) }} <span class="text-white/40">/</span> {{ formatTime(duration) }}
        </div>
      </div>

      <div class="flex items-center gap-4">
        <button class="text-white hover:text-primary transition-colors" title="Subtitles/CC">
          <span class="material-symbols-outlined text-[24px]">closed_caption</span>
        </button>
        <button class="text-white hover:text-primary transition-colors" title="Settings">
          <span class="material-symbols-outlined text-[24px]">settings</span>
        </button>
        <button 
            @click="emit('toggleFullscreen')"
            class="text-white hover:text-primary transition-colors" title="Fullscreen">
          <span class="material-symbols-outlined text-[28px]">fullscreen</span>
        </button>
      </div>
    </div>
  </div>
</template>
