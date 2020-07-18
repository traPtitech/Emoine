<template>
  <div :class="$style.presentation">
    <template v-if="!showEdit">
      <h3 :class="$style.inputContainer">{{ presentation.name }}</h3>
      <p :class="$style.inputContainer">
        <span :class="$style.inputTitle">説明:</span>
        {{ presentation.description }}
      </p>
      <p :class="$style.inputContainer">
        <span :class="$style.inputTitle">発表者:</span>
        {{ presentation.speakers }}
      </p>
      <template v-if="!showInsert">
        <button :class="$style.button" @click="startEdit">編集</button>
        <button :class="$style.button" @click="startInsert">移動</button>
        <button :class="$style.button" @click="deleteThis">削除</button>
      </template>
    </template>
    <template v-if="showEdit">
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
      <button :class="$style.button" @click="finishEdit">編集完了</button>
    </template>
    <template v-if="showInsert">
      <div :class="$style.inputContainer">
        <span :class="$style.inputTitle">n番目に移動:</span>
        <input v-model="insertN" :class="$style.input" type="number" />
      </div>
      <button :class="$style.button" @click="finishInsert">移動完了</button>
    </template>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, ref, reactive } from 'vue'
import apis, { Presentation } from '/@/lib/apis'

export default defineComponent({
  name: 'Presentation',
  props: {
    presentation: {
      type: Object as PropType<Presentation>,
      required: true
    },
    presentationList: {
      type: Array as PropType<Presentation[]>,
      required: true
    }
  },
  setup(props, { emit }) {
    const state = reactive({
      name: props.presentation.name,
      description: props.presentation.description,
      speakers: props.presentation.speakers
    })
    const showEdit = ref(false)
    const startEdit = () => {
      showEdit.value = true
    }
    const finishEdit = async () => {
      const data: Partial<Presentation> = state
      showEdit.value = false
      await apis.editPresentation(
        '' + props.presentation.id,
        data as Presentation
      )
      emit('need-update')
    }

    const insertN = ref(0)
    const showInsert = ref(false)
    const startInsert = () => {
      showInsert.value = true
    }
    const finishInsert = async () => {
      showInsert.value = false
      const prev = props.presentationList[insertN.value - 1 - 1] ?? null
      const current = props.presentationList[insertN.value - 1] ?? null
      if (prev === null) {
        // eslint-disable-next-line no-console
        console.error('no prev')
        return
      }

      try {
        await apis.editPresentation('' + prev.id, {
          next: props.presentation.id
        } as Presentation)
        if (current) {
          await apis.editPresentation('' + current.id, {
            prev: props.presentation.id
          } as Presentation)
        }
        await apis.editPresentation('' + props.presentation.id, {
          prev: prev.id,
          next: current.id
        } as Presentation)
      } catch (e) {
        console.error('failed to insert', e)
      }
      emit('need-update')
    }

    const deleteThis = async () => {
      if (!window.confirm('本当に削除しますか？')) return
      await apis.deletePresentation('' + props.presentation.id)
      emit('need-update')
    }

    return {
      state,
      showEdit,
      startEdit,
      finishEdit,
      insertN,
      showInsert,
      startInsert,
      finishInsert,
      deleteThis
    }
  }
})
</script>

<style lang="scss" module>
.presentation {
  padding: 8px;
  background-color: #ccc;
  border-radius: 4px;
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
