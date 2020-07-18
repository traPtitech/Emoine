import { Ref } from 'vue'
import { Stamp } from '/@/lib/pb'

const ReactionRatioSize = 10

export const addReaction = (
  $base: Ref<HTMLDivElement | undefined>,
  baseHeight: Ref<number>,
  stamp: Stamp
): void => {
  if (!$base.value) return

  const size = baseHeight.value / ReactionRatioSize

  const $reaction = document.createElement('div')
  $reaction.className = 'animation-reaction'
  $reaction.textContent = `:${Stamp[stamp]}:`

  $reaction.style.top = '0'
  $reaction.style.left = '0'

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
