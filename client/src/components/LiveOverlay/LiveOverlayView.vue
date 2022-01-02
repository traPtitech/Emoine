<template>
  <div :class="$style.container">
    <div :class="$style.containerMain">
      <top-controls :show="show" :class="$style.topControls" />
      <viewer-counter v-show="show" :class="$style.viewerCounter" />

      <nico :show="show" :class="$style.liveContainer">
        <live v-show="!notShowLive" :class="$style.live" />
      </nico>

      <comment-panel
        v-show="show && !notShowCommentPanel"
        :class="$style.commentPanel"
      />
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
import CommentPanel from '/@/components/LiveOverlay/CommentPanel.vue'
import Nico from '/@/components/LiveOverlay/Nico.vue'
import Live from '/@/components/Live.vue'
import Descriptions from '/@/components/Descriptions.vue'
import BottomControls from '/@/components/LiveOverlay/BottomControls.vue'

export default defineComponent({
  name: 'LiveOverlayView',
  components: {
    TopControls,
    ViewerCounter,
    CommentPanel,
    Nico,
    Live,
    Descriptions,
    BottomControls
  },
  props: {
    notShowCommentPanel: {
      type: Boolean,
      default: false
    },
    notShowLive: {
      type: Boolean,
      default: false
    }
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

    return { show, showDesc, toggle, toggleDesc }
  }
})
</script>

<style lang="scss" module>
.container {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
}
.containerMain {
  flex: 1;
  position: relative;
  display: grid;
  grid-template-rows: 75px 1fr;
}
.topControls {
  z-index: 2;
  width: 100%;
  position: absolute;
  grid-row: 1;
}
.viewerCounter {
  z-index: 2;
  position: absolute;
  right: 0;
  grid-row: 2;
}
.liveContainer {
  z-index: 1;
  grid-row: 2;
  overflow-y: visible;
}
.live {
  width: 100%;
  position: absolute;
  top: -75px;
  bottom: 0;
}
.commentPanel {
  z-index: 2;
  position: absolute;
  right: 0;
  bottom: 0;
}
.desc {
  z-index: 3;
}
.bottomControls {
  width: 100%;
}
</style>
