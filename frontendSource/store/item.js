export const itemStore = defineStore('item', () => {
  const item = ref({Content: null, Name: 'Нет имени', Tags: []})
  function setItem(data) {
    item.value = data
  }
  return {item, setItem}
})