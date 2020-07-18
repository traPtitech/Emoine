import { ref, Ref, watchEffect } from 'vue'
import { throttle } from 'throttle-debounce'

// 注意: windowのリサイズだけ見る
const useElementSize = (
  $element: Ref<HTMLElement | undefined>
): { width: Ref<number>; height: Ref<number> } => {
  const width = ref($element.value?.clientWidth ?? 0)
  const height = ref($element.value?.clientHeight ?? 0)

  watchEffect(() => {
    width.value = $element.value?.clientWidth ?? 0
    height.value = $element.value?.clientHeight ?? 0
  })

  window.addEventListener(
    'resize',
    throttle(100, true, () => {
      width.value = $element.value?.clientWidth ?? 0
      height.value = $element.value?.clientHeight ?? 0
    })
  )

  return { width, height }
}

export default useElementSize
