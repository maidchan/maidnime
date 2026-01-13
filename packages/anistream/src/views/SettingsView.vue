<script setup lang="ts">
import ProfileSection from '../components/settings/ProfileSection.vue';
import PreferencesSection from '../components/settings/PreferencesSection.vue';
import StorageSection from '../components/settings/StorageSection.vue';
import AboutSection from '../components/settings/AboutSection.vue';
</script>

<template>
  <div class="bg-background-light dark:bg-background-dark min-h-screen text-slate-900 dark:text-white font-display flex flex-col antialiased selection:bg-primary selection:text-white overflow-x-hidden">
    <div class="flex h-full grow flex-col">
      <!-- Main Content Wrapper -->
      <div class="px-4 md:px-8 lg:px-40 flex flex-1 justify-center py-8">
        <div class="flex flex-col w-full max-w-[800px] flex-1 gap-8">
          
          <!-- Header -->
          <header class="flex flex-col gap-2">
            <router-link to="/" class="inline-flex items-center gap-2 text-[#ad92c9] hover:text-white transition-colors w-fit group">
              <span class="material-symbols-outlined text-[20px] group-hover:-translate-x-1 transition-transform">arrow_back</span>
              <span class="text-sm font-medium">Back</span>
            </router-link>
            <h1 class="text-3xl md:text-4xl font-black leading-tight tracking-[-0.033em]">Settings</h1>
          </header>

          <!-- Sections -->
          <ProfileSection />
          <PreferencesSection />
          <StorageSection />
          <AboutSection />
          
        </div>
      </div>
    </div>
  </div>
</template>
