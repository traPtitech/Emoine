<template>
  <div>
    <live-id-settings />
    <presentations-settings />
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
