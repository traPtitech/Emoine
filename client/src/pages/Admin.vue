<template>
  <routes />
  <div>Admin Page</div>
  <!--
    - 進行設定
    - すべての統計が見れる(リアクションもレビューも)
  -->
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Routes from '/@/components/Routes.vue'

export default defineComponent({
  name: 'Admin',
  components: {
    Routes
  }
})
</script>
