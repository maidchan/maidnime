<script setup lang="ts">
// In a real app, these would emit events. Visual only for now.
</script>

<template>
  <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-3 md:gap-4 pt-4 border-t border-white/5">
    <div class="flex flex-wrap items-center gap-2 md:gap-3">
        <!-- Genre Filter -->
        <button class="flex items-center gap-1.5 md:gap-2 h-10 md:h-10 px-3 md:px-4 rounded-lg bg-primary/20 text-primary border border-primary/30 hover:bg-primary/30 active:bg-primary/30 transition-colors text-xs md:text-sm font-semibold min-h-[44px]">
            <span>Genre: Action</span>
            <span class="material-symbols-outlined text-base md:text-[18px]">keyboard_arrow_down</span>
        </button>
        <!-- Status Filter -->
        <button class="flex items-center gap-1.5 md:gap-2 h-10 md:h-10 px-3 md:px-4 rounded-lg bg-surface-dark text-gray-300 border border-white/5 hover:bg-white/10 active:bg-white/10 hover:text-white transition-colors text-xs md:text-sm font-medium min-h-[44px]">
            <span>Status</span>
            <span class="material-symbols-outlined text-base md:text-[18px]">keyboard_arrow_down</span>
        </button>
        <!-- Year Filter -->
        <button class="flex items-center gap-1.5 md:gap-2 h-10 md:h-10 px-3 md:px-4 rounded-lg bg-surface-dark text-gray-300 border border-white/5 hover:bg-white/10 active:bg-white/10 hover:text-white transition-colors text-xs md:text-sm font-medium min-h-[44px]">
            <span>Year</span>
            <span class="material-symbols-outlined text-base md:text-[18px]">keyboard_arrow_down</span>
        </button>
    </div>
    <div class="flex items-center gap-2 self-start md:self-auto">
        <span class="text-xs md:text-sm text-gray-400">Sort by:</span>
        <button class="flex items-center gap-1 text-xs md:text-sm font-bold text-white hover:text-primary active:text-primary transition-colors min-h-[44px] px-2">
            <span>Popularity</span>
            <span class="material-symbols-outlined text-base md:text-[18px]">sort</span>
        </button>
    </div>
  </div>
</template>
