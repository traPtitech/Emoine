<template>
  <div :class="$style.container">
    <comment-list v-if="isCommentsShown" :class="$style.list" />
    <button @click="popupCommentsList">コメントパネルポップアップ</button>
    <button @click="toggleIsCommentsShown">コメントパネル表示切替</button>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from 'vue'
import { connectTarget } from '/@/lib/connect'
import apis from '/@/lib/apis'
import { useStore } from '/@/store'
import CommentList from '/@/components/CommentList.vue'

export default defineComponent({
  name: 'CommentPanel',
  components: {
    CommentList
  },
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

    const popupCommentsList = () => {
      const width = 200
      const height = 300
      const top = screen.height - height
      const left = screen.width - width
      window.open(
        '/popup-comment-list',
        '_blank',
        `top=${top},left=${left},width=${width},height=${height},menubar=no,toolbar=no,locaion=no,status=no,scrollbars=no`
      )
    }

    return {
      isCommentsShown,
      toggleIsCommentsShown,
      comments,
      popupCommentsList
    }
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
  overflow-y: scroll;
}
</style>
