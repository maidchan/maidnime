<script setup lang="ts">
import { ref } from 'vue';
import FavoritesSidebar from '../components/FavoritesSidebar.vue';
import ContinueWatchingCard from '../components/ContinueWatchingCard.vue';
import WatchlistCard from '../components/WatchlistCard.vue';

// Tabs
const activeTab = ref('Continue Watching');
const tabs = ['Continue Watching', 'Watchlist', 'Done'];

// Mock Data
const continueWatching = [
    {
        title: "Jujutsu Kaisen",
        episode: "Episode 14: The Shibuya Incident",
        progress: 75,
        timeLeft: "18m",
        image: "https://lh3.googleusercontent.com/aida-public/AB6AXuAb2s3-xKXEuSEJYtRYEXt360xX0nL5nufY_ukmOPZdKIDlM69cneUC3uTW2CRXw-ggQWN-i8F2Qf9I02GkAaPAVWg0uE75VB5osPebN_JrDGW5O3JbCoeOPXo0-8VajUQOgd1oX8rohoVFNAOLYy4bdxe4YT8xjtTbkyes51nQOmYXNb4otF3E-kUkMTX145HShA0J0PdOOdvmnIGnxrgCS3VB5_R7toEep5zaEdvKjOS8Iv-CNRzemIgQiK1rZ3IO0prX-Tq1",
        season: "S2"
    },
    {
        title: "Cyberpunk: Edgerunners",
        episode: "Episode 6: Girl on Fire",
        progress: 10,
        timeLeft: "22m",
        image: "https://lh3.googleusercontent.com/aida-public/AB6AXuBy-rhFHc7Pf7zhZ_ZVEoQ3YlzIXZfYtocc9SGC237sfh1vggtox7txaHn748uzyNft8ocgQ6JwKGj_w83ZY9tfhyk6Rx2KRMnyWaAxVS62DB-XlQMFU-kEMjDcqH-O5P8EmHTIunO32peQF500j_V52OJY-h7e-biTBM_TNBqsJE3-LZovDme3BsW9RcNAlZ3cLvNS-1hiWjJ3tRTUF4bHQX5XK5mVMIleY6JO4-b1zKOZdGUk3k7JdQYcn0-_12zJBTfqFPRV",
        season: "S1"
    },
    {
        title: "Frieren: Beyond Journey's End",
        episode: "Episode 18: First Class Mage Exam",
        progress: 92,
        timeLeft: "2m",
        image: "https://lh3.googleusercontent.com/aida-public/AB6AXuC8zoNUzToUf163JZG2n1h6cLaTWYk4BvigY2Czm5iQ8M7agQPGT29YbOrQ-HOtyc98bCHyMjrWIvSgmNVpYM-8z07ElvHN5aRSS63UsExLCdwTwx7wX1qdK_gTkOPjVl7qbQN4EKNmTRJ_8qJfiVEwCCWze51vLuFmkpDvoVag93CayIP8zAUn3bOWYvANkzzgcw1zWh7Am4t74AK-9b9DMbJkowLOD3EW8n9u67ZlP0TFMXM92uKwKTZh0uOmRtAwSB5sO_4j",
        season: "S1"
    }
];

const watchlist = [
    {
        title: "One Piece",
        image: "https://lh3.googleusercontent.com/aida-public/AB6AXuCjzB9EcopqJC0yT5EdY-_lan4gZa94Lmnx-fiXu2vivVu9fjJnruuch7RJNbtppOoqY_PbKZY3krB34pI93BjEn9gM7rZsYmhvwizTSt4ys5hC0TEZO94B_HgeDbV4hEZhw7wqH1ooPr33J-6kJgTf350x7ovOOVC8mCxvWWtISsgXzjAqpp1-3lp0r0h4p-OL8fm79xJszo3mhxWzzWw3Q5S-mxTqkQJdIiLJpX5oDRilA8JC4d5z9ikuOg2nj9Ew4QJJ0n0U",
        genre: "Adventure, Action",
        badge: "SUB"
    },
    {
        title: "Chainsaw Man",
        image: "https://lh3.googleusercontent.com/aida-public/AB6AXuCa5P7GvC-WwrEq457kBWjpE_774e4BOmRGscHxpT_ad0dMZ84ToTeStl9fqC1_psV17bsxAN5fPuGMGvKtidD3BWtql5bdtxYuOFrM_vshmn9IapxkChfaC2NCfTHXRBbqYv-kYiaq_mlN9YN632J4OfN9CEk3QIY5e0XwWNkT2zRM6SobNYaehtv4UXXD6LgdCo8VmhBLNzZsWcD0yUQSMwYscSFJUu-l9PlU6EcA0nBYQJ5EEV71Lv_55aNjvtpqxvWtAtxs",
        genre: "Dark Fantasy, Action",
        badge: "DUB"
    },
    {
        title: "Spy x Family",
        image: "https://lh3.googleusercontent.com/aida-public/AB6AXuCupzInRwpmHB9pdEMlk1vDyscPIpl68MVXv0-y6wcqWTLJf7cr0hy4EfIyrJ2v3NG5HEUcHC83bl10Jzs8Y41EjT5JO9onBbyKDZNXn2UFks2FhNXJ0soMQXJybd-pnXUSJotPTFYiKAwxfRJmtxqaFKJ0ARamu4Htw4l8EydSjradyjLZ0-H5B5izZbqC1bzUinayPxRZuDhrzW3YHgVvCQkehA-dAZu2ftESd03YMupWw2y61cahCBi7iS7vi2YOXyqvi_LM",
        genre: "Comedy, Slice of Life"
    },
    {
        title: "Attack on Titan",
        image: "https://lh3.googleusercontent.com/aida-public/AB6AXuC5gJIyjt77_SLddq4mDWfDeOhXz8_r_dxlKpmcTLGKBnox2jnFt4w1iX7XT62rybvtP-y7bovN0eaPUYQDh1gIQBKhYR2_mA7WBlk2zCwPOxxTTfEFfj70FaXM89jkwNhLFcFVHx2qk2xHEwO1AD39ZYIIntxr0gZ2A7emn27kUuvc4RjQXcKFV2soAOUZ4y5pK05h0TyKJhsXNfPrR90Z-Vnb7iSINdiisG41epvRDcBzrXoXcVkBbJBIcYEFWfrIeUJcNSsA",
        genre: "Action, Drama",
        badgeColor: "primary"
    },
    {
        title: "Blue Lock",
        image: "https://lh3.googleusercontent.com/aida-public/AB6AXuAGyB_8jrLWUaBVtHml_YlCNYZyNhQboIkeLupOwsF8DZUmZi4r_9mIeBW_oHCeJE0cSHQ7J00OKhCmIMYv64gkEd_Iomy8hjymPrOOblXdVMW-9V25kATPp21qSb4LcR2y32w4S0D1j7Vv_l6dyEepKOGKGCO8cnvoqizQINzmGkQPylmtRd_IlI1ttTu19djMTJsTEHTOQiXvS679O98pIPBpiA4RdbB4EpAv3n2DY-DeoxZxp0uUMqFEFTSBddGnUwip1BtC",
        genre: "Sports, Thriller"
    }
];
</script>

<template>
  <div class="bg-background-light dark:bg-background-dark text-slate-900 dark:text-white h-screen overflow-hidden font-display flex flex-row">
    <!-- Side Layout -->
    <FavoritesSidebar />

    <!-- Main Content -->
    <main class="flex-1 flex flex-col h-full overflow-hidden bg-background-dark relative">
      <!-- Top Gradient Overlay -->
      <div class="absolute top-0 left-0 w-full h-64 bg-gradient-to from-primary/5 to-transparent pointer-events-none"></div>
      
      <div class="flex-1 overflow-y-auto p-4 md:p-8 lg:px-12 scroll-smooth">
        <div class="max-w-6xl mx-auto flex flex-col gap-6">
          <!-- Page Header -->
          <header class="flex flex-wrap items-center justify-between gap-4 py-4">
            <h2 class="text-white text-3xl md:text-4xl font-black tracking-tight">My Favorites</h2>
            <button class="flex items-center justify-center h-10 px-6 rounded-lg bg-surface-dark border border-white/10 hover:bg-white/5 text-white text-sm font-bold tracking-wide transition-all hover:scale-105 active:scale-95">
              Edit List
            </button>
          </header>

          <!-- Tabs -->
          <div class="sticky top-0 z-10 bg-background-dark/95 backdrop-blur-sm pt-2">
            <div class="border-b border-[#4d3267]">
              <div class="flex gap-8">
                <button 
                  v-for="tab in tabs" 
                  :key="tab"
                  @click="activeTab = tab"
                  class="flex items-center justify-center border-b-[3px] pb-3 px-2 transition-colors"
                  :class="activeTab === tab ? 'border-b-primary text-white' : 'border-b-transparent text-[#ad92c9] hover:text-white'"
                >
                  <p class="text-sm font-bold tracking-wide">{{ tab }}</p>
                </button>
              </div>
            </div>
          </div>

          <!-- Content: Continue Watching -->
          <section class="flex flex-col gap-4 mt-2">
            <div class="flex items-center gap-2 mb-2">
              <span class="material-symbols-outlined text-primary">history</span>
              <h3 class="text-white text-lg font-bold">Pick up where you left off</h3>
            </div>
            <!-- Horizontal Card List -->
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-4">
                <ContinueWatchingCard 
                    v-for="item in continueWatching" 
                    :key="item.title" 
                    v-bind="item" 
                />
            </div>
          </section>

          <hr class="border-white/5 my-2"/>

          <!-- Content: Watchlist Grid -->
          <section class="flex flex-col gap-4">
            <div class="flex items-center justify-between mb-2">
              <div class="flex items-center gap-2">
                <span class="material-symbols-outlined text-[#ad92c9]">grid_view</span>
                <h3 class="text-white text-lg font-bold">Watchlist</h3>
              </div>
              <div class="flex gap-2">
                <button class="p-1 hover:bg-white/5 rounded text-[#ad92c9] hover:text-white transition-colors"><span class="material-symbols-outlined">filter_list</span></button>
                <button class="p-1 hover:bg-white/5 rounded text-[#ad92c9] hover:text-white transition-colors"><span class="material-symbols-outlined">sort</span></button>
              </div>
            </div>
            <!-- Responsive Grid -->
            <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-x-4 gap-y-8 pb-10">
                <WatchlistCard 
                    v-for="item in watchlist" 
                    :key="item.title"
                    v-bind="item"
                />
            </div>
          </section>

          <!-- Spacer -->
          <div class="h-10"></div>
        </div>
      </div>
    </main>

    <!-- Mobile Nav Overlay -->
    <div class="md:hidden fixed bottom-0 left-0 w-full bg-surface-dark border-t border-white/5 p-2 flex justify-around z-50">
      <router-link to="/" class="flex flex-col items-center p-2 text-[#ad92c9]">
        <span class="material-symbols-outlined">home</span>
        <span class="text-[10px] mt-1">Home</span>
      </router-link>
      <router-link to="/browse" class="flex flex-col items-center p-2 text-[#ad92c9]">
        <span class="material-symbols-outlined">search</span>
        <span class="text-[10px] mt-1">Search</span>
      </router-link>
      <router-link to="/favorites" class="flex flex-col items-center p-2 text-primary">
        <span class="material-symbols-outlined fill-current">favorite</span>
        <span class="text-[10px] mt-1 font-bold">Favorites</span>
      </router-link>
      <a href="#" class="flex flex-col items-center p-2 text-[#ad92c9]">
        <span class="material-symbols-outlined">settings</span>
        <span class="text-[10px] mt-1">Settings</span>
      </a>
    </div>
  </div>
</template>
