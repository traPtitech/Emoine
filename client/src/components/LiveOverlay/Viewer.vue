<template>
  <div :class="$style.container">
    {{ viewer }}
  </div>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'Viewer',
  setup() {
    const store = useStore()
    const viewer = computed(() => store.state.viewer)
    return { viewer }
  }
})
</script>

<style lang="scss" module>
.container {
  position: absolute;
  right: 0;
  top: calc(1rem + min(5rem, 10vh));

  display: inline-block;
  color: rgba(255, 255, 255, 0.8); /* 文字の色 */
  font-size: 36pt; /* 文字のサイズ */
  text-shadow: 2px 2px 1px #333, -2px 2px 1px #333, 2px -2px 1px #333,
    -2px -2px 1px #333, 2px 0px 1px #333, 0px 2px 1px #333, -2px 0px 1px #333,
    0px -2px 1px #333;
}
</style>
