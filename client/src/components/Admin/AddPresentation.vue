<template>
  <div :class="$style.addPresentation">
    <h3>追加</h3>
    <div :class="$style.inputContainer">
      <span :class="$style.inputTitle">タイトル:</span>
      <input v-model="state.name" :class="$style.input" type="text" />
    </div>
    <div :class="$style.inputContainer">
      <span :class="$style.inputTitle">説明:</span>
      <input v-model="state.description" :class="$style.input" type="text" />
    </div>
    <div :class="$style.inputContainer">
      <span :class="$style.inputTitle">発表者:</span>
      <input v-model="state.speakers" :class="$style.input" type="text" />
    </div>
    <button :class="$style.button" @click="add">追加</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue'
import apis from '/@/lib/apis'

export default defineComponent({
  name: 'AddPresentation',
  setup(_, { emit }) {
    const state = reactive({
      name: '',
      description: '',
      speakers: ''
    })

    const add = async () => {
      if (state.name === '') return
      await apis.postPresentations(state)

      state.name = ''
      state.description = ''
      state.speakers = ''
      emit('need-update')
    }

    return { state, add }
  }
})
</script>

<style lang="scss" module>
.addPresentation {
}
.inputContainer {
  display: flex;
  width: 100%;
  margin: 4px 0;
  &:first-child {
    margin-top: 0;
  }
  &:last-child {
    margin-bottom: 0;
  }
}
.inputTitle {
  margin-right: 4px;
  font-weight: bold;
}
.input {
  background-color: #fff;
  flex: 1;
}
.button {
  margin: 0 4px;
  padding: 2px 4px;
  border: solid 2px #333;
}
</style>
