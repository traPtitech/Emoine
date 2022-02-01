<template>
  <div :class="$style.container">
    <live-overlay-view :class="$style.overlay" :show="show" />
    <stamp-controls :class="$style.stampControls" />
    <comment-panel :class="$style.commentPanel" />
    <bottom-controls
      :class="$style.bottomControls"
      @toggle="toggle"
      @toggle-desc="$emit('toggleDesc')"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import LiveOverlayView from '/@/components/LiveOverlay/LiveOverlayView.vue'
import StampControls from '/@/components/LiveOverlay/StampControls.vue'
import CommentPanel from '/@/components/LiveOverlay/CommentPanel.vue'
import BottomControls from '/@/components/LiveOverlay/BottomControls.vue'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'LiveOverlay',
  components: {
    LiveOverlayView,
    StampControls,
    CommentPanel,
    BottomControls
  },
  emits: {
    toggleDesc: () => true
  },
  setup() {
    const store = useStore()
    store.dispatch.fetchLive()

    const show = ref(true)
    const toggle = () => {
      show.value = !show.value
    }

    return { show, toggle }
  }
})
</script>

<style lang="scss" module>
.container {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template: 1fr min-content / 1fr min-content;
}
.overlay {
  grid-row: 1;
  grid-column: 1;
}
.stampControls {
  grid-row: 2;
  grid-column: 1;
}
.commentPanel {
  grid-row: 1;
  grid-column: 2;
  min-height: 0;
}
.bottomControls {
  grid-row: 2;
  grid-column: 2;
}
</style>
