import { Ref } from 'vue'

/*
- フォントサイズ:  画面に CommentLines 行表示できるサイズ。ただし、最低でも MinimumFontSize px
- コメント表示位置: 上から詰める
  - 現状、コメントはその文字長に関わらず画面内にいる間行を専有する(ニコニコでは文字列幅だけ専有している)
*/

const CommentLines = 10
const MinimumFontSize = 20
const LineHeight = 1.5

class FlowController {
  private countPerLine = Array(CommentLines).fill(0) // 行ごとのコメント数
  private lineCount = CommentLines // 画面に表示される行数
  private lineHeight = MinimumFontSize * LineHeight

  updateVideoHeight(height: number) {
    this.lineCount = Math.min(
      Math.floor(height / (MinimumFontSize * LineHeight)),
      CommentLines
    )
    this.lineHeight = height / this.lineCount
  }
  next(): { top: number; fontSize: number; release: () => void } {
    const countPerLine = this.countPerLine.slice(0, this.lineCount)
    const minCount = Math.min(...countPerLine)
    const nextIndex = countPerLine.findIndex(o => o === minCount)

    this.countPerLine[nextIndex]++
    return {
      top: nextIndex * this.lineHeight,
      fontSize: this.lineHeight / LineHeight,
      release: () => this.countPerLine[nextIndex]--
    }
  }
}

const flowController = new FlowController()

export const addComment = (
  baseEle: Ref<HTMLDivElement | undefined>,
  baseHeight: Ref<number>,
  text: string
): void => {
  if (!baseEle.value) return

  flowController.updateVideoHeight(baseHeight.value)
  const { top, fontSize, release } = flowController.next()

  const $comment = document.createElement('div')
  $comment.className = 'animation-comment'
  $comment.textContent = text
  $comment.style.top = `${top}px`
  $comment.style.fontSize = `${fontSize}px`

  $comment.addEventListener(
    'animationend',
    () => {
      $comment.remove()
      release()
    },
    { once: true }
  )
  baseEle.value.append($comment)
}
