<template>
  <div :class="$style.container">
    <div v-if="show" :class="$style.overlay">
      <h3>LiveOverlay {{ liveId }}</h3>
      <div v-for="a in received" :key="a">{{ a }}</div>
    </div>
    <div :class="$style.sends">
      <button @click="reaction">
        :iine:
      </button>
      <button @click="comment">
        コメント「ぽ」
      </button>
    </div>
    <button :class="$style.toggle" @click="toggle">
      Toggle overlay
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { useStore } from '/@/store'
import { Stamp } from '/@/lib/pb'
import { connectTarget, sendReaction, sendComment } from '/@/lib/connect'

export default defineComponent({
  name: 'LiveOverlay',
  setup() {
    const store = useStore()
    const liveId = computed(() => store.state.liveId)

    const show = ref(true)
    const toggle = () => {
      show.value = !show.value
    }

    const received = ref([])
    connectTarget.addEventListener('reaction', e => {
      received.value.push(e.detail.stamp)
    })
    connectTarget.addEventListener('comment', e => {
      received.value.push(e.detail.text)
    })

    const reaction = () => {
      sendReaction({ presentationId: 1, stamp: Stamp.iine })
    }
    const comment = () => {
      sendComment({ presentationId: 1, text: 'ぽ' })
    }

    return { liveId, show, toggle, received, reaction, comment }
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
  pointer-events: none;
}
.overlay {
  height: 100%;
  width: 100%;
  background: rgba(255, 0, 0, 0.1);
  pointer-events: auto;
}
.sends {
  position: absolute;
  bottom: 0;
  left: 0;
  background: rgba(255, 255, 255, 0.8);
  pointer-events: auto;
}
.toggle {
  position: absolute;
  bottom: 0;
  right: 0;
  background: rgba(255, 255, 255, 0.8);
  pointer-events: auto;
}
</style>
