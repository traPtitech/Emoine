<template>
  <div :class="$style.container">
    <input
      ref="inputRef"
      v-model="text"
      :class="$style.input"
      type="text"
      @keydown.enter="comment"
    />

    <div :class="$style.bottomContents">
      <button :class="$style.send" @click="comment">送信</button>
      <button :class="$style.button" @click="$emit('toggleDesc')">説明</button>
      <button :class="$style.button" @click="$emit('toggle')">
        表示/非表示
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { useStore } from '/@/store'
import { sendComment } from '/@/lib/connect'
import useShortcut from '/@/use/shortcut'

export default defineComponent({
  name: 'BottomControls',
  emits: {
    toggle: () => true,
    toggleDesc: () => true
  },
  setup() {
    const store = useStore()
    const presentationId = computed(() => store.state.presentation?.id ?? null)

    const text = ref('')
    const comment = () => {
      if (!text.value) return
      sendComment({ presentationId: presentationId.value, text: text.value })
      text.value = ''
    }

    const inputRef = ref<HTMLInputElement>()
    useShortcut({ key: ' ', ctrlKey: true }, () => {
      inputRef.value?.focus()
    })

    return { text, inputRef, comment }
  }
})
</script>

<style lang="scss" module>
.container {
  pointer-events: auto;
  width: 100%;
  background: rgba(255, 255, 255, 0.8);
}
.input {
  border-radius: 10px;
  background-color: #c9c1b1;
}
.bottomContents {
  display: flex;
  flex-direction: row-reverse;
}
.button {
  margin: 0 8px;
  padding: 0 8px;
}
</style>
