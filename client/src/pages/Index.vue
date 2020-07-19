<template>
  <div :class="$style.container">
    <live :class="$style.live" />
    <live-overlay :class="$style.overlay" @toggle-desc="toggleDesc" />
    <review v-if="isReview" :class="$style.review" />
    <descriptions v-if="showDesc" :class="$style.desc" @toggle="toggleDesc" />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import Live from '/@/components/Live.vue'
import LiveOverlay from '/@/components/LiveOverlay/LiveOverlay.vue'
import Review from '/@/components/Review.vue'
import Descriptions from '/@/components/Descriptions.vue'
import { useStore } from '/@/store'
import { Status } from '/@/lib/pb'

export default defineComponent({
  name: 'Index',
  components: {
    Live,
    LiveOverlay,
    Review,
    Descriptions
  },
  setup() {
    const store = useStore()
    const isReview = computed(() => store.state.state === Status.reviewing)

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
  position: relative;
  height: 100%;
  width: 100%;
}
.live {
  z-index: 1;
}
.overlay {
  z-index: 2;
}
.review {
  z-index: 3;
}
.desc {
  z-index: 4;
}
</style>
