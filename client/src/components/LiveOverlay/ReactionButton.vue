<template>
  <button :class="$style.reactionButton" @click="reaction">
    <img
      draggable="false"
      :class="$style.img"
      :src="`/assets/${fileName}.webp`"
    />
  </button>
</template>

<script lang="ts">
import { defineComponent, PropType, computed } from 'vue'
import { Stamp } from '/@/lib/pb'
import { sendReaction } from '/@/lib/connect'
import { stampToFileName } from '/@/use/reactionRenderer'

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
    const fileName = computed(() => stampToFileName(props.stamp))

    const reaction = () => {
      sendReaction({
        presentationId: props.presentationId ?? null,
        stamp: props.stamp
      })
    }

    return { fileName, reaction }
  }
})
</script>

<style lang="scss" module>
.reactionButton {
  height: min(5em, 10vh);
  width: min(5em, 10vh);
  margin: 0.5em;
}
.img {
  height: 100%;
  width: 100%;
  object-fit: contain;
  user-select: none;
}
</style>
