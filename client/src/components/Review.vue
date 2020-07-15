<template>
  <div :class="$style.review">
    <h3>Review</h3>
    <!--
      - レビュー
        - する
        - 編集できるように
    -->
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'Review',
  setup() {
    return {}
  }
})
</script>

<style lang="scss" module>
.review {
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  background: rgba(0, 255, 0, 0.1);
}
</style>
