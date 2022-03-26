<template>
  <div :class="$style.container">
    <header :class="$style.header">
      <h2>チャット</h2>
      <button :class="$style.button" @click="popupCommentsList">
        <ExternalLinkIcon :class="$style.icon" />
      </button>
    </header>
    <comment-list :class="$style.list" />
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, ref } from 'vue'
import { connectTarget } from '/@/lib/connect'
import apis from '/@/lib/apis'
import { useStore } from '/@/store'
import CommentList from '/@/components/CommentList.vue'
import { ExternalLinkIcon } from '@heroicons/vue/outline'

export default defineComponent({
  name: 'CommentPanel',
  components: {
    CommentList,
    ExternalLinkIcon
  },
  setup() {
    const store = useStore()
    const presentationId = computed(() => store.state.presentation?.id ?? null)

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
      const width = 500
      const height = 700
      const top = screen.height - height
      const left = screen.width - width
      window.open(
        '/popup-comment-list',
        '_blank',
        `top=${top},left=${left},width=${width},height=${height},menubar=no,toolbar=no,locaion=no,status=no,scrollbars=no`
      )
    }

    return {
      comments,
      popupCommentsList
    }
  }
})
</script>

<style lang="scss" module>
.header {
  display: flex;
  justify-content: center;
  align-items: center;
}
.container {
  display: flex;
  flex-direction: column;
  width: min(20vw, 20rem);
  background: #fff9ed;
  pointer-events: auto;
}
.list {
  flex: 1;
  overflow-y: scroll;
}
.button {
  height: min-content;
  display: flex;
  justify-content: end;
  align-items: center;
}
.icon {
  height: 1.3rem;
  margin: 0.5rem;
}
</style>
