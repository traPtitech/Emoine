<template>
  <div :class="$style.container">
    <transition-group
      v-if="isCommentsShown"
      name="transition"
      tag="ul"
      :class="$style.list"
    >
      <li v-for="c in comments" :key="c">{{ c }}</li>
    </transition-group>
    <button :class="$style.button" @click="toggleIsCommentsShown">
      コメントパネル切替
    </button>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from 'vue'
import { connectTarget } from '/@/lib/connect'
import apis from '/@/lib/apis'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'CommentList',
  setup() {
    const store = useStore()
    const presentationId = computed(() => store.state.presentation?.id ?? null)

    const isCommentsShown = ref(true)
    const toggleIsCommentsShown = () => {
      isCommentsShown.value = !isCommentsShown.value
    }

    const comments = ref<string[]>([])
    onMounted(async () => {
      if (!presentationId.value) return
      const { data } = await apis.getPresentationComments(
        '' + presentationId.value
      )
      comments.value = data.map(c => c.text).reverse()

      connectTarget.addEventListener('comment', e => {
        if (!e.detail) return
        comments.value.unshift(e.detail.text)
      })
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
