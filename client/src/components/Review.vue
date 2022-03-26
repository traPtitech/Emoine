<template>
  <div :class="$style.reviewContainer">
    <div :class="$style.review">
      <div :class="$style.ranges">
        <div v-for="presentation in presentations" :key="presentation.id">
          <label>
            <input v-model="votes" type="checkbox" :value="presentation.id" />
            {{ presentation.name }}
          </label>
        </div>
      </div>
      <button
        :class="$style.send"
        :disabled="sending || presentations === undefined"
        @click="send"
      >
        送信
      </button>
      <div v-if="err !== ''" :class="$style.error">{{ err }}</div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import apis, { Presentation } from '/@/lib/apis'

export default defineComponent({
  name: 'Review',
  setup() {
    const presentations = ref<Presentation[]>()
    const fetchPresentations = async () => {
      const res = await apis.getPresentations()
      presentations.value = res.data
    }
    fetchPresentations()

    const votes = ref<number[]>([])
    const sending = ref(false)
    const err = ref('')

    apis
      .getMyPresentationReviews()
      .then(res => {
        votes.value = res.data
      })
      .catch(() => {
        err.value = '前回の回答を取得できませんでした'
        votes.value = []
      })

    const send = async () => {
      if (votes.value.length > 3) {
        err.value = '3件までしか投票できません'
        return
      }

      sending.value = true
      try {
        // 選択したLT x 3の配列をput
        await apis.putPresentationReview(votes.value)
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } catch (e: any) {
        // eslint-disable-next-line no-console
        console.error(e)
        err.value = e.toString()
      } finally {
        sending.value = false
      }
    }

    return { presentations, votes, send, sending, err }
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
.error {
  color: red;
}
</style>
