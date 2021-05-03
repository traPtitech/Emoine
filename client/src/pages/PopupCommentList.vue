<template>
  <div :class="$style.container">
    <comment-list />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import CommentList from '/@/components/CommentList.vue'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'PopupCommentList',
  components: {
    CommentList
  },
  setup() {
    const store = useStore()
    store.dispatch.fetchLive()

    return {}
  }
})
</script>

<style lang="scss" module>
.container {
  height: 100%;
  width: 100%;
}
</style>
