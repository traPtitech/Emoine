<template>
  <div :class="$style.reviewContainer">
    <div :class="$style.review">
      <div v-if="err !== ''">{{ err }}</div>
      <div v-else-if="done">レビューを送信しました</div>
      <template v-else-if="presentation">
        <h3>{{ presentation.name }}</h3>
        <p>説明: {{ presentation.description }}</p>
        <div :class="$style.ranges">
          <range v-model:value="state.skill" label="技術" />
          <range v-model:value="state.artistry" label="芸術" />
          <range v-model:value="state.idea" label="アイデア" />
          <range v-model:value="state.presentation" label="プレゼン" />
        </div>
        <button :class="$style.send" @click="send">送信</button>
      </template>
      <div v-else>レビュー画面でエラーが発生しました</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, reactive, ref } from 'vue'
import { useStore } from '/@/store'
import apis, { Review } from '/@/lib/apis'
import Range from './Range.vue'

export default defineComponent({
  name: 'Review',
  components: {
    Range
  },
  setup() {
    const store = useStore()
    const presentation = computed(() => store.state.presentation)

    const state = reactive<Review>({
      userId: '', // 未使用だけど型定義が対応してないので記述
      skill: 3,
      artistry: 3,
      idea: 3,
      presentation: 3
    })

    const done = ref(false)
    const err = ref('')

    const send = async () => {
      if (!presentation.value) return

      try {
        const res = await apis.postPresentationReview(
          '' + presentation.value.id,
          state
        )
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

    return { presentation, state, send, done, err }
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
