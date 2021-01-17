<template>
  <div :class="$style.container">
    <comment-sender v-show="show" />
    <button v-show="show" :class="$style.button" @click="$emit('toggleDesc')">
      説明
    </button>
    <button :class="$style.button" @click="$emit('toggle')">
      オーバーレイ
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { useStore } from '/@/store'
import { sendComment } from '/@/lib/connect'
import CommentSender from './CommentSender.vue'

export default defineComponent({
  name: 'BottomControls',
  components: {
    CommentSender
  },
  props: {
    show: {
      type: Boolean,
      required: true
    }
  },
  emits: {
    toggle: () => true,
    toggleDesc: () => true
  },
  setup() {
    const store = useStore()
    const presentationId = computed(() => store.state.presentation?.id ?? null)

    const text = ref('')
    const comment = () => {
      sendComment({ presentationId: presentationId.value, text: text.value })
      text.value = ''
    }

    return { text, comment }
  }
})
</script>

<style lang="scss" module>
.container {
  display: flex;
  pointer-events: auto;
  background: rgba(255, 255, 255, 0.8);
}
.button {
  margin: 0 8px;
  padding: 0 8px;
}
</style>
