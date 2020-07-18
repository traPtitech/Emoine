<template>
  <div ref="$base" :class="$style.nico" :data-is-shown="show" />
</template>

<script lang="ts">
import { defineComponent, ref, Ref } from 'vue'
import { connectTarget } from '/@/lib/connect'

const addComment = ($base: Ref<HTMLDivElement | undefined>, text: string) => {
  if (!$base.value) return
  const baseHeight = $base.value.clientHeight
  const windowHeight = window.innerHeight

  const $comment = document.createElement('div')
  $comment.className = 'animation-comment'
  $comment.textContent = text
  $comment.addEventListener(
    'animationend',
    () => {
      $comment.remove()
    },
    { once: true }
  )
  $base.value.append($comment)
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

    connectTarget.addEventListener('reaction', e => {
      //received.value.push(e.detail.stamp.toString())
    })
    connectTarget.addEventListener('comment', e => {
      addComment($base, e.detail.text)
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
