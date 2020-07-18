import { createDirectStore } from 'direct-vuex'
import { IState, Status } from '/@/lib/pb'
import apis, { User, Presentation } from '/@/lib/apis'

interface State {
  liveId: string
  state: IState
  presentation: Presentation | null
  me: User | null
}

const state: State = {
  liveId: 'MXMCe6J3YA8',
  state: {
    status: Status.pause,
    info: '準備中'
  },
  presentation: null,
  me: null
}

const { store, rootActionContext } = createDirectStore({
  state,
  getters: {
    // countString(state) {
    //   if (state.count === 0) return 'zero'
    //   if (state.count === 1) return 'once'
    //   return `${state.count} times`
    // }
  },
  mutations: {
    setMe(state, me: User) {
      state.me = me
    },
    setState(state, s: IState) {
      state.state = s
    },
    setPresentation(state, presentation: Presentation) {
      state.presentation = presentation
    }
  },
  actions: {
    async fetchMe(context) {
      const { commit } = rootActionContext(context)
      try {
        const { data } = await apis.getMe()
        commit.setMe(data)
      } catch {
        return
      }
    },
    async fetchAndSetPresentation(
      context,
      { presentationId }: { presentationId: number }
    ) {
      const { commit } = rootActionContext(context)
      const { data: presentation } = await apis.getPresentation(
        '' + presentationId
      )
      commit.setPresentation(presentation)
    }
  }
})

export default store.original

export type Store = typeof store
export const useStore = (): Store => store
