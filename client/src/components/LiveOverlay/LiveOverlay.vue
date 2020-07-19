<template>
  <div :class="$style.container">
    <top-controls :show="show" />
    <nico :show="show" :class="$style.overlay" />
    <bottom-controls
      :class="$style.bottomControls"
      :data-is-shown="show"
      :show="show"
      @toggle="toggle"
      @toggle-desc="$emit('toggle-desc')"
    />
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
  display: flex;
  flex-direction: column;
  pointer-events: none;
}
.overlay {
  width: 100%;
  flex: 1;
}
.bottomControls {
  &:not([data-is-shown='true']) {
    align-self: flex-end;
  }
}
</style>
