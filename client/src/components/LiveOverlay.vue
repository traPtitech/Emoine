<template>
  <div :class="$style.container">
    <div v-if="show" :class="$style.overlay">
      <h3>LiveOverlay {{ liveId }}</h3>
      <!--
        - リアクション
        - コメント
      -->
    </div>
    <button :class="$style.toggle" @click="toggle">
      Toggle overlay
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'LiveOverlay',
  setup() {
    const store = useStore()
    const liveId = computed(() => store.state.liveId)

    const show = ref(true)
    const toggle = () => {
      show.value = !show.value
    }

    return { liveId, show, toggle }
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
  pointer-events: all;
}
.toggle {
  position: absolute;
  bottom: 0;
  right: 0;
  background: rgba(255, 255, 255, 0.8);
  pointer-events: all;
}
</style>
