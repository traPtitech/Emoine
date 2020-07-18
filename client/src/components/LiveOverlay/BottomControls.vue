<template>
  <div :class="$style.bottomControls">
    <input v-model="text" type="text" />
    <button @click="comment">
      コメント
    </button>
    <button :class="$style.toggle" @click="emit('toggle')">
      Toggle overlay
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { useStore } from '/@/store'
import { sendComment } from '/@/lib/connect'

export default defineComponent({
  name: 'BottomControls',
  setup(_, { emit }) {
    const store = useStore()
    const presentationId = computed(() => 1) //store.state.presentation?.id ?? null)

    const text = ref('')
    const comment = () => {
      sendComment({ presentationId: presentationId.value, text: text.value })
      text.value = ''
    }

    return { text, comment, emit }
  }
})
</script>

<style lang="scss" module>
.bottomControls {
  background: rgba(255, 255, 255, 0.8);
}
</style>
