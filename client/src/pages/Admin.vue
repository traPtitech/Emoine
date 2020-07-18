<template>
  <div :class="$style.wrapper">
    <ul :class="$style.tabs">
      <li :class="$style.tab" @click="changeTab('settings')">設定</li>
      <li :class="$style.tab" @click="changeTab('statistics')">統計情報</li>
    </ul>
    <settings v-if="tab === 'settings'" :class="$style.content" />
    <statistics v-else-if="tab === 'statistics'" :class="$style.content" />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import Settings from '/@/components/Admin/Settings.vue'
import Statistics from '/@/components/Admin/Statistics.vue'

type Tab = 'settings' | 'statistics'

export default defineComponent({
  name: 'Admin',
  components: {
    Settings,
    Statistics
  },
  setup() {
    const tab = ref<Tab>('settings')
    const changeTab = (name: Tab) => {
      tab.value = name
    }

    return { tab, changeTab }
  }
})
</script>

<style lang="scss" module>
.wrapper {
  max-width: 720px;
  margin: 5% auto;
}
.tabs {
  display: flex;
  margin: 24px;
  margin-bottom: 0;
}
.tab {
  flex: 1;
  padding: 4px 8px;
  background-color: #ccc;
  border-top: solid 2px #666;
  border-left: solid 2px #666;
  border-bottom: none;
  text-align: center;
  cursor: pointer;
  &:last-child {
    border-right: solid 2px #666;
  }
}
.content {
  margin: 24px;
  margin-top: 0;
  padding: 24px;
  background-color: #eee;
  border: solid 2px #666;
}
</style>
