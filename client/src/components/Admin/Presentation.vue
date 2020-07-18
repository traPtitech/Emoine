<template>
  <div :class="$style.presentaion">
    <h3>{{ presentation.name }}</h3>
    <p>{{ presentation.description }}</p>
    <p>{{ presentation.speakers }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'
import { Presentation } from '/@/lib/apis'

export default defineComponent({
  name: 'Presentation',
  props: {
    presentation: {
      type: Object as PropType<Presentation>,
      required: true
    }
  }
})
</script>

<style lang="scss" module>
.presentation {
}
</style>
