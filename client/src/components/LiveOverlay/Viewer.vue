<template>
  <div :class="$style.container">
    {{ viewer }}
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { viewerTarget } from '/@/lib/connect'
import apis from '/@/lib/apis'

export default defineComponent({
  name: 'Viewer',
  setup() {
    const viewer = ref(0)
    viewerTarget.addEventListener('viewer', e => {
      viewer.value = e.detail.count
    })
    onMounted(async () => {
      viewer.value = (await apis.getViewer()).data.count
    })
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
  -webkit-text-stroke-width: 2px;
  -webkit-text-stroke-color: #333;
}
</style>
