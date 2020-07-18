<template>
  <div :class="$style.nico" :data-is-shown="show">
    <div v-for="a in received" :key="a">{{ a }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { connectTarget } from '/@/lib/connect'

export default defineComponent({
  name: 'Nico',
  props: {
    show: {
      type: Boolean,
      required: true
    }
  },
  setup() {
    const received = ref<string[]>([])
    connectTarget.addEventListener('reaction', e => {
      if (e.detail.stamp !== null && e.detail.stamp !== undefined) {
        received.value.push(e.detail.stamp.toString())
      }
    })
    connectTarget.addEventListener('comment', e => {
      if (e.detail.text) {
        received.value.push(e.detail.text)
      }
    })

    return { received }
  }
})
</script>

<style lang="scss" module>
.nico {
  background: rgba(255, 0, 0, 0.1);
  pointer-events: auto;
  &:not([data-is-shown='true']) {
    visibility: hidden;
    pointer-events: none;
  }
}
</style>
