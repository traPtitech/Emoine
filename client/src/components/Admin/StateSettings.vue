<template>
  <div :class="$style.stateSettings">
    <p>
      <span :class="$style.title">現在の発表:</span>
      {{ presentationTitle }}
    </p>
    <button :class="$style.button" @click="next">この発表を終了</button>
    <button v-if="isSpeaking" :class="$style.button" @click="pause">
      発表を一時停止
    </button>
    <button v-if="isPaused" :class="$style.button" @click="resume">
      発表を再開
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '/@/store'
import { Status } from '/@/lib/pb'
import apis, { StateChangeType } from '/@/lib/apis'

export default defineComponent({
  name: 'StateSettings',
  setup() {
    const store = useStore()
    const presentationTitle = computed(
      () => store.state.presentation?.name ?? 'なし'
    )
    const status = computed(() => store.state.state.status ?? null)
    const isSpeaking = computed(() => status.value === Status.speaking)
    const isPaused = computed(() => status.value === Status.pause)

    const next = async () => {
      try {
        await apis.postState(StateChangeType.Next)
      } catch (e) {
        // eslint-disable-next-line no-console
        console.error(e)
      }
    }

    const pause = async () => {
      try {
        await apis.postState(StateChangeType.Pause)
      } catch (e) {
        // eslint-disable-next-line no-console
        console.error(e)
      }
    }

    const resume = async () => {
      try {
        await apis.postState(StateChangeType.Resume)
      } catch (e) {
        // eslint-disable-next-line no-console
        console.error(e)
      }
    }

    return { presentationTitle, isSpeaking, isPaused, next, pause, resume }
  }
})
</script>

<style lang="scss" module>
.title {
  font-weight: bold;
}
.button {
  margin: 0 4px;
  padding: 2px 4px;
  border: solid 2px #333;
}
</style>
