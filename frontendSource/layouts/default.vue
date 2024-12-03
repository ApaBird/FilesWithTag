<template>
  <div class="flex bg-[#131314] w-full">
    <div class="h-screen p-4 flex flex-col gap-4 w-full">
      <UIPath></UIPath>
      <!-- <input
        class="bg-[#1f1f1f] rounded-[8px] px-2 py-1 text-white"
        :placeholder="pathAddress"
        v-model.lazy="addressStore().path"
      /> -->
      <slot></slot>
    </div>
    <UISidebar v-if="activePicture" />
  </div>
</template>

<script setup>
import { addressStore } from "~/store/pathAddress";
import { itemStore } from "~/store/item";

const pathAddress = computed(() => addressStore().path);

const activePicture = computed(() => itemStore().item.Content);
</script>

<style scoped></style>
