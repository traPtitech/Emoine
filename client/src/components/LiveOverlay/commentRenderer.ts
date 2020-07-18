import { Ref } from 'vue'

const CommentLines = 10
let commentLineNow = 0

const incrementCommentLine = () => {
  commentLineNow++
  if (commentLineNow === CommentLines) {
    commentLineNow = 0
  }
}

export const addComment = (
  $base: Ref<HTMLDivElement | undefined>,
  baseHeight: Ref<number>,
  text: string
): void => {
  if (!$base.value) return

  const lineHeight = baseHeight.value / CommentLines

  const $comment = document.createElement('div')
  $comment.className = 'animation-comment'
  $comment.textContent = text
  $comment.style.top = `${commentLineNow * lineHeight}px`
  $comment.style.fontSize = `${lineHeight}px`

  $comment.addEventListener(
    'animationend',
    () => {
      $comment.remove()
    },
    { once: true }
  )
  $base.value.append($comment)

  incrementCommentLine()
}
