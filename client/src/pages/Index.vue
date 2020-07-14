<template>
  <routes />
  <div>Index Page (Viewer)</div>
  <img alt="Vue logo" src="/@/assets/logo.png" />
  <live />
  <!--
    - タイトルと説明と参加者とか？
    - 統計情報(今映ってるもの)
    - リアクション
    - コメント
    - レビュー
      - する
      - 編集できるように
  -->
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import Routes from '/@/components/Routes.vue'
import Live from '/@/components/Live.vue'

export default defineComponent({
  name: 'Index',
  components: {
    Routes,
    Live
  }
})
</script>
