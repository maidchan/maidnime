<script setup lang="ts">
// In a real app, these values would come from a store or prop
</script>

<template>
  <section aria-label="Profile Settings" class="relative overflow-hidden rounded-xl bg-surface-dark border border-white/5 p-6 shadow-lg">
    <!-- Subtle background glow -->
    <div class="absolute -top-24 -right-24 w-64 h-64 bg-primary/20 rounded-full blur-3xl pointer-events-none"></div>
    
    <div class="relative z-10 flex flex-col md:flex-row items-center gap-6 justify-between">
      <div class="flex items-center gap-5 w-full md:w-auto">
        <div class="relative group cursor-pointer">
          <div class="bg-center bg-no-repeat aspect-square bg-cover rounded-full h-20 w-20 ring-4 ring-[#362348] group-hover:ring-primary/50 transition-all" style='background-image: url("https://lh3.googleusercontent.com/aida-public/AB6AXuAuv1XWdzT_WHxWJUburH5hZTm0SzsjKOaZLAZMZKgn0gLxd9JSXCKfn_FAH8N7aHsGErn8PtbykZahrVW5tsPOoGSpMD-8QralqCJzAoWRnSI5hTC1KR2KkTf4yOi5Zwj6pv4DTsH00LL8IGtFpa913_1ohFkHRa6sS8Of2jylVIjxJH1RBq-Mcl0Ikgjf8-no2YDld4tcVmwN8fyM8UVTo4Yi5DOY4MPqdYbliY408rN_FojWPrFGcgBrxzGP8wEnHL8IFWmz");'></div>
          <div class="absolute bottom-0 right-0 bg-primary rounded-full p-1 border-2 border-surface-dark flex items-center justify-center">
            <span class="material-symbols-outlined text-white text-[14px] font-bold">edit</span>
          </div>
        </div>
        
        <div class="flex flex-col">
          <h3 class="text-xl text-white font-bold leading-tight">OtakuMaster99</h3>
          <p class="text-[#ad92c9] text-sm">otakumaster@example.com</p>
          <div class="mt-2 inline-flex items-center gap-1 px-2 py-0.5 rounded bg-primary/10 border border-primary/20 w-fit">
            <span class="material-symbols-outlined text-primary text-[12px]">crown</span>
            <span class="text-primary text-xs font-bold uppercase tracking-wide">Premium</span>
          </div>
        </div>
      </div>
      
      <button class="w-full md:w-auto flex items-center justify-center gap-2 h-10 px-6 bg-primary hover:bg-[#6b10c6] text-white text-sm font-bold rounded-lg transition-all shadow-md hover:shadow-primary/25">
        <span>Edit Profile</span>
      </button>
    </div>
  </section>
</template>
