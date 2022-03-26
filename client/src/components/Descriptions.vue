<template>
  <div :class="$style.container">
    <div :class="$style.button_position">
      <button :class="$style.button" @click="$emit('toggle')">
        <XIcon :class="$style.icon" />
      </button>
    </div>
    <h1>{{ presentation?.name ?? 'プレゼン情報' }}</h1>
    <p :class="$style.description">
      {{ presentation?.description ?? 'No info' }}
    </p>
    <h2>ショートカットキー</h2>
    <div :class="$style.shortcuts">
      <div>Ctrl+Space</div>
      <div>入力欄をフォーカス</div>
      <div>Shift+Enter</div>
      <div>コメント送信</div>
      <div>Ctrl+0..8</div>
      <div>スタンプを送信(左から0)</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'
import { XIcon } from '@heroicons/vue/solid'

export default defineComponent({
  name: 'Descriptions',
  components: {
    XIcon
  },
  emits: {
    toggle: () => true
  },
  setup() {
    const store = useStore()
    const presentation = computed(() => store.state.presentation)
    return { presentation }
  }
})
</script>

<style lang="scss" module>
.container {
  border-radius: 0.2rem;
  position: absolute;
  top: 50%;
  left: 50%;
  padding: 24px;
  transform: translate(-50%, -50%);
  background: rgba(245, 245, 245);
}
.description {
  padding: 4px 8px;
}
.button_position {
  position: absolute;
  top: 0;
  right: 0;
}
.button {
  margin: 1rem;
  height: 1rem;
  width: 1rem;
  display: flex;
  justify-content: center;
  align-items: center;
}
.shortcuts {
  display: grid;
  grid-template-columns: min-content 1fr;
  gap: 8px;
  padding: 4px 8px;
}
</style>
