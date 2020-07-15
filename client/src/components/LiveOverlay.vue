<template>
  <div :class="$style.overlay">
    <h3>LiveOverlay {{ liveId }}</h3>
    <!--
      - リアクション
      - コメント
    -->
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'LiveOverlay',
  setup() {
    const store = useStore()
    const liveId = computed(() => store.state.liveId)
    return { liveId }
  }
})
</script>

<style lang="scss" module>
.overlay {
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  background: rgba(255, 0, 0, 0.1);
}
</style>
