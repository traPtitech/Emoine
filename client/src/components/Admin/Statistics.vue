<template>
  <div>
    <h3>統計情報</h3>
    <presentation-statics
      v-for="presentation in presentationList"
      :key="presentation.id"
      :presentation="presentation"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed, onMounted } from 'vue'
import apis, { Presentation } from '/@/lib/apis'
import { linkedListMapToArray } from '/@/lib/util'
import PresentationStatics from './PresentationStatics.vue'

export default defineComponent({
  name: 'Statistics',
  components: {
    PresentationStatics
  },
  setup() {
    const presentationMap = ref(new Map<number, Presentation>())
    const presentationList = computed(() =>
      linkedListMapToArray(presentationMap.value)
    )

    onMounted(async () => {
      const { data } = await apis.getPresentations()
      presentationMap.value = new Map(data.map(p => [p.id, p]))
    })

    return { presentationList }
  }
})
</script>

<style module>
.re {
  font-weight: bold;
}
</style>
