<template>
  <div :class="$style.container">
    <top-controls :show="show" :class="$style.topControls" />
    <div ref="baseEle" :class="$style.liveContainer">
      <live />
      <bottom-controls
        :class="$style.bottomControls"
        :data-is-shown="show"
        @toggle="toggle"
      />
    </div>
    <review v-if="isReview" :class="$style.review" />
    <descriptions v-if="showDesc" :class="$style.desc" @toggle="toggleDesc" />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import TopControls from '/@/components/LiveOverlay/TopControls.vue'
import BottomControls from '/@/components/LiveOverlay/BottomControls.vue'
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

    return { isReview, show, showDesc, toggle, toggleDesc }
  }
})
</script>

<style lang="scss" module>
.container {
  position: relative;
  height: 100%;
  width: 100%;
}
.topControls {
  z-index: 2;
}
.liveContainer {
  z-index: 1;
  height: 100vh;
}
.bottomControls {
  position: absolute;
  height: 24px;
  bottom: 0;
  &:not([data-is-shown='true']) {
    align-self: flex-end;
  }
}
.review {
  z-index: 3;
}
.desc {
  z-index: 4;
}
</style>
