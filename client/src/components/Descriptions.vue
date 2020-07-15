<template>
  <div v-if="presentation" :class="$style.desc">
    <h1>{{ presentation.name }}</h1>
    <p>{{ presentation.description }}</p>
    <!--
      - 統計情報(リアクション数とか)
    -->
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
  top: 0;
  right: 0;
  background: rgba(255, 255, 255, 0.3);
}
</style>
