export const filtersStore = defineStore('filters', () => {
  const filters = ref([])
  function addFilter(data) {
    if(data == null) return
    if(filters.value.includes(data)) return
    filters.value.push(data)
  }
  function removeFilter(data) {
    if(data == null) return
    filters.value = filters.value.filter(filter => filter != data)
  }
  return {filters, addFilter, removeFilter}
})