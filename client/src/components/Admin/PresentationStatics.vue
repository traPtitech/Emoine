<template>
  <div :class="$style.presentaionStatics">
    <p>統計が見れる(リアクションもレビューも)</p>
    <h3>{{ presentation.name }}</h3>
    <p>{{ presentation.description }}</p>
    <p>{{ presentation.speakers }}</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, reactive, onMounted } from 'vue'
import apis, { Presentation } from '/@/lib/apis'

export default defineComponent({
  name: 'PresentationStatics',
  props: {
    presentation: {
      type: Object as PropType<Presentation>,
      required: true
    }
  },
  setup(props) {
    const state = reactive({})
    onMounted(async () => {})
  }
})
</script>

<style lang="scss" module>
.presentationStatics {
}
</style>
