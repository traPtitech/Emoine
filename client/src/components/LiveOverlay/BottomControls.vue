<template>
  <div :class="$style.container">
    <div :class="$style.input_container">
      <input
        ref="inputRef"
        v-model="text"
        :class="$style.input"
        type="text"
        placeholder="コメント"
        @keydown.enter="comment"
      />
    </div>
    <div :class="$style.bottomContents">
      <button :class="$style.button" @click="$emit('toggle')">
        表示/非表示
      </button>
      <button :class="$style.button" @click="$emit('toggleDesc')">説明</button>
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
    const comment = (e: KeyboardEvent) => {
      if (!e.shiftKey) return
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
  background: white;
  text-align: center;
  padding: 10px 5%;
}
.input_container {
  width: 100%;
  height: 2rem;
  border-radius: 20px;
  border: 2px solid #c9c1b1;
  &:focus-within {
    border-color: #ffd904;
  }
}
.input {
  width: 100%;
  height: 100%;
  margin-left: 1rem;
}
.bottomContents {
  display: flex;
  flex-direction: row;
  justify-content: center;
}
.button {
  margin: 0.5rem 1rem;
}
</style>
