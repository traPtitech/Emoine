<template>
  <div>
    <state-settings :class="$style.settings" />
    <live-id-settings :class="$style.settings" />
    <presentations-settings :class="$style.settings" />
    <generate-token :class="$style.settings" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import StateSettings from './StateSettings.vue'
import LiveIdSettings from './LiveIdSettings.vue'
import PresentationsSettings from './PresentationsSettings.vue'
import GenerateToken from './GenerateToken.vue'

export default defineComponent({
  name: 'Settings',
  components: {
    StateSettings,
    LiveIdSettings,
    PresentationsSettings,
    GenerateToken
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
