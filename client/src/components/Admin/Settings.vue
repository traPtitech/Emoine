<template>
  <div>
    <live-id-settings :class="$style.settings" />
    <presentations-settings :class="$style.settings" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import LiveIdSettings from './LiveIdSettings.vue'
import PresentationsSettings from './PresentationsSettings.vue'

export default defineComponent({
  name: 'Settings',
  components: {
    LiveIdSettings,
    PresentationsSettings
  }
})
</script>

<style lang="scss" module>
.settings {
  margin: 24px 0;
  &:first-child {
    margin-top: 0;
  }
  &:last-child {
    margin-bottom: 0;
  }
}
</style>
