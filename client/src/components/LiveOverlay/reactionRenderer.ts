import { Ref } from 'vue'
import { Stamp } from '/@/lib/pb'

const ReactionRatioSize = 8

const random = (max: number): number => Math.random() * max

export const addReaction = (
  $base: Ref<HTMLDivElement | undefined>,
  baseHeight: Ref<number>,
  baseWidth: Ref<number>,
  stamp: Stamp
): void => {
  if (!$base.value) return

  const size = baseHeight.value / ReactionRatioSize

  const $img = document.createElement('img')
  const fileName = stampToFileName(stamp)
  $img.src = `/assets/${fileName}.png`

  const $reaction = document.createElement('div')
  $reaction.className = 'animation-reaction'
  $reaction.appendChild($img)

  $reaction.style.top = `${random(baseHeight.value - 2 * size) + size}px`
  $reaction.style.left = `${random(baseWidth.value - 2 * size) + size}px`

  $reaction.style.height = `${size}px`
  $reaction.style.width = `${size}px`

  $reaction.addEventListener(
    'animationend',
    () => {
      $reaction.remove()
    },
    { once: true }
  )
  $base.value.append($reaction)
}

export const stampToFileName = (stamp: Stamp): string => {
  switch (stamp) {
    case Stamp.iine:
      return 'iine'
    case Stamp.pro:
      return 'pro'
    case Stamp.emoi:
      return 'emoi'
    case Stamp.kandoushita:
      return 'kandoushita'
    case Stamp.sugoi:
      return 'sugoi'
    case Stamp.kami:
      return 'kami'
    case Stamp.suko:
      return 'suko'
    case Stamp.yosa:
      return 'yosa'
    case Stamp.kusa:
      return 'kusa'
    default: {
      const invalid: never = stamp
      // eslint-disable-next-line no-console
      console.warn('invalid stamp', invalid)
      return 'iine'
    }
  }
}
