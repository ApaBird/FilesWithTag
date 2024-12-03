export const addressStore = defineStore('address', () => {
  const path = ref('C:/')
  const pathHistory = ref(['C:/'])
  function setPath(newPath) {
      path.value = newPath
      pathHistory.value.push(newPath)
  }
  
  function back() {
    path.value = pathHistory.value[pathHistory.value.length - 2]
    pathHistory.value.pop()
  }
  watch(() =>path.value, (value) => {
    if(pathHistory.value.find(x=> x === value)) return
    pathHistory.value.push(value)
  })
  return {path, setPath, back, pathHistory}
})