<template>
  <div :class="$style.other">
    <input
      v-model="text"
      :class="$style.input"
      type="text"
      @keydown.enter="comment"
    />
    <button :class="$style.send" @click="comment">送信</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { useStore } from '/@/store'
import { sendComment } from '/@/lib/connect'

export default defineComponent({
  name: 'CommentSender',
  setup() {
    const store = useStore()
    const presentationId = computed(() => store.state.presentation?.id ?? null)

    const text = ref('')
    const comment = () => {
      if (!text.value) return
      sendComment({ presentationId: presentationId.value, text: text.value })
      text.value = ''
    }

    return { text, comment }
  }
})
</script>

<style lang="scss" module>
.other {
  display: flex;
  flex: 1;
}
.input {
  background-color: #fff;
  flex: 1;
}
.send {
  color: #eee;
  background-color: #333;
}
</style>
