<template>
  <div>
    <h2>トークン生成</h2>
    <button
      v-if="generatedToken === ''"
      :class="$style.button"
      :disabled="isGenerating"
      @click="generate"
    >
      トークンを生成
    </button>
    <template v-else>
      <text-with-copy-button :value="indexPageWithToken" />
      <text-with-copy-button :value="overlayViewerPageWithToken" />
    </template>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from 'vue'
import apis from '/@/lib/apis'
import TextWithCopyButton from './TextWithCopyButton.vue'

export default defineComponent({
  name: 'GenerateToken',
  components: {
    TextWithCopyButton
  },
  setup() {
    const generatedToken = ref('')

    const indexPageWithToken = computed(
      () => `${location.origin}/?token=${generatedToken.value}`
    )
    const overlayViewerPageWithToken = computed(
      () => `${location.origin}/overlay-viewer?token=${generatedToken.value}`
    )

    const isGenerating = ref(false)
    const generate = async () => {
      isGenerating.value = true

      const {
        data: { token }
      } = await apis.createToken()
      generatedToken.value = token

      isGenerating.value = false
    }

    return {
      generatedToken,
      indexPageWithToken,
      overlayViewerPageWithToken,
      isGenerating,
      generate
    }
  }
})
</script>

<style lang="scss" module>
.button {
  margin: 0 4px;
  padding: 2px 4px;
  border: solid 2px #333;
}
</style>
