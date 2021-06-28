<template>
  <div :class="$style.overlayWrapper">
    <nico :show="show" :class="$style.overlay" />
    <comment-panel
      v-if="!withoutCommentPanel"
      v-show="show"
      :class="$style.commentPanel"
    />
    <viewer-counter :class="$style.viewerCounter" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Nico from './Nico.vue'
import CommentPanel from './CommentPanel.vue'
import ViewerCounter from './ViewerCounter.vue'

export default defineComponent({
  name: 'LiveOverlay',
  components: {
    Nico,
    ViewerCounter,
    CommentPanel
  },
  props: {
    show: {
      type: Boolean,
      required: true
    },
    withoutCommentPanel: {
      type: Boolean,
      default: false
    }
  }
})
</script>

<style lang="scss" module>
.overlayWrapper {
  position: relative;
  width: 100%;
  flex: 1;
}
.overlay {
  height: 100%;
  width: 100%;
}
.commentPanel {
  position: absolute;
  right: 0;
  bottom: 0;
}
.viewerCounter {
  position: absolute;
  top: 0;
  right: 0;
}
</style>
