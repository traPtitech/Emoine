<template>
  <div :class="$style.desc">
    <template v-if="presentation">
      <h1>{{ presentation.name }}</h1>
      <p>{{ presentation.description }}</p>
    </template>
    <p v-else>No info</p>
    <button :class="$style.button" @click="$emit('toggle')">Close</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'

export default defineComponent({
  name: 'Descriptions',
  setup() {
    const store = useStore()
    const presentation = computed(() => store.state.presentation)
    return { presentation }
  }
})
</script>

<style lang="scss" module>
.desc {
  position: absolute;
  top: 50%;
  left: 50%;
  padding: 24px;
  transform: translate(-50%, -50%);
  background: rgba(255, 255, 255, 0.8);
}
.button {
  margin: 0 4px;
  padding: 2px 4px;
  border: solid 2px #333;
}
</style>
