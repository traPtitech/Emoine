<template>
  <div>Admin Page</div>
  <settings />
  <statistics />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Settings from '/@/components/Admin/Settings.vue'
import Statistics from '/@/components/Admin/Statistics.vue'

export default defineComponent({
  name: 'Admin',
  components: {
    Settings,
    Statistics
  }
})
</script>
