<template>
  <div class="border-l border-[#444746] w-[500px] flex justify-start items-center p-4 text-white flex-col gap-3">
    <p>Старик</p>
    <div class="w-[300px]">
      <img src="C:\Users\kuram\Downloads\Telegram Desktop\image_2024-07-22_16-54-14 (2).png"></img>
    </div>
    <div class="flex flex-col gap-5 justify-center items-center">
      <div class="grid grid-cols-3 gap-3">
        <div v-for="m in tags"
          class="bg-[#1f1f1f] rounded-[15px] w-[106px] h-[30px] flex  justify-center items-center p-1">
          <p class="truncate" :title="m">{{ m }}</p>
        </div>
      </div>
      <div class="flex flex-row w-full gap-3">
        <input
        class="bg-[#1f1f1f] rounded-[8px] px-2 py-1 text-white w-full"
        :placeholder="pathAddress"
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

const tags = ref(['Старик', 'Поел говна', 'Пенсия'])
const newTag = ref('')


function addTag() {
  if (newTag.value) {
    tags.value.push(newTag.value);
    newTag.value = '';
  }
}

</script>
