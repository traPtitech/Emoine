<template>
  <transition-group name="transition" tag="ul" :class="$style.list">
    <li v-for="c in comments" :key="c.id" :class="$style.comment">
      {{ c.text }}
    </li>
  </transition-group>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from 'vue'
import { connectTarget } from '/@/lib/connect'
import apis from '/@/lib/apis'
import { useStore } from '/@/store'
import { CommentSafe } from '/@/lib/connect'

export default defineComponent({
  name: 'CommentList',
  setup() {
    const store = useStore()
    const presentationId = computed(() => store.state.presentation?.id ?? null)

    const comments = ref<CommentSafe[]>([])
    onMounted(async () => {
      if (presentationId.value !== null) {
        const { data } = await apis.getPresentationComments(
          '' + presentationId.value
        )
        comments.value = data.reverse()
      }

      connectTarget.addEventListener('comment', e => {
        if (!e.detail) return
        comments.value.unshift(e.detail)
      })
    })

    return { comments }
  }
})
</script>

<style lang="scss" module>
.list {
  padding: 0.5rem;
  text-align: center;
  word-break: break-all;
}
.comment {
  background-color: white;
  border-radius: 10px;
  padding: 0.5rem 10px;
  margin: 5px 0;
}
</style>
