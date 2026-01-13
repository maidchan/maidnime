<script setup lang="ts">
</script>

<template>
  <section aria-label="About Application" class="pt-4 border-t border-white/5 mt-2">
    <h3 class="text-white text-lg font-bold leading-tight tracking-[-0.015em] pb-4 px-1">About</h3>
    
    <div class="flex flex-col gap-3">
      <div class="flex items-center justify-between py-2 px-1">
        <span class="text-[#ad92c9] text-sm">Version</span>
        <span class="text-white font-mono text-sm bg-surface-dark px-2 py-1 rounded border border-white/5">v2.4.0</span>
      </div>
      
      <a href="#" class="flex items-center justify-between py-3 px-4 rounded-lg bg-surface-dark hover:bg-white/5 border border-transparent hover:border-white/10 transition-all group">
        <span class="text-white text-sm font-medium">Privacy Policy</span>
        <span class="material-symbols-outlined text-[#ad92c9] text-[18px] group-hover:translate-x-1 transition-transform">chevron_right</span>
      </a>
      
      <a href="#" class="flex items-center justify-between py-3 px-4 rounded-lg bg-surface-dark hover:bg-white/5 border border-transparent hover:border-white/10 transition-all group">
        <span class="text-white text-sm font-medium">Terms of Service</span>
        <span class="material-symbols-outlined text-[#ad92c9] text-[18px] group-hover:translate-x-1 transition-transform">chevron_right</span>
      </a>
    </div>
    
    <div class="mt-8 flex justify-center">
      <button class="text-[#ad92c9] hover:text-red-400 text-sm font-medium transition-colors py-2 px-4 rounded hover:bg-red-500/5 cursor-pointer">
        Log Out
      </button>
    </div>
    
    <div class="mt-8 text-center pb-8">
      <p class="text-xs text-[#ad92c9]/50">© 2024 AnimeStream Inc. All rights reserved.</p>
    </div>
  </section>
</template>
