<template>
  <div :class="$style.desc">
    <h3>Descriptions {{ liveId }}</h3>
    <!--
      - タイトルと説明
      - 統計情報(リアクション数とか)
    -->
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'Descriptions',
  setup() {
    const store = useStore()
    const liveId = computed(() => store.state.liveId)
    return { liveId }
  }
})
</script>

<style lang="scss" module>
.desc {
  position: absolute;
  top: 0;
  right: 0;
  background: rgba(255, 255, 255, 0.3);
}
</style>
