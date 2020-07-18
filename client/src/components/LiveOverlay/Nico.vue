<template>
  <div :class="$style.nico">
    <h3>LiveOverlay {{ liveId }}</h3>
    <div v-for="a in received" :key="a">{{ a }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { useStore } from '/@/store'
import { connectTarget } from '/@/lib/connect'

export default defineComponent({
  name: 'Nico',
  setup() {
    const store = useStore()
    const liveId = computed(() => store.state.liveId)

    const received = ref<string[]>([])
    connectTarget.addEventListener('reaction', e => {
      if (e.detail.stamp !== null && e.detail.stamp !== undefined) {
        received.value.push(e.detail.stamp.toString())
      }
    })
    connectTarget.addEventListener('comment', e => {
      if (e.detail.text) {
        received.value.push(e.detail.text)
      }
    })

    return { liveId, received }
  }
})
</script>

<style lang="scss" module>
.nico {
}
</style>
