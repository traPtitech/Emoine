import { getWs } from '/@/lib/ws'
import { Message, IMessage, IReaction, IComment } from '/@/lib/pb'

interface ConnectEventMap {
  reaction: CustomEvent<IReaction>
  comment: CustomEvent<IComment>
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
  if (!m.reaction) return
  connectTarget.dispatchEvent(
    new CustomEvent('reaction', {
      detail: m.reaction
    })
  )
}

const onComment = (m: Message) => {
  if (!m.comment) return
  connectTarget.dispatchEvent(
    new CustomEvent('comment', {
      detail: m.comment
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
