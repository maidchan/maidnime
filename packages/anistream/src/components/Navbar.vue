<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const isScrolled = ref(false);
const searchQuery = ref('');
const isUserMenuOpen = ref(false);
const isNotificationsOpen = ref(false);

const handleScroll = () => {
    isScrolled.value = window.scrollY > 20;
};

const handleSearch = () => {
    if (searchQuery.value.trim()) {
        router.push({ path: '/browse', query: { q: searchQuery.value } });
        searchQuery.value = '';
    }
};

onMounted(() => {
    window.addEventListener('scroll', handleScroll);
});

onUnmounted(() => {
    window.removeEventListener('scroll', handleScroll);
});
</script>

<template>
    <nav 
        class="fixed top-0 z-50 w-full transition-all duration-300 border-b border-transparent"
        :class="[isScrolled ? 'glass-nav py-2' : 'bg-transparent py-4']"
    >
        <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
            <div class="flex h-12 items-center justify-between gap-4">
                <!-- Logo -->
                <router-link to="/" class="flex items-center gap-2 cursor-pointer group">
                    <div class="flex size-8 items-center justify-center rounded-lg bg-primary text-white group-hover:scale-110 transition-transform">
                        <span class="material-symbols-outlined">play_arrow</span>
                    </div>
                    <span class="text-xl font-bold tracking-tight text-white">Anistream</span>
                </router-link>

                <!-- Navigation Links -->
                <nav class="hidden md:flex items-center gap-6 lg:gap-9">
                    <router-link to="/" class="text-white hover:text-primary transition-colors text-sm font-medium leading-normal" active-class="text-primary font-bold">Home</router-link>
                    <router-link to="/browse" class="text-white hover:text-primary transition-colors text-sm font-medium leading-normal" active-class="text-primary font-bold">Browse</router-link>
                    <router-link to="/mylist" class="text-white hover:text-primary transition-colors text-sm font-medium leading-normal" active-class="text-primary font-bold">My List</router-link>
                    <router-link to="/schedule" class="text-white hover:text-primary transition-colors text-sm font-medium leading-normal" active-class="text-primary font-bold">Schedule</router-link>
                </nav>

                <!-- Search Bar (Desktop) -->
                <div class="hidden md:flex flex-1 max-w-md items-center">
                    <div class="relative w-full group">
                        <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                            <span class="material-symbols-outlined text-gray-400">search</span>
                        </div>
                        <input 
                            v-model="searchQuery"
                            type="text" 
                            class="block w-full rounded-full border-0 bg-surface-dark py-2 pl-10 pr-3 text-sm text-white placeholder-gray-400 focus:ring-2 focus:ring-primary focus:bg-surface-dark-hover transition-colors" 
                            placeholder="Find anime..." 
                            @keyup.enter="handleSearch"
                        />
                    </div>
                </div>

                <!-- Right Actions -->
                <div class="flex items-center gap-4">
                    <!-- Notifications -->
                    <div class="relative">
                        <button 
                            @click="isNotificationsOpen = !isNotificationsOpen"
                            class="relative rounded-full p-2 text-gray-300 hover:bg-white/10 hover:text-white transition-colors"
                        >
                            <span class="material-symbols-outlined">notifications</span>
                            <span class="absolute top-2 right-2 size-2 rounded-full bg-red-500 border border-background-dark"></span>
                        </button>
                        
                        <div v-if="isNotificationsOpen" class="absolute right-0 mt-2 w-80 bg-surface-dark border border-white/10 rounded-xl shadow-2xl overflow-hidden z-50">
                            <div class="p-4 border-b border-white/5 flex justify-between items-center">
                                <h4 class="font-bold text-white text-sm">Notifications</h4>
                                <span class="text-[10px] text-primary uppercase font-bold tracking-wider">2 New</span>
                            </div>
                            <div class="max-h-96 overflow-y-auto">
                                <div class="p-4 hover:bg-white/5 cursor-pointer border-b border-white/5">
                                    <p class="text-sm text-white font-medium">New Episode Available!</p>
                                    <p class="text-xs text-gray-400 mt-1">Solo Leveling Episode 12 is now streaming.</p>
                                </div>
                                <div class="p-4 hover:bg-white/5 cursor-pointer border-b border-white/5">
                                    <p class="text-sm text-white font-medium">Added to Favorites</p>
                                    <p class="text-xs text-gray-400 mt-1">Gintama has been successfully added to your list.</p>
                                </div>
                            </div>
                            <div class="p-3 text-center bg-white/5">
                                <button class="text-xs text-gray-400 hover:text-white transition-colors">Mark all as read</button>
                            </div>
                        </div>
                    </div>

                    <!-- User Profile -->
                    <div class="relative">
                        <div 
                            class="h-9 w-9 cursor-pointer rounded-full bg-cover bg-center border border-white/10 hover:border-primary transition-all hover:scale-105"
                            style="background-image: url('https://lh3.googleusercontent.com/aida-public/AB6AXuC8ml6mGUK1rOI6-uTUxPIxntQP9uuQaFSbU9tefb3dwrQ-t-h7cG30OB3YHRVDZQigM0Za24mj_baka0MadNXCJzwBHcXHuV3KSUk1U6vY4w5KDU39LKTBE43mpWKqMX3nx01TLoKnufUoAOeEOWPeHvaE0ahVpOVOZAlf-H8Td1Tjtd19nnxqzlbDvW__2TuPYVDxzxG0ogWyFGHZ-HgE_MGrZtiOFbhXp9yHeY9zfSunQJUEvH7_tzLJorpI4mka0Mtah25Q');"
                            aria-label="User profile avatar"
                            @click="isUserMenuOpen = !isUserMenuOpen"
                        ></div>
                        
                        <div v-if="isUserMenuOpen" class="absolute right-0 mt-2 w-48 bg-surface-dark border border-white/10 rounded-xl shadow-2xl overflow-hidden z-50 py-2">
                            <router-link to="/settings" class="flex items-center gap-3 px-4 py-2 text-sm text-gray-300 hover:bg-white/5 hover:text-white transition-colors" @click="isUserMenuOpen = false">
                                <span class="material-symbols-outlined text-lg">person</span>
                                Profile
                            </router-link>
                            <router-link to="/mylist" class="flex items-center gap-3 px-4 py-2 text-sm text-gray-300 hover:bg-white/5 hover:text-white transition-colors" @click="isUserMenuOpen = false">
                                <span class="material-symbols-outlined text-lg">bookmark</span>
                                Watchlist
                            </router-link>
                            <router-link to="/settings" class="flex items-center gap-3 px-4 py-2 text-sm text-gray-300 hover:bg-white/5 hover:text-white transition-colors" @click="isUserMenuOpen = false">
                                <span class="material-symbols-outlined text-lg">settings</span>
                                Settings
                            </router-link>
                            <div class="h-px bg-white/5 my-2"></div>
                            <button class="w-full flex items-center gap-3 px-4 py-2 text-sm text-red-400 hover:bg-red-400/10 transition-colors">
                                <span class="material-symbols-outlined text-lg">logout</span>
                                Sign Out
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        
        <!-- Click Outside Backdrop for Dropdowns -->
        <div v-if="isUserMenuOpen || isNotificationsOpen" class="fixed inset-0 z-[-1]" @click="isUserMenuOpen = false; isNotificationsOpen = false"></div>
    </nav>
</template>
