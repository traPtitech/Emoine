import { getWs } from '/@/lib/ws'
import { Message, IMessage } from '/@/lib/pb'

export const setup = (): void => {
  const ws = getWs()
  ws.addEventListener('message', e => {
    console.log(Message.decode(new Uint8Array(e.data)))
  })
}

export const send = (m: IMessage): void => {
  const ws = getWs()
  const buff = Message.encode(m).finish()
  ws.send(buff.buffer.slice(buff.byteOffset, buff.byteOffset + buff.length))
}
