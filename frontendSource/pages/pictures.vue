<template>
  <div class="flex flex-col gap-3 max-h-[calc(100vh-130px)]">
    <UIFilters />
    <div
       ref="itemsRef" class="mx-auto flex gap-4 flex-wrap overflow-y-auto scrollbar max-h-[calc(100%-100px)]"
    >
      <div
        class="flex max-h-[360px] xl:max-w-[400px] sm:max-w-[200px] flex flex-col gap-1"
        v-for="item in itemsFolder"
      >
        <div class="xl:max-w-[400px] min-w-[300px] max-h-[360px] min-h-[330px] sm:max-w-[240px] border border-[#444746] rounded-xl ">
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
import { addressStore } from "~/store/pathAddress";
const storePathAddress = addressStore();

const pathAddress = computed(() => {
  return storePathAddress.path;
});

const activePicture = itemStore();

function onClickPicture(picture) {
  activePicture.setItem(picture);
}

const itemsRef = ref(null)
const lastLength = ref(0)

onMounted(() => {
  itemsRef.value.addEventListener('scroll',() => {
    if (Math.floor(itemsRef.value.scrollHeight - itemsRef.value.scrollTop - itemsRef.value.clientHeight) <= 1) {
      if(itemsFolder.value.length == lastLength.value) return
      lastLength.value = itemsFolder.value.length
      $fetch(`http://127.0.0.1:8050/Files?Count=20&Offset=${itemsFolder.value.length}&Path=${pathAddress.value}&Ftype=Image`).then((t) => {
        if(!t.Files) return
        itemsFolder.value = [...itemsFolder.value, ...t.Files];
      });
    }
  })
})


const itemsFolder = ref([]);


watch(
  () => pathAddress.value,
  async (value) => {
    itemsFolder.value = []
    $fetch(`http://127.0.0.1:8050/Files?Count=20&Offset=0&Path=${value}&Ftype=Image`).then((t) => {
      itemsFolder.value = t.Files;
    });
  }, {
    immediate: true
  }
);
</script>
