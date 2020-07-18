<template>
  <div :class="$style.topControls">
    <button @click="reaction">
      :iine:
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'
import { Stamp } from '/@/lib/pb'
import { sendReaction } from '/@/lib/connect'

export default defineComponent({
  name: 'TopControls',
  setup() {
    const store = useStore()
    const presentationId = computed(() => 1) //store.state.presentation?.id ?? null)

    const reaction = () => {
      sendReaction({ presentationId: presentationId.value, stamp: Stamp.iine })
    }

    return { reaction }
  }
})
</script>

<style lang="scss" module>
.topControls {
  background: rgba(255, 255, 255, 0.8);
}
</style>
