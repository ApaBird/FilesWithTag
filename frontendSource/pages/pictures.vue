<template>
  <div class="flex flex-col gap-3 max-h-[calc(100vh-60px)]">
    <UIFilters />
    <div
      class="mx-auto flex gap-4 flex-wrap overflow-y-auto scrollbar h-[calc(100%-100px)]"
    >
      <div
        class="flex max-h-[360px] xl:max-w-[400px] sm:max-w-[200px] flex flex-col gap-1"
        v-for="item in itemsFolder"
      >
        <div class="xl:max-w-[400px] max-h-[360px] min-h-[330px] sm:max-w-[240px] ">
          <img
            class="w-full h-full object-cover rounded-xl"
            :src="`data:image/jpg;base64, ${item.Content}`"
            @click="
              onClickPicture({
                Content: item.Content,
                Name: item.Name,
                Tags: item.Tags,
              })
            "
          />
        </div>
        <div class="flex justify-center items-center w-full">
          <p class="h-[30px] text-white truncate">{{ item.Name }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { itemStore } from "~/store/item";

const props = defineProps({
  itemsFolder: Array,
});

const activePicture = itemStore();

function onClickPicture(picture) {
  activePicture.setItem(picture);
}
</script>
