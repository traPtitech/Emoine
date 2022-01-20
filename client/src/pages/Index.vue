<template>
  <div :class="$style.container">
    <div :class="$style.main">
      <live-overlay-view :class="$style.overlay" :show="show" />
      <stamp-controls :class="$style.stampControls" />
      <comment-panel :class="$style.commentPanel" />
      <bottom-controls
        :class="$style.bottomControls"
        @toggle="toggle"
        @toggle-desc="toggleDesc"
      />
    </div>

    <review v-if="isReview" :class="$style.review" />
    <descriptions v-if="showDesc" :class="$style.desc" @toggle="toggleDesc" />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import LiveOverlayView from '/@/components/LiveOverlay/LiveOverlayView.vue'
import Review from '/@/components/Review.vue'
import Descriptions from '/@/components/Descriptions.vue'
import StampControls from '../components/LiveOverlay/StampControls.vue'
import CommentPanel from '/@/components/LiveOverlay/CommentPanel.vue'
import BottomControls from '/@/components/LiveOverlay/BottomControls.vue'
import { useStore } from '/@/store'
import { Status } from '/@/lib/pb'

export default defineComponent({
  name: 'Index',
  components: {
    LiveOverlayView,
    Review,
    Descriptions,
    StampControls,
    CommentPanel,
    BottomControls
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

    return { isReview, show, showDesc, toggle, toggleDesc }
  }
})
</script>

<style lang="scss" module>
.container {
  height: 100%;
  width: 100%;
  position: relative;
}
.main {
  z-index: 0;
  height: 100%;
  width: 100%;
  display: grid;
  grid-template: 1fr min-content / 1fr min-content;
}
.overlay {
  grid-row: 1;
  grid-column: 1;
}
.stampControls {
  grid-row: 2;
  grid-column: 1;
}
.commentPanel {
  grid-row: 1;
  grid-column: 2;
  min-height: 0;
}
.bottomControls {
  grid-row: 2;
  grid-column: 2;
}
.review {
  z-index: 1;
}
.desc {
  z-index: 2;
}
</style>
