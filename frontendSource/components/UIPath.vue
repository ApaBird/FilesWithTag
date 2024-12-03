<template>
  <!-- {{pathAdressParts}} -->
  <div class="flex w-full h-[30px] bg-[#1f1f1f] rounded-[8px] p-2 items-center text-white" @click="activeInput = true, currentInput = pathAddress + '/'">
    <!-- <p class="hover:bg-[#444746] rounded hover:cursor-pointer" @click.stop="addressStore().setPath('C:/')">{{pathAdressParts.slice(1)}}</p> -->
    <div v-for="item, index in pathAdressParts" class="flex" v-if="activeInput == false">
      <p class="hover:bg-[#444746] rounded hover:cursor-pointer truncate" @click.stop="onClickPath(index)">{{item}}</p>
      <p>/</p>
    </div>
    <input class="flex items-center w-full h-[30px] mb-[1px] bg-[#1f1f1f] rounded-[8px] outline-none"
     v-if="activeInput == true" v-model="currentInput" @keydown.enter="activeInput = false, addressStore().setPath($event.target.value)" @keydown.esc="activeInput = false">
  </div>
</template>

<script setup>
import { addressStore } from "~/store/pathAddress";

const pathAddress = computed(() => addressStore().path);

const activeInput = ref(false)

const currentInput = ref('')

const pathAdressParts = computed(() => {
  let path = pathAddress.value.split("/");
  if(path.at(-1) == '') path.pop();
  return path
});

function onClickPath(index) {
  if(index == 0) 
    addressStore().setPath(pathAdressParts.value.slice(0, index + 1).join('/') + '/');
  else addressStore().setPath(pathAdressParts.value.slice(0, index + 1).join('/'));
  
}

</script>