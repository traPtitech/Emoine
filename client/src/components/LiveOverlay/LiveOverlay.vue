<template>
  <div :class="$style.container">
    <top-controls :show="show" />
    <live-overlay-view :show="show" />
    <bottom-controls
      :class="$style.bottomControls"
      :data-is-shown="show"
      :show="show"
      @toggle="toggle"
      @toggle-desc="$emit('toggleDesc')"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import TopControls from './TopControls.vue'
import BottomControls from './BottomControls.vue'
import LiveOverlayView from './LiveOverlayView.vue'

export default defineComponent({
  name: 'LiveOverlay',
  components: {
    TopControls,
    LiveOverlayView,
    BottomControls
  },
  emits: {
    toggleDesc: () => true
  },
  setup() {
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
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  display: flex;
  flex-direction: column;
  pointer-events: none;
}
.bottomControls {
  &:not([data-is-shown='true']) {
    align-self: flex-end;
  }
}
</style>
