<template>
  <div ref="baseEle" :class="$style.nico" />
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
  setup(props) {
    const baseEle = ref<HTMLDivElement>()
    const { height: baseHeight, width: baseWidth } = useElementSize(baseEle)
    const { addComment } = useCommentRenderer(baseEle, baseHeight)

    connectTarget.addEventListener('reaction', e => {
      if (!props.show) return

      addReaction(baseEle, baseHeight, baseWidth, e.detail.stamp)
    })
    connectTarget.addEventListener('comment', e => {
      if (!props.show) return

      addComment(e.detail.text)
    })

    return { baseEle }
  }
})
</script>

<style lang="scss" module>
.nico {
  position: relative;
  color: white;
  pointer-events: none;
}
</style>
