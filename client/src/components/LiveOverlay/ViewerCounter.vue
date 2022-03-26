<template>
  <div :class="$style.container">
    {{ count }}
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { viewerTarget } from '/@/lib/connect'
import apis from '/@/lib/apis'

export default defineComponent({
  name: 'ViewerCounter',
  setup() {
    const count = ref(0)
    onMounted(async () => {
      const { data } = await apis.getViewer()
      count.value = data.count

      viewerTarget.addEventListener('viewer', e => {
        count.value = e.detail.count
      })
    })
    return { count }
  }
})
</script>

<style lang="scss" module>
.container {
  color: rgba(255, 255, 255, 0.8);
  font-size: 3rem;
  -webkit-text-stroke: 0.02em #333;
  margin: 0 1rem;
}
</style>
