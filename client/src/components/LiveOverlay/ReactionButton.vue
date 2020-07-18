<template>
  <button :class="$style.reactionButton" @click="reaction">
    {{ name }}
  </button>
</template>

<script lang="ts">
import { defineComponent, PropType, computed } from 'vue'
import { Stamp } from '/@/lib/pb'
import { sendReaction } from '/@/lib/connect'

const stampToName = (stamp: Stamp) => {
  switch (stamp) {
    case Stamp.iine:
      return 'いいね'
    case Stamp.pro:
      return 'プロ'
    case Stamp.emoi:
      return 'エモい'
    case Stamp.kandoushita:
      return '感動した'
    case Stamp.sugoi:
      return 'すごい'
    case Stamp.kami:
      return '神'
    case Stamp.suko:
      return 'すこ'
    case Stamp.yosa:
      return 'よさ'
    case Stamp.kusa:
      return '草'
    default: {
      const invalid: never = stamp
      // eslint-disable-next-line no-console
      console.warn('invalid stamp', invalid)
    }
  }
}

export default defineComponent({
  name: 'ReactionButton',
  props: {
    // eslint-disable-next-line vue/require-default-prop
    presentationId: Number,
    stamp: {
      type: Number as PropType<Stamp>,
      required: true
    }
  },
  setup(props) {
    const name = computed(() => stampToName(props.stamp))

    const reaction = () => {
      sendReaction({
        presentationId: props.presentationId ?? null,
        stamp: props.stamp
      })
    }

    return { name, reaction }
  }
})
</script>

<style lang="scss" module>
.reactionButton {
}
</style>
