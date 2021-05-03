<template>
  <div :class="$style.container">
    <ul v-if="isCommentsShown" :class="$style.list">
      <li v-for="c in comments" :key="c">{{ c }}</li>
    </ul>
    <button :class="$style.button" @click="toggleIsCommentsShown">
      コメントパネル切替
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue'
import { connectTarget } from '/@/lib/connect'

export default defineComponent({
  name: 'CommentList',
  setup() {
    const isCommentsShown = ref(true)
    const toggleIsCommentsShown = () => {
      isCommentsShown.value = !isCommentsShown.value
    }

    const comments = reactive<string[]>([])
    connectTarget.addEventListener('comment', e => {
      if (!e.detail) return
      comments.unshift(e.detail.text)
    })

    return { isCommentsShown, toggleIsCommentsShown, comments }
  }
})
</script>

<style lang="scss" module>
.container {
  display: flex;
  flex-direction: column;
  width: min(20vw, 20rem);
  background: rgba(255, 255, 255, 0.8);
  pointer-events: auto;
}
.list {
  height: min(30vh, 30rem);
  padding: 0.5rem;
  overflow-y: scroll;
  word-break: break-all;
}
</style>
