export const itemStore = defineStore('item', () => {
  const item = ref({Content: null, Name: 'Нет имени', Tags: []})
  function setItem(data) {
    item.value = data
  }
  function closeItem() {
    item.value = {Content: null, Name: 'Нет имени', Tags: []}
  }
  return {item, setItem, closeItem}
})