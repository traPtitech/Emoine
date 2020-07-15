import { createDirectStore } from 'direct-vuex'
import { IState, Status } from '/@/lib/pb'
import { Presentation } from '/@/lib/apis'

interface State {
  liveId: string
  state: IState
  presentation: Presentation | null
}

const state: State = {
  liveId: 'MXMCe6J3YA8',
  state: {
    status: Status.pause,
    info: '準備中'
  },
  presentation: null
}

const { store } = createDirectStore({
  state,
  getters: {
    // countString(state) {
    //   if (state.count === 0) return 'zero'
    //   if (state.count === 1) return 'once'
    //   return `${state.count} times`
    // }
  },
  mutations: {
    // increment(state) {
    //   state.count++
    // }
  }
})

export default store.original

export type Store = typeof store
export const useStore = (): Store => store
