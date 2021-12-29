<template>
  <div :class="$style.live">
    <iframe :class="$style.iframe" :src="liveUrl" />
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

<style lang="scss" module>
.live {
  position: absolute;
  top: 0;
  left: 0;
  bottom: 24px;
  right: 0;
}
.iframe {
  height: 100%;
  width: 100%;
  border: none;
}
</style>
