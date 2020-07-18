import { getWs } from '/@/lib/ws'
import { Message, IMessage, IReaction, IComment } from '/@/lib/pb'

export const ConnectTarget = document.createDocumentFragment()

const onReaction = (m: Message) => {
  if (!m.reaction) return
  ConnectTarget.dispatchEvent(
    new CustomEvent('reaction', {
      detail: m.reaction
    })
  )
}

const onComment = (m: Message) => {
  if (!m.comment) return
  ConnectTarget.dispatchEvent(
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

export const send = (m: IMessage): void => {
  const ws = getWs()
  const buff = Message.encode(m).finish()
  ws.send(buff.buffer.slice(buff.byteOffset, buff.byteOffset + buff.length))
}

const sendReaction = (reaction: Required<IReaction>): void => {
  send({ reaction })
}

const sendComment = (comment: Required<IComment>): void => {
  send({ comment })
}
