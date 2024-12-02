<template>
  <div class="flex w-full gap-3 max-h-[30px]">
    <div v-for="item in filters">
      <div
        class="flex bg-[#1f1f1f] rounded-[8px] h-[32px] gap-3 p-2 text-white justify-between items-center"
      >
        <div class="truncate" :title="item">
          {{ item }}
          {{ item.length }}
        </div>
        <div class=" min-w-[18px] max-w-[18px] h-[18px] flex items-center justify-center">
          <Icon name="carbon:close-large" class="text-white text-[18px] hover:cursor-pointer hover:bg-rose-800" @click="removeFilter(item)"/>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { filtersStore } from "~/store/filters";

const filters = computed(() => {
  return filtersStore().filters;
});

function removeFilter(item) {
  filtersStore().removeFilter(item)
}
</script>
