export const addressStore = defineStore('address', () => {
  const path = ref('C:/')
  function setPath(newPath) {
    path.value = newPath
  }  
  return {path,setPath}
})