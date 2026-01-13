<script setup lang="ts">
// Logic for clearing cache would go here
</script>

<template>
  <section aria-label="Storage Management">
    <h3 class="text-white text-lg font-bold leading-tight tracking-[-0.015em] pb-4 px-1">Storage</h3>
    
    <div class="bg-surface-dark rounded-xl border border-white/5 p-5 flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
      <div class="flex items-start gap-4">
        <div class="p-3 bg-[#1a1122] rounded-full border border-white/10 text-primary">
          <span class="material-symbols-outlined text-[24px]">database</span>
        </div>
        <div class="flex flex-col gap-1">
          <span class="text-base font-medium text-white">Cache Size</span>
          <div class="flex items-center gap-2">
            <span class="text-2xl font-bold text-white">450 MB</span>
            <span class="text-xs font-medium px-2 py-0.5 rounded bg-white/5 text-[#ad92c9]">Used</span>
          </div>
          <p class="text-xs text-[#ad92c9] mt-1">Clearing cache may free up space but images will reload.</p>
        </div>
      </div>
      
      <button class="w-full sm:w-auto flex items-center justify-center gap-2 h-10 px-4 bg-transparent hover:bg-red-500/10 border border-white/10 hover:border-red-500/50 text-white hover:text-red-400 text-sm font-bold rounded-lg transition-all whitespace-nowrap group cursor-pointer">
        <span class="material-symbols-outlined text-[18px] group-hover:text-red-400 transition-colors">delete_sweep</span>
        Clear Cache
      </button>
    </div>
  </section>
</template>
