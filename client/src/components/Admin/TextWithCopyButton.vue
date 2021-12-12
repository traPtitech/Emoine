<template>
  <div :class="$style.container">
    <div :class="$style.text">{{ value }}</div>
    <button :class="$style.button" @click="copy">Copy</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'TextWithCopyButton',
  props: {
    value: {
      type: String,
      required: true
    }
  },
  setup(props) {
    const copy = async () => {
      await navigator.clipboard.writeText(props.value)
    }
    return { copy }
  }
})
</script>

<style lang="scss" module>
.container {
  display: flex;
  align-items: center;
}

.text {
  flex: 1;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.button {
  flex-shrink: 0;
  margin: 0 4px;
  padding: 2px 4px;
  border: solid 2px #333;
}
</style>
