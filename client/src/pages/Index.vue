<template>
  <div :class="$style.container">
    <live-overlay :class="$style.main" @toggle-desc="toggleDesc" />

    <review v-if="isReview" :class="$style.review" />
    <descriptions v-if="showDesc" :class="$style.desc" @toggle="toggleDesc" />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import LiveOverlay from '/@/components/LiveOverlay/LiveOverlay.vue'
import Review from '/@/components/Review.vue'
import Descriptions from '/@/components/Descriptions.vue'
import { useStore } from '/@/store'
import { Status } from '/@/lib/pb'

export default defineComponent({
  name: 'Index',
  components: {
    LiveOverlay,
    Review,
    Descriptions
  },
  setup() {
    const store = useStore()
    const isReview = computed(
      () => store.state.state.status === Status.reviewing
    )
    store.dispatch.fetchLive()

    const showDesc = ref(false)
    const toggleDesc = () => {
      showDesc.value = !showDesc.value
    }

    return { isReview, showDesc, toggleDesc }
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
  height: 100%;
  width: 100%;
}
.review {
  z-index: 1;
}
.desc {
  z-index: 2;
}
</style>
