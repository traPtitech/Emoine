<template>
  <div :class="$style.container">
    <nico v-show="show" :class="$style.overlay" />
    <top-controls :class="$style.topControls" />
    <bottom-controls :class="$style.bottomControls" @toggle="toggle" />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import Nico from './Nico.vue'
import TopControls from './TopControls.vue'
import BottomControls from './BottomControls.vue'

export default defineComponent({
  name: 'LiveOverlay',
  components: {
    Nico,
    TopControls,
    BottomControls
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
  pointer-events: none;
}
.overlay {
  height: 100%;
  width: 100%;
  background: rgba(255, 0, 0, 0.1);
  pointer-events: auto;
}
.topControls {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  pointer-events: auto;
}
.bottomControls {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  pointer-events: auto;
}
</style>
