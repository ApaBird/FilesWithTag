<template>
  <div class="grid grid-cols-6 gap-4">
    <div
      v-if="pathAddress != 'C:/'"
      class="p-4 h-[40px] rounded-[8px] text-white flex justify-start items-center bg-[#1f1f1f] hover:cursor-pointer hover:bg-[#444746]"
      @click="storePathAddress.setPath(pathAddress.split('/').slice(0, -1).join('/'))"
    >
      <div class="flex w-[30px]">
        <Icon name="simple-line-icons:action-undo" style="color: white" />
      </div>
      <p>Назад</p>
    </div>
    <div
      v-for="folder in activeFolderContent"
      class="h-[40px] rounded-[8px] text-white flex justify-start items-center bg-[#1f1f1f] hover:cursor-pointer hover:bg-[#444746] p-4"
      @click="storePathAddress.setPath(folder.dir)"
    >
      <div class="flex w-[30px] ">
        <Icon name="simple-line-icons:folder-alt" style="color: white" />
      </div>
      <p class="truncate" :title="folder.name">{{ folder.name }}</p>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from "vue";

import { addressStore } from '~/store/pathAddress';

const storePathAddress = addressStore()

const pathAddress = computed(() => {
  return storePathAddress.path
})


const folders = ref({});

async function getFolders() {
  folders.value = await $fetch("http://127.0.0.1:8050/OsTree").then((t) => t);
}


const activeFolderContent = computed(() => {
  if (folders.value.content)
    return getFolderContent([folders.value], pathAddress.value);
});

function getFolderContent(start, path) {
  console.log(path)
  for (let i = 0; i < start.length; i++) {
    if (start[i].dir === path) {
      return start[i].content;
    } else if (start[i].content.length > 0) {
      let result = getFolderContent(start[i].content, path);
      if (result) return result;
    }
  }
}

onMounted(async () => {
  await getFolders();
});
</script>
