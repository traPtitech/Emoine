<template>
  <div :class="$style.container">
    <live-overlay-view hide-live show />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import LiveOverlayView from '/@/components/LiveOverlay/LiveOverlayView.vue'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'Index',
  components: {
    LiveOverlayView
  },
  setup() {
    const store = useStore()
    store.dispatch.fetchLive()

    return {}
  }
})
</script>

<style lang="scss" module>
.container {
  height: 100%;
  width: 100%;
}
</style>
