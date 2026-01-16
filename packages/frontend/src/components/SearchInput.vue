<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{
    modelValue: string;
}>();

const emit = defineEmits<{
    (e: 'update:modelValue', value: string): void;
    (e: 'search', value: string): void;
}>();

const localQuery = ref(props.modelValue);

watch(() => props.modelValue, (newVal) => {
    localQuery.value = newVal;
});

watch(localQuery, (newVal) => {
    emit('update:modelValue', newVal);
});

const handleEnter = () => {
    emit('search', localQuery.value);
};
</script>

<template>
  <div class="flex flex-col md:flex-row gap-3 md:gap-4 items-stretch md:items-center w-full">
    <!-- Back Button -->
    <router-link 
        to="/"
        class="flex shrink-0 w-full md:w-auto h-12 cursor-pointer items-center justify-center rounded-xl bg-surface-dark hover:bg-[#4a3061] active:bg-[#4a3061] transition-colors text-white gap-2 px-5 md:px-6 text-sm font-bold shadow-sm min-h-[44px]"
    >
      <span class="material-symbols-outlined text-lg md:text-[20px]">arrow_back</span>
      <span>Back</span>
    </router-link>

    <!-- Search Input -->
    <div class="flex-1 relative group">
      <div class="absolute inset-y-0 left-0 pl-3 md:pl-4 flex items-center pointer-events-none">
        <span class="material-symbols-outlined text-[#ad92c9] text-lg md:text-xl">search</span>
      </div>
      <input 
        v-model="localQuery"
        @keyup.enter="handleEnter"
        class="block w-full h-12 rounded-xl border-none bg-surface-dark pl-11 md:pl-12 pr-20 md:pr-24 text-white placeholder:text-[#ad92c9] focus:ring-2 focus:ring-primary focus:bg-[#422a57] transition-all text-sm md:text-base min-h-[44px]" 
        placeholder="Search for anime, characters, or genre..." 
        type="text" 
      />
      <div class="absolute inset-y-0 right-0 pr-1 md:pr-2 flex items-center gap-0.5 md:gap-1">
        <button class="p-2 text-[#ad92c9] hover:text-white rounded-full hover:bg-white/10 active:bg-white/10 transition-colors flex items-center justify-center min-h-[44px] min-w-[44px]" @click="handleEnter">
          <span class="material-symbols-outlined text-lg md:text-[20px]">search</span>
        </button>
        <button 
            v-if="localQuery"
            @click="localQuery = ''"
            class="p-2 text-[#ad92c9] hover:text-white rounded-full hover:bg-white/10 active:bg-white/10 transition-colors flex items-center justify-center min-h-[44px] min-w-[44px]"
        >
          <span class="material-symbols-outlined text-lg md:text-[20px]">close</span>
        </button>
      </div>
    </div>
  </div>
</template>
