<template>
  <div :class="$style.reviewContainer">
    <div :class="$style.review">
      <div v-if="err !== ''">{{ err }}</div>
      <div v-else-if="done">レビューを送信しました</div>
      <template v-else-if="presentation">
        <h3>{{ presentation.name }}</h3>
        <p>説明: {{ presentation.description }}</p>
        <div :class="$style.ranges">
          <!-- ここにLTの一覧を持つ -->
          <!-- 好きなときにレビュー画面を開いて変更できる -->
        </div>
        <button :class="$style.send" @click="send">送信</button>
      </template>
      <div v-else>レビュー画面でエラーが発生しました</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from 'vue'
import { useStore } from '/@/store'
import apis from '/@/lib/apis'

export default defineComponent({
  name: 'Review',
  setup() {
    const store = useStore()
    const presentation = computed(() => store.state.presentation)

    const done = ref(false)
    const err = ref('')

    const send = async () => {
      if (!presentation.value) return

      try {
        // 選択したLT x 3の配列をput
        await apis.putPresentationReview()
        done.value = true
      } catch (e) {
        if (e.response.status === 409) {
          err.value = '既に回答済みです'
          done.value = true
        } else {
          // eslint-disable-next-line no-console
          console.error(e)
        }
      }
    }

    return { presentation, send, done, err }
  }
})
</script>

<style lang="scss" module>
.reviewContainer {
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  backdrop-filter: blur(4px);
  z-index: 4;
}
.review {
  height: 90%;
  width: 90%;
  margin: 5%;
  padding: 24px;
  background: #fff;
}
.ranges {
  margin: 8px 0;
}
.send {
  padding: 4px 16px;
  background-color: #eee;
  border: solid 2px #333;
}
</style>
