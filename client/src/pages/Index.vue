<template>
  <routes />
  <div>Index Page (Viewer)</div>
  <img alt="Vue logo" src="/@/assets/logo.png" />
  <live />
  <live-overlay />
  <review />
  <descriptions />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Routes from '/@/components/Routes.vue'
import Live from '/@/components/Live.vue'
import LiveOverlay from '/@/components/LiveOverlay.vue'
import Review from '/@/components/Review.vue'
import Descriptions from '/@/components/Descriptions.vue'

export default defineComponent({
  name: 'Index',
  components: {
    Routes,
    Live,
    LiveOverlay,
    Review,
    Descriptions
  }
})
</script>
