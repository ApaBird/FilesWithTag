<template>
  <div class="flex bg-[#131314] w-full">
    <div class="h-screen p-4 flex flex-col gap-4 w-full">
      <input
        class="bg-[#1f1f1f] rounded-[8px] px-2 py-1 text-white"
        :placeholder="pathAddress"
        disabled
      />
      <slot></slot>
    </div>
    <UISidebar v-show="true"></UISidebar>
  </div>
</template>

<script setup>
import { addressStore } from "~/store/pathAddress";

const pathAddress = computed(() => addressStore().path);
</script>
