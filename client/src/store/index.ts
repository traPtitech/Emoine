import { createDirectStore } from 'direct-vuex'
import { IState, Status } from '/@/lib/pb'
import apis, { Presentation } from '/@/lib/apis'
import { stateTarget } from '/@/lib/connect'

interface State {
  liveId: string
  state: IState
  presentation: Presentation | null
}

const state: State = {
  liveId: 'pvu2mYY8tf4',
  state: {
    status: Status.pause,
    info: '準備中'
  },
  presentation: null
}

const { store, rootActionContext } = createDirectStore({
  state,
  mutations: {
    setState(state, s: IState) {
      state.state = s
    },
    setPresentation(state, presentation: Presentation) {
      state.presentation = presentation
    },
    setLiveId(state, liveId: string) {
      state.liveId = liveId
    }
  },
  actions: {
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

export default store.original

export type Store = typeof store
export const useStore = (): Store => store
