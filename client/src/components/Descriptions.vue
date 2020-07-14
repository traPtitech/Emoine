<template>
  <div>
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

<style module>
.re {
  font-weight: bold;
}
</style>
