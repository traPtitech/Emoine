<template>
  <div ref="$base" :class="$style.nico" :data-is-shown="show" />
</template>

<script lang="ts">
import { defineComponent, ref, Ref } from 'vue'
import { connectTarget } from '/@/lib/connect'
import useElementSize from '/@/use/elementSize'

const CommentLines = 10
let commentLineNow = 0

const incrementCommentLine = () => {
  commentLineNow++
  if (commentLineNow === CommentLines) {
    commentLineNow = 0
  }
}

const addComment = (
  $base: Ref<HTMLDivElement | undefined>,
  baseHeight: Ref<number>,
  text: string
) => {
  if (!$base.value) return

  const lineHeight = baseHeight.value / CommentLines

  const $comment = document.createElement('div')
  $comment.className = 'animation-comment'
  $comment.textContent = text
  $comment.style.top = `${commentLineNow * lineHeight}px`
  $comment.style.fontSize = `${lineHeight}px`

  $comment.addEventListener(
    'animationend',
    () => {
      $comment.remove()
    },
    { once: true }
  )
  $base.value.append($comment)

  incrementCommentLine()
}

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
      //received.value.push(e.detail.stamp.toString())
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
