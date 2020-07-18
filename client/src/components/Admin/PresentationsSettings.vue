<template>
  <div>
    <h2>進行設定</h2>
    <presentation
      v-for="presentation in presentationList"
      :key="presentation.id"
      :class="$style.presentation"
      :presentation="presentation"
      @need-update="refetch"
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

    const refetch = async () => {
      const { data } = await apis.getPresentations()
      presentationMap.value = new Map(data.map(p => [p.id, p]))
    }

    onMounted(() => {
      refetch()
    })

    return { presentationList, refetch }
  }
})
</script>

<style module>
.presentation {
  margin: 8px 0;
  &:first-child {
    margin-top: 0;
  }
  &:last-child {
    margin-bottom: 0;
  }
}
</style>
