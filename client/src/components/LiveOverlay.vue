<template>
  <div>
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

<style module>
.re {
  font-weight: bold;
}
</style>
