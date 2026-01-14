<script setup lang="ts">
import Navbar from './components/Navbar.vue';
import Footer from './components/Footer.vue';
</script>

<template>
  <div class="bg-background-light dark:bg-background-dark min-h-screen text-gray-900 dark:text-white font-display">
    <Navbar v-if="!$route.meta.hideNavbar" />
    
    <router-view />
    
    <Footer v-if="!$route.meta.hideNavbar" />
  </div>
</template>
