<template>
  <div :class="$style.container">
    <reaction-button
      v-for="stamp in stamps"
      :key="stamp"
      :presentation-id="presentationId"
      :stamp="stamp"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'
import { Stamp } from '/@/lib/pb'
import ReactionButton from './ReactionButton.vue'
import useShortcut from '/@/use/shortcut'
import { sendReaction } from '/@/lib/connect'

export default defineComponent({
  name: 'StampControls',
  components: {
    ReactionButton
  },
  setup() {
    const store = useStore()
    const presentationId = computed(() => store.state.presentation?.id)

    const stamps = [
      Stamp.iine,
      Stamp.pro,
      Stamp.emoi,
      Stamp.kandoushita,
      Stamp.sugoi,
      Stamp.kami,
      Stamp.suko,
      Stamp.yosa,
      Stamp.kusa
    ]

    for (const [i, stamp] of stamps.entries()) {
      useShortcut({ key: '' + i, ctrlKey: true }, () => {
        sendReaction({
          presentationId: presentationId.value ?? null,
          stamp
        })
      })
    }

    return { presentationId, stamps }
  }
})
</script>

<style lang="scss" module>
.container {
  display: flex;
  justify-content: space-around;
  background: rgba(255, 255, 255, 0.8);
  pointer-events: auto;
}
</style>
