<template>
  <div>
    進行設定
    <presentation
      v-for="presentation in presentationList"
      :key="presentation.id"
      :presentation="presentation"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, computed } from 'vue'
import apis, { Presentation } from '/@/lib/apis'
import { linkedListMapToArray } from '/@/lib/util'
import PresentaionComponent from './Presentation.vue'

export default defineComponent({
  name: 'PresentationsSettings',
  components: {
    Presentation: PresentaionComponent
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
