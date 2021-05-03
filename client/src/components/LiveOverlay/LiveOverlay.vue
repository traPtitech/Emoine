<template>
  <div :class="$style.container">
    <top-controls :show="show" />
    <div :class="$style.overlayWrapper">
      <nico :show="show" :class="$style.overlay" />
      <comment-list v-show="show" :class="$style.commentList" />
    </div>
    <bottom-controls
      :class="$style.bottomControls"
      :data-is-shown="show"
      :show="show"
      @toggle="toggle"
      @toggle-desc="$emit('toggleDesc')"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import Nico from './Nico.vue'
import TopControls from './TopControls.vue'
import CommentList from './CommentList.vue'
import BottomControls from './BottomControls.vue'

export default defineComponent({
  name: 'LiveOverlay',
  components: {
    Nico,
    TopControls,
    CommentList,
    BottomControls
  },
  emits: {
    toggleDesc: () => true
  },
  setup() {
    const show = ref(true)
    const toggle = () => {
      show.value = !show.value
    }

    return { show, toggle }
  }
})
</script>

<style lang="scss" module>
.container {
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  display: flex;
  flex-direction: column;
  pointer-events: none;
}
.overlayWrapper {
  position: relative;
  width: 100%;
  flex: 1;
}
.overlay {
  height: 100%;
  width: 100%;
}
.commentList {
  position: absolute;
  right: 0;
  bottom: 0;
}
.bottomControls {
  &:not([data-is-shown='true']) {
    align-self: flex-end;
  }
}
</style>
