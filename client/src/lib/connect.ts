import { getWs } from '/@/lib/ws'
import {
  Message,
  IMessage,
  IReaction,
  IComment,
  Stamp,
  IState,
  Status,
  IViewer
} from '/@/lib/pb'

type ReactionSafe = Omit<IReaction, 'stamp'> & { stamp: Stamp }
type CommentSafe = Omit<IComment, 'text'> & { text: string }
type StateSafe = Omit<IState, 'status'> & { status: Status }
type ViewerSafe = Omit<IViewer, 'count'> & { count: number }

interface ConnectEventMap {
  reaction: CustomEvent<ReactionSafe>
  comment: CustomEvent<CommentSafe>
}

interface StateEventMap {
  state: CustomEvent<StateSafe>
}

interface ViewerEventMap {
  viewer: CustomEvent<ViewerSafe>
}

interface CustomTarget<M> extends Omit<EventTarget, 'addEventListener'> {
  addEventListener<K extends keyof M>(
    name: K,
    listener: (e: M[K]) => void,
    options?: boolean | AddEventListenerOptions
  ): void
}

export const connectTarget = document.createDocumentFragment() as CustomTarget<ConnectEventMap>
export const stateTarget = document.createDocumentFragment() as CustomTarget<StateEventMap>
export const viewerTarget = document.createDocumentFragment() as CustomTarget<ViewerEventMap>

const onReaction = (m: Message) => {
  const reaction = m.reaction
  if (!reaction) return
  if (reaction.stamp === null && reaction.stamp === undefined) return
  const reactionSafe = reaction as ReactionSafe

  connectTarget.dispatchEvent(
    new CustomEvent('reaction', { detail: reactionSafe })
  )
}

const onComment = (m: Message) => {
  const comment = m.comment
  if (!comment) return
  if (!comment.text) return
  const commentSafe = comment as CommentSafe

  connectTarget.dispatchEvent(
    new CustomEvent('comment', { detail: commentSafe })
  )
}

const onState = (m: Message) => {
  const state = m.state
  if (!state) return
  if (state.status === undefined || state.status === null) return
  const stateSafe = state as CommentSafe

  stateTarget.dispatchEvent(new CustomEvent('viewer', { detail: stateSafe }))
}

const onViewer = (m: Message) => {
  const viewer = m.viewer
  if (!viewer) return
  if (!viewer.count) return
  const viewerSafe = viewer as ViewerSafe

  viewerTarget.dispatchEvent(new CustomEvent('viewer', { detail: viewerSafe }))
}

export const setup = (): void => {
  const ws = getWs()
  ws.addEventListener('message', e => {
    const message = Message.decode(new Uint8Array(e.data))
    switch (message.payload) {
      case 'reaction': {
        onReaction(message)
        break
      }
      case 'comment': {
        onComment(message)
        break
      }
      case 'state': {
        onState(message)
        break
      }
      case 'viewer': {
        onViewer(message)
        break
      }
      default:
        // eslint-disable-next-line no-console
        console.log(message)
    }
  })
}

const send = (m: IMessage): void => {
  const ws = getWs()
  const buff = Message.encode(m).finish()
  ws.send(buff.buffer.slice(buff.byteOffset, buff.byteOffset + buff.length))
}

export const sendReaction = (reaction: Required<IReaction>): void => {
  send({ reaction })
}

export const sendComment = (comment: Required<IComment>): void => {
  send({ comment })
}
