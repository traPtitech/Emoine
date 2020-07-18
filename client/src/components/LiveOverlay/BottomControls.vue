<template>
  <div :class="$style.bottomControls">
    <button @click="comment">
      コメント「ぽ」
    </button>
    <button :class="$style.toggle" @click="emit('toggle')">
      Toggle overlay
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'
import { sendComment } from '/@/lib/connect'

export default defineComponent({
  name: 'BottomControls',
  setup(_, { emit }) {
    const store = useStore()
    const presentationId = computed(() => 1) //store.state.presentation?.id ?? null)

    const comment = () => {
      sendComment({ presentationId: presentationId.value, text: 'ぽ' })
    }

    return { comment, emit }
  }
})
</script>

<style lang="scss" module>
.bottomControls {
  background: rgba(255, 255, 255, 0.8);
}
</style>
