<script setup lang="ts">
</script>

<template>
  <aside class="hidden md:flex w-72 flex-col justify-between border-r border-white/5 bg-background-dark p-6 shrink-0 h-full overflow-y-auto">
    <div class="flex flex-col gap-8">
      <!-- User Profile -->
      <div class="flex items-center gap-3">
        <div class="bg-center bg-no-repeat bg-cover rounded-full h-12 w-12 border-2 border-primary" style='background-image: url("https://lh3.googleusercontent.com/aida-public/AB6AXuB9q7cTy--hT5khMjlapADHl-EFP3Y3qRMoiIGIGnDaTwSRY7pucPYy8-Q2U_JRRD5HYYScTyItu2wtP7XeBdhDzClX18D5LaqFwjrvdM4C10jhwZH93MeXi5M_n24HCSMOllVLd9-s0AXYZChqJDbJJ7yQBonPbJJFEzB7sDnVuOM8eCwbX72gwvAbYahMvA5JSNtH_82PVysj63t8ZenfHtcqskO56ACc4TDGrEvGKupVu2c4ZNYIyUaJZY2TSTRcBAJKvKuG");'></div>
        <div class="flex flex-col">
          <h1 class="text-white text-lg font-bold leading-tight">AnimeStream</h1>
          <p class="text-[#ad92c9] text-sm font-medium">Premium Member</p>
        </div>
      </div>
      
      <!-- Navigation Links -->
      <nav class="flex flex-col gap-2">
        <router-link to="/" class="flex items-center gap-3 px-4 py-3 rounded-lg hover:bg-white/5 transition-colors text-[#ad92c9] hover:text-white group">
          <span class="material-symbols-outlined text-[24px]">home</span>
          <span class="text-sm font-medium">Home</span>
        </router-link>
        <router-link to="/browse" class="flex items-center gap-3 px-4 py-3 rounded-lg hover:bg-white/5 transition-colors text-[#ad92c9] hover:text-white group">
          <span class="material-symbols-outlined text-[24px]">search</span>
          <span class="text-sm font-medium">Search</span>
        </router-link>
        <!-- Active State -->
        <router-link to="/favorites" class="flex items-center gap-3 px-4 py-3 rounded-lg bg-[#362348] text-white shadow-md shadow-primary/10">
          <span class="material-symbols-outlined text-[24px] fill-current">favorite</span>
          <span class="text-sm font-bold">Favorites</span>
        </router-link>
        <router-link to="/settings" class="flex items-center gap-3 px-4 py-3 rounded-lg hover:bg-white/5 transition-colors text-[#ad92c9] hover:text-white group">
          <span class="material-symbols-outlined text-[24px]">settings</span>
          <span class="text-sm font-medium">Settings</span>
        </router-link>
      </nav>
    </div>
    
    <!-- Bottom Actions -->
    <div class="flex flex-col gap-4 mt-auto">
      <div class="bg-gradient-to-br from-primary/20 to-[#261933] p-4 rounded-xl border border-primary/20">
        <p class="text-white text-sm font-bold mb-1">Upgrade Plan</p>
        <p class="text-[#ad92c9] text-xs mb-3">Get access to 4K streaming.</p>
        <button class="w-full py-2 bg-primary hover:bg-[#5e0eb0] text-white text-xs font-bold rounded-lg transition-colors">Go Premium</button>
      </div>
    </div>
  </aside>
</template>
