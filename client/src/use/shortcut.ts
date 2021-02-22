import { onMounted, onUnmounted } from 'vue'

const useShortcut = (
  { key, ctrlKey }: { key: string; ctrlKey?: boolean },
  handler: (e: KeyboardEvent) => void
): void => {
  let pressing = false

  const isTargetKey = (e: KeyboardEvent) =>
    e.key === key && (ctrlKey === undefined || e.ctrlKey === ctrlKey)

  const onDown = (e: KeyboardEvent) => {
    if (pressing || !isTargetKey(e)) return
    e.preventDefault()
    handler(e)
    pressing = true
  }
  const onUp = (e: KeyboardEvent) => {
    if (!isTargetKey(e)) return
    e.preventDefault()
    pressing = false
  }

  onMounted(() => {
    window.addEventListener('keydown', onDown)
    window.addEventListener('keyup', onUp)
  })
  onUnmounted(() => {
    window.removeEventListener('keydown', onDown)
    window.removeEventListener('keyup', onUp)
  })
}

export default useShortcut
