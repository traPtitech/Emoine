<template>
  <div>
    <h3>Live {{ liveUrl }}</h3>
    <iframe :src="liveUrl" />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'Live',
  setup() {
    const store = useStore()
    const liveUrl = computed(
      () => `https://www.youtube.com/embed/${store.state.liveId}`
    )
    return { liveUrl }
  }
})
</script>

<style module>
.re {
  font-weight: bold;
}
</style>
