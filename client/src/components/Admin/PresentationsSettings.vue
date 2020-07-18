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
import PresentaionComponent from './Presentation.vue'

export default defineComponent({
  name: 'PresentationsSettings',
  components: {
    Presentation: PresentaionComponent
  },
  setup() {
    const presentationMap = ref(new Map<number, Presentation>())
    const presentationList = computed(() => {
      const list: Presentation[] = []
      const first = [...presentationMap.value.values()].find(
        p => p.prev === null
      )

      let now = first ?? null
      while (now !== null) {
        list.push(now)
        if (now.next) {
          now = presentationMap.value.get(now.next) ?? null
        } else {
          now = null
        }
      }

      return list
    })

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
