<template>
  <routes />
  <div>Admin Page</div>
  <settings />
  <statistics />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Routes from '/@/components/Routes.vue'
import Settings from '/@/components/Admin/Settings.vue'
import Statistics from '/@/components/Admin/Statistics.vue'

export default defineComponent({
  name: 'Admin',
  components: {
    Routes,
    Settings,
    Statistics
  }
})
</script>
