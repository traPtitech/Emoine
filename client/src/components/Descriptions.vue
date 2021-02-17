<template>
  <div :class="$style.desc">
    <h1>{{ presentation?.name ?? 'プレゼン情報' }}</h1>
    <p>{{ presentation?.description ?? 'No info' }}</p>
    <h2>ショートカットキー</h2>
    <div :class="$style.shortcuts">
      <div>Ctrl+Space</div>
      <div>入力欄をフォーカス</div>
      <div>Ctrl+0～8</div>
      <div>スタンプを送信(左から0)</div>
    </div>
    <button :class="$style.button" @click="$emit('toggle')">Close</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'Descriptions',
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
.desc {
  position: absolute;
  top: 50%;
  left: 50%;
  padding: 24px;
  transform: translate(-50%, -50%);
  background: rgba(255, 255, 255, 0.8);
}
.button {
  margin: 4px;
  padding: 2px 4px;
  border: solid 2px #333;
}
.shortcuts {
  display: grid;
  grid-template-columns: min-content 1fr;
  gap: 8px;
}
</style>
