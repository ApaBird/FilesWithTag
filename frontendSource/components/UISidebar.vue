<template>
  <div class="border-l border-[#444746] w-[500px] flex justify-start items-center p-4 text-white flex-col gap-3">
    <p>{{ activePicture.Name ?? 'No name' }}</p>
    <div class="max-h-[calc(100vh-250px)]">
      <img  class="w-full h-full object-cover rounded-xl" :src="`data:image/jpg;base64, ${activePicture.Content}`"/>
    </div>
    <div class="flex flex-col gap-4 w-full">
      <div class="grid grid-cols-3 gap-3">
        <div 
          v-for="tag in activePicture.Tags"
          @click="filters.addFilter(tag)"
          class="bg-[#1f1f1f] rounded-[15px] h-[32px] hover:cursor-pointer hover:bg-[#444746]">
          <p class="truncate flex justify-center items-center p-1" :title="tag">{{ tag }}</p>
        </div>
      </div>
      <div class="flex flex-row w-full gap-3">
        <input
        class="bg-[#1f1f1f] rounded-[8px] px-2 py-1 text-white w-full"
        placeholder="Tagname"
        v-model="newTag"
      />
      <div @click="addTag"
        class="rounded-[8px] w-[32px] flex h-[32px] bg-[#1f1f1f] justify-center items-center hover:cursor-pointer hover:bg-[#444746]">
        <Icon name="carbon:add" style="color: white" size="24px" />
      </div>
    </div>
    </div>
  </div>
</template>



<script setup>
const newTag = ref('')

import { itemStore } from '~/store/item';
import { filtersStore } from '~/store/filters';

const filters = filtersStore()

const activePicture = computed(() => {
  return itemStore().item
})




function addTag() {
  if (newTag.value) {
    activePicture.value.Tags.push(newTag.value);
    newTag.value = '';
  }
}

</script>

<style scoped>
::placeholder {
  color: rgba(255, 255, 255, 0.2);
}
</style>
