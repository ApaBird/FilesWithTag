export const addressStore = defineStore('address', () => {
  const path = ref('C:/')
  function setPath(newPath) {
    console.log(newPath)
    path.value = newPath
  }  
  return {path,setPath}
})