<template>
  <div :class="$style.container">
    <div :class="$style.containerMain">
      <top-controls :show="show" :class="$style.topControls" />
      <viewer-counter v-show="show" :class="$style.viewerCounter" />
      <div ref="baseEle" :class="$style.liveContainer">
        <slot />
      </div>
      <comment-panel v-show="show" :class="$style.commentPanel" />
      <descriptions v-if="showDesc" :class="$style.desc" @toggle="toggleDesc" />
    </div>
    <bottom-controls
      :class="$style.bottomControls"
      :data-is-shown="show"
      @toggle="toggle"
      @toggle-desc="toggleDesc"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import TopControls from '/@/components/LiveOverlay/TopControls.vue'
import ViewerCounter from '/@/components/LiveOverlay/ViewerCounter.vue'
import BottomControls from '/@/components/LiveOverlay/BottomControls.vue'
import CommentPanel from '/@/components/LiveOverlay/CommentPanel.vue'
import Descriptions from '/@/components/Descriptions.vue'
import { connectTarget } from '/@/lib/connect'
import useElementSize from '/@/use/elementSize'
import { useCommentRenderer } from '/@/use/commentRenderer'
import { addReaction } from '/@/use/reactionRenderer'

export default defineComponent({
  name: 'LiveOverlayView',
  components: {
    TopControls,
    ViewerCounter,
    BottomControls,
    CommentPanel,
    Descriptions
  },
  setup() {
    const show = ref(true)
    const toggle = () => {
      show.value = !show.value
    }

    const showDesc = ref(false)
    const toggleDesc = () => {
      showDesc.value = !showDesc.value
    }

    const baseEle = ref<HTMLDivElement>()
    const { height: baseHeight, width: baseWidth } = useElementSize(baseEle)
    const { addComment } = useCommentRenderer(baseEle, baseHeight)

    connectTarget.addEventListener('reaction', e => {
      if (document.visibilityState === 'hidden' || !show.value) return

      addReaction(baseEle, baseHeight, baseWidth, e.detail.stamp)
    })
    connectTarget.addEventListener('comment', e => {
      if (document.visibilityState === 'hidden' || !show.value) return

      addComment(e.detail.text)
    })

    return { baseEle, show, showDesc, toggleDesc, toggle }
  }
})
</script>

<style lang="scss" module>
.container {
  height: 100%;
  width: 100%;
}
.containerMain {
  position: relative;
  height: calc(100vh - 24px);
}
.topControls {
  z-index: 2;
  position: absolute;
  top: 0;
  width: 100%;
}
.viewerCounter {
  z-index: 2;
  position: absolute;
  top: 70px;
  right: 0;
}
.liveContainer {
  z-index: 1;
  height: 85%;
}
.desc {
  z-index: 3;
}
.commentPanel {
  z-index: 2;
  position: absolute;
  right: 0;
  bottom: 0;
}
.bottomControls {
  position: absolute;
  bottom: 0;
  &:not([data-is-shown='true']) {
    align-self: flex-end;
  }
}
</style>
