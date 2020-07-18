<template>
  <div :class="$style.bottomControls">
    <div v-show="show" :class="$style.other">
      <input v-model="text" :class="$style.input" type="text" />
      <button @click="comment">
        コメント
      </button>
    </div>
    <button :class="$style.toggle" @click="emit('toggle')">
      Toggle overlay
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { useStore } from '/@/store'
import { sendComment } from '/@/lib/connect'

export default defineComponent({
  name: 'BottomControls',
  props: {
    show: {
      type: Boolean,
      required: true
    }
  },
  setup(_, { emit }) {
    const store = useStore()
    const presentationId = computed(() => 1) //store.state.presentation?.id ?? null)

    const text = ref('')
    const comment = () => {
      sendComment({ presentationId: presentationId.value, text: text.value })
      text.value = ''
    }

    return { text, comment, emit }
  }
})
</script>

<style lang="scss" module>
.bottomControls {
  display: flex;
  pointer-events: auto;
  background: rgba(255, 255, 255, 0.8);
}
.other {
  display: flex;
  flex: 1;
  pointer-events: auto;
}
.input {
  flex: 1;
}
</style>
