<template>
  <div :class="$style.container">
    <live-overlay-view :class="$style.overlay" />

    <review v-if="isReview" :class="$style.review" />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import LiveOverlayView from '/@/components/LiveOverlay/LiveOverlayView.vue'
import Review from '/@/components/Review.vue'
import { useStore } from '/@/store'
import { Status } from '/@/lib/pb'

export default defineComponent({
  name: 'Index',
  components: {
    LiveOverlayView,
    Review
  },
  setup() {
    const store = useStore()
    const isReview = computed(
      () => store.state.state.status === Status.reviewing
    )
    store.dispatch.fetchLive()

    return { isReview }
  }
})
</script>

<style lang="scss" module>
.container {
  height: 100%;
  width: 100%;
  position: relative;
}
.overlay {
  z-index: 1;
  height: 100%;
  width: 100%;
}
.review {
  z-index: 2;
}
</style>
