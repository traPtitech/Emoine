<template>
  <div ref="baseEle" :class="$style.nico" :data-is-shown="show" />
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { connectTarget } from '/@/lib/connect'
import useElementSize from '/@/use/elementSize'
import { useCommentRenderer } from '/@/use/commentRenderer'
import { addReaction } from '/@/use/reactionRenderer'

export default defineComponent({
  name: 'Nico',
  props: {
    show: {
      type: Boolean,
      required: true
    }
  },
  setup() {
    const baseEle = ref<HTMLDivElement>()
    const { height: baseHeight, width: baseWidth } = useElementSize(baseEle)
    const { addComment } = useCommentRenderer(baseEle, baseHeight)

    connectTarget.addEventListener('reaction', e => {
      if (document.visibilityState === 'hidden') return

      addReaction(baseEle, baseHeight, baseWidth, e.detail.stamp)
    })
    connectTarget.addEventListener('comment', e => {
      if (document.visibilityState === 'hidden') return

      addComment(e.detail.text)
    })

    return { baseEle }
  }
})
</script>

<style lang="scss" module>
.nico {
  color: white;
  pointer-events: none;
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  &:not([data-is-shown='true']) {
    visibility: hidden;
    pointer-events: none;
  }
}
</style>
