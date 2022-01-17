<template>
  <div :class="$style.container">
    <viewer-counter v-show="show" :class="$style.viewerCounter" />
    <nico v-show="show" :class="$style.nico" />
    <live v-if="!hideLive" :class="$style.live" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ViewerCounter from '/@/components/LiveOverlay/ViewerCounter.vue'
import Nico from '/@/components/LiveOverlay/Nico.vue'
import Live from '/@/components/Live.vue'

export default defineComponent({
  name: 'LiveOverlayView',
  components: {
    ViewerCounter,
    Nico,
    Live
  },
  props: {
    show: {
      type: Boolean,
      required: true
    },
    hideLive: {
      type: Boolean,
      default: false
    }
  }
})
</script>

<style lang="scss" module>
.container {
  height: 100%;
  width: 100%;
  position: relative;
  overflow: hidden;
}
.viewerCounter {
  z-index: 3;
  position: absolute;
  right: 0;
  top: 0;
}
.nico {
  z-index: 2;
}
.live {
  z-index: 1;
}
</style>
