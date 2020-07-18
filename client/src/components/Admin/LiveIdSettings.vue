<template>
  <div>
    <h2>LiveId</h2>
    <input v-model="id" :class="$style.input" type="text" />
    <button :class="$style.button" @click="set">設定</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import apis from '/@/lib/apis'

export default defineComponent({
  name: 'LiveIdSettings',
  setup() {
    const id = ref('')
    onMounted(async () => {
      try {
        const { data } = await apis.getLiveId()
        id.value = data.liveId
      } catch (e) {
        // eslint-disable-next-line no-console
        console.error(e)
      }
    })

    const set = async () => {
      try {
        await apis.putLiveId({ liveId: id.value })
      } catch (e) {
        // eslint-disable-next-line no-console
        console.error(e)
      }
    }

    return { id, set }
  }
})
</script>

<style module>
.input {
  background-color: #fff;
}
.button {
  margin: 0 4px;
  padding: 2px 4px;
  border: solid 2px #333;
}
</style>
