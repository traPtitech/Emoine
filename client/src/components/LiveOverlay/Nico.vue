<template>
  <div ref="$base" :class="$style.nico" :data-is-shown="show" />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { connectTarget } from '/@/lib/connect'
import useElementSize from '/@/use/elementSize'
import { addComment } from './commentRenderer'
import { addReaction } from './reactionRenderer'

export default defineComponent({
  name: 'Nico',
  props: {
    show: {
      type: Boolean,
      required: true
    }
  },
  setup() {
    const $base = ref<HTMLDivElement>()
    const { height: baseHeight } = useElementSize($base)

    connectTarget.addEventListener('reaction', e => {
      addReaction($base, baseHeight, e.detail.stamp)
    })
    connectTarget.addEventListener('comment', e => {
      addComment($base, baseHeight, e.detail.text)
    })

    return { $base }
  }
})
</script>

<style lang="scss" module>
.nico {
  position: relative;
  color: white;
  background: rgba(255, 0, 0, 0.1);
  pointer-events: auto;
  overflow: hidden;
  &:not([data-is-shown='true']) {
    visibility: hidden;
    pointer-events: none;
  }
}
</style>
