<template>
  <div :class="$style.container">
    <live :class="$style.live" />
    <live-overlay :class="$style.overlay" />
    <review v-if="false" :class="$style.review" />
    <descriptions :class="$style.desc" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Live from '/@/components/Live.vue'
import LiveOverlay from '/@/components/LiveOverlay.vue'
import Review from '/@/components/Review.vue'
import Descriptions from '/@/components/Descriptions.vue'

export default defineComponent({
  name: 'Index',
  components: {
    Live,
    LiveOverlay,
    Review,
    Descriptions
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
