<template>
  <div :class="$style.rangeContainer">
    <label :class="$style.label">{{ label }}</label>
    <input
      :class="$style.input"
      :value="value"
      type="range"
      min="1"
      max="5"
      step="1"
      @input="onInput"
    />
    <p>{{ value }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Range',
  props: {
    label: {
      type: String,
      required: true
    },
    value: {
      type: Number,
      required: true
    }
  },
  setup(_, { emit }) {
    const onInput = (e: InputEvent) => {
      emit('update:value', (e.target as HTMLInputElement).valueAsNumber)
    }
    return { onInput }
  }
})
</script>

<style lang="scss" module>
.rangeContainer {
  display: flex;
}
.label {
  font-weight: bold;
}
.input {
  margin: 0 8px;
}
</style>
