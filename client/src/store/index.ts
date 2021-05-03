import { createDirectStore } from 'direct-vuex'
import { IState, IViewer, Status } from '/@/lib/pb'
import apis, { User, Presentation } from '/@/lib/apis'
import { stateTarget, viewerTarget } from '/@/lib/connect'

interface State {
  liveId: string
  state: IState
  presentation: Presentation | null
  me: User | null
  viewer: number
}

const state: State = {
  liveId: 'pvu2mYY8tf4',
  state: {
    status: Status.pause,
    info: '準備中'
  },
  presentation: null,
  me: null,
  viewer: 0
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
    },
    setLiveId(state, liveId: string) {
      state.liveId = liveId
    },
    setViewer(state, viewer: IViewer) {
      state.viewer = viewer.count ?? 0
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
    async fetchLive(context) {
      const { commit } = rootActionContext(context)
      try {
        const { data } = await apis.getLiveId()
        commit.setLiveId(data.liveId)
      } catch {
        return
      }
    },
    async setState(context, state: IState) {
      const { commit, dispatch } = rootActionContext(context)
      commit.setState(state)
      if (state.presentationId) {
        await dispatch.fetchAndSetPresentation({
          presentationId: state.presentationId
        })
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

stateTarget.addEventListener('state', e => {
  store.dispatch.setState(e.detail)
})

viewerTarget.addEventListener('viewer', e => {
  store.commit.setViewer(e.detail)
})

export default store.original

export type Store = typeof store
export const useStore = (): Store => store
