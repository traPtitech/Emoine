<template>
  <div :class="$style.container">
    <div :class="$style.containerMain">
      <top-controls :show="show" :class="$style.topControls" />
      <div ref="baseEle" :class="$style.liveContainer">
        <live />
      </div>
      <comment-panel v-show="show" :class="$style.commentPanel" />
      <review v-if="isReview" :class="$style.review" />
      <descriptions v-if="showDesc" :class="$style.desc" @toggle="toggleDesc" />
    </div>
    <bottom-controls
      :class="$style.bottomControls"
      :data-is-shown="show"
      @toggle="toggle"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import TopControls from '/@/components/LiveOverlay/TopControls.vue'
import BottomControls from '/@/components/LiveOverlay/BottomControls.vue'
import CommentPanel from '/@/components/LiveOverlay/CommentPanel.vue'
import Live from '/@/components/Live.vue'
import Review from '/@/components/Review.vue'
import Descriptions from '/@/components/Descriptions.vue'
import { useStore } from '/@/store'
import { Status } from '/@/lib/pb'
import { connectTarget } from '/@/lib/connect'
import useElementSize from '/@/use/elementSize'
import { useCommentRenderer } from '/@/components/LiveOverlay/commentRenderer'
import { addReaction } from '/@/components/LiveOverlay/reactionRenderer'

export default defineComponent({
  name: 'Index',
  components: {
    TopControls,
    BottomControls,
    CommentPanel,
    Live,
    Review,
    Descriptions
  },
  setup() {
    const store = useStore()
    const isReview = computed(
      () => store.state.state.status === Status.reviewing
    )
    store.dispatch.fetchLive()

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

    return { baseEle, isReview, show, showDesc, toggle, toggleDesc }
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
.liveContainer {
  z-index: 1;
  height: 85%;
}
.commentPanel {
  z-index: 2;
  position: absolute;
  right: 0;
  bottom: 0;
}
.review {
  z-index: 3;
}
.desc {
  z-index: 4;
}
.bottomControls {
  position: absolute;
  height: 24px;
  bottom: 0;
  &:not([data-is-shown='true']) {
    align-self: flex-end;
  }
}
</style>
