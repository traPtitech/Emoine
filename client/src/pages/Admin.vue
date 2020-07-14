<template>
  <routes />
  <div>Admin Page</div>
  <HelloWorld msg="Hello Vue 3.0 + Vite" />
  <!--
    - 進行設定
    - すべての統計が見れる(リアクションもレビューも)
  -->
</template>

<script lang="ts">
import Routes from '/@/components/Routes.vue'
import HelloWorld from '/@/components/HelloWorld.vue'
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Admin',
  components: {
    Routes,
    HelloWorld
  }
})
</script>
