import { getWs } from '/@/lib/ws'
import { Message, IMessage, IReaction, IComment, Stamp } from '/@/lib/pb'

type ReactionSafe = Omit<IReaction, 'stamp'> & { stamp: Stamp }
type CommentSafe = Omit<IComment, 'text'> & { text: string }

interface ConnectEventMap {
  reaction: CustomEvent<ReactionSafe>
  comment: CustomEvent<CommentSafe>
}

interface ConnectTarget extends Omit<EventTarget, 'addEventListener'> {
  addEventListener<K extends keyof ConnectEventMap>(
    name: K,
    listener: (e: ConnectEventMap[K]) => void,
    options?: boolean | AddEventListenerOptions
  ): void
}

export const connectTarget = document.createDocumentFragment() as ConnectTarget

const onReaction = (m: Message) => {
  const reaction = m.reaction
  if (!reaction) return
  if (reaction.stamp === null && reaction.stamp === undefined) return
  const reactionSafe = reaction as ReactionSafe

  connectTarget.dispatchEvent(
    new CustomEvent('reaction', {
      detail: reactionSafe
    })
  )
}

const onComment = (m: Message) => {
  const comment = m.comment
  if (!comment) return
  if (!comment.text) return
  const commentSafe = comment as CommentSafe

  connectTarget.dispatchEvent(
    new CustomEvent('comment', {
      detail: commentSafe
    })
  )
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
      default:
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
