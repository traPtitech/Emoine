import { computed, ref, Ref } from 'vue'

/*
- フォントサイズ: 画面に CommentLines 行表示できるサイズ。ただし、最低でも MinimumFontSize px
- コメント表示位置: 上から詰める
  - 現状、コメントはその文字長に関わらず画面内にいる間行を専有する(ニコニコでは文字列幅だけ専有している)
*/

const CommentLines = 10
const MinimumFontSize = 20
const LineHeight = 1.5

const indexOfMin = (arr: number[]): number => {
  const minCount = Math.min(...arr)
  return arr.findIndex(o => o === minCount)
}

export const useCommentRenderer = (
  baseEle: Ref<HTMLDivElement | undefined>,
  baseHeight: Ref<number>
): { addComment: (text: string) => void } => {
  const countPerLineAll = ref(Array(CommentLines).fill(0)) // 行ごとのコメント数(画面外の行を含む)

  const lineCount = computed(() =>
    Math.min(
      Math.floor(baseHeight.value / (MinimumFontSize * LineHeight)),
      CommentLines
    )
  )
  const lineHeight = computed(() => baseHeight.value / lineCount.value)
  const fontSize = computed(() => lineHeight.value / LineHeight)

  const addComment = (text: string) => {
    if (!baseEle.value) return

    const linesInScreen = countPerLineAll.value.slice(0, lineCount.value)
    const index = indexOfMin(linesInScreen)
    const top = index * lineHeight.value

    countPerLineAll.value[index]++

    const $comment = document.createElement('div')
    $comment.className = 'animation-comment'
    $comment.textContent = text
    $comment.style.top = `${top}px`
    $comment.style.fontSize = `${fontSize.value}px`

    $comment.addEventListener(
      'animationend',
      () => {
        $comment.remove()
        countPerLineAll.value[index]--
      },
      { once: true }
    )
    baseEle.value.append($comment)
  }
  return { addComment }
}
