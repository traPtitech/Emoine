<template>
  <div :class="$style.container">
    <div :class="$style.containerMain">
      <top-controls :show="show" :class="$style.topControls" />
      <viewer-counter v-show="show" :class="$style.viewerCounter" />

      <nico :show="show" :class="$style.liveContainer">
        <slot :class="$style.liveContent" />
      </nico>

      <comment-panel v-show="show" :class="$style.commentPanel" />
      <descriptions v-if="showDesc" :class="$style.desc" @toggle="toggleDesc" />
    </div>

    <bottom-controls
      :class="$style.bottomControls"
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
import Nico from '/@/components/LiveOverlay/Nico.vue'

export default defineComponent({
  name: 'LiveOverlayView',
  components: {
    TopControls,
    ViewerCounter,
    BottomControls,
    CommentPanel,
    Descriptions,
    Nico
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

    return { show, showDesc, toggleDesc, toggle }
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
  height: calc(100% - 24px);
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
  right: 0;
}
.liveContainer {
  z-index: 1;
  height: 80vh;
  width: 100%;
  position: absolute;
  bottom: 0;
  overflow-y: visible;
}
.liveContent {
  top: 0;
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
}
</style>
