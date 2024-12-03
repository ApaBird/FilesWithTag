<template>
  <div class="h-screen flex flex-col gap-4 overflow-hidden">
    <div class="max-h-[220px] min-h-[100px] overflow-y-auto scrollbar">
      <div class="grid grid-cols-6 gap-4 mr-2">
        <div
          v-if="pathAddress != 'C:/'"
          class="p-4 h-[40px] rounded-[8px] text-white flex justify-start items-center bg-[#1f1f1f] hover:cursor-pointer hover:bg-[#444746]"
          @click="storePathAddress.back()"
        >
          <div class="flex w-[30px]">
            <Icon name="simple-line-icons:action-undo" style="color: white" />
          </div>
          <p>Назад</p>
        </div>
        <div
          v-for="folder in folders"
          class="h-[40px] rounded-[8px] min-h-[40px] text-white flex justify-start items-center bg-[#1f1f1f] hover:cursor-pointer hover:bg-[#444746] p-4"
          @click="onClickFolder(folder)"
        >
          <div class="flex w-[30px]">
            <Icon name="simple-line-icons:folder-alt" style="color: white" />
          </div>
          <p class="truncate" :title="folder.name">{{ folder.Name}}</p>
        </div>
      </div>
    </div>
      <Pictures/>
  </div>
</template>

<script setup>
import { addressStore } from "~/store/pathAddress";
import Pictures from "./pictures.vue";

const storePathAddress = addressStore();

const pathAddress = computed(() => {
  return storePathAddress.path;
});

const folders = ref({});

async function getFolders() {
  folders.value = []
  folders.value = await $fetch(`http://127.0.0.1:8050/Dirs?Path=${pathAddress.value}`).then((t) => t.dirs);
}


function getFolderContent(start, path) {
  for (let i = 0; i < start.length; i++) {
    if (start[i].dir === path) {
      return start[i].content;
    } else if (start[i].content.length > 0) {
      let result = getFolderContent(start[i].content, path);
      if (result) return result;
    }
  }
}

function onClickFolder(folder) {
  storePathAddress.setPath(folder.Path)
}

watch(() => pathAddress.value,() => {
  getFolders();
}, {immediate:true})

</script>

