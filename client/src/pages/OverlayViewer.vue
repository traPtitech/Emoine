<template>
  <div :class="$style.container">
    <live-overlay-view
      :class="$style.overlay"
      not-show-comment-panel
      not-show-live
    />
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
  display: flex;
  height: 100%;
  width: 100%;
}
</style>
