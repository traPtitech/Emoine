<template>
  <div :class="$style.topControls" :data-is-shown="show">
    <reaction-button
      v-for="stamp in stamps"
      :key="stamp"
      :presentation-id="presentationId"
      :stamp="stamp"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'
import { Stamp } from '/@/lib/pb'
import ReactionButton from './ReactionButton.vue'

export default defineComponent({
  name: 'TopControls',
  components: {
    ReactionButton
  },
  props: {
    show: {
      type: Boolean,
      required: true
    }
  },
  setup() {
    const store = useStore()
    const presentationId = computed(() => store.state.presentation?.id)

    const stamps = [
      Stamp.iine,
      Stamp.pro,
      Stamp.emoi,
      Stamp.kandoushita,
      Stamp.sugoi,
      Stamp.kami,
      Stamp.suko,
      Stamp.yosa,
      Stamp.kusa
    ]

    return { presentationId, stamps }
  }
})
</script>

<style lang="scss" module>
.topControls {
  display: flex;
  justify-content: space-around;
  background: rgba(255, 255, 255, 0.8);
  pointer-events: auto;
  &:not([data-is-shown='true']) {
    visibility: hidden;
    pointer-events: none;
  }
}
</style>
