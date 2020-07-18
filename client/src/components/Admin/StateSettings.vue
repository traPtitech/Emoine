<template>
  <div :class="$style.stateSettings">
    <button :class="$style.button" @click="next">次の発表を開始</button>
    <button v-if="!isPaused" :class="$style.button" @click="pause">
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
    const status = computed(() => store.state.state.status ?? null)
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

    return { isPaused, next, pause, resume }
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
