<template>
  <div :class="$style.presentaionStatics">
    <h3>{{ presentation.name }}</h3>
    <p>{{ presentation.description }}</p>
    <p>{{ presentation.speakers }}</p>
    <div v-if="state.reactions" :class="$style.reactions">
      <div v-for="reaction in reactions" :key="reaction.id">
        <img :class="$style.img" :src="`/assets/${reaction.imgPath}.webp`" />
        {{ reaction.count }}
      </div>
    </div>
    <div>
      <div>
        コメント数 {{ state.comments.length }}
        <button :class="$style.commentToggle" @click="toggleComments">
          コメント表示切替
        </button>
      </div>
      <div v-if="showComments" :class="$style.comments">
        <p v-for="c in state.comments" :key="c.id">{{ c.text }}</p>
      </div>
    </div>
    <div v-if="state.reviews" :class="$style.reviews">
      <div>回答数 {{ state.reviews.count }}</div>
      <div>技術 {{ state.reviews.avgSkill }}</div>
      <div>芸術 {{ state.reviews.avgArtistry }}</div>
      <div>アイデア {{ state.reviews.avgIdea }}</div>
      <div>プレゼン {{ state.reviews.avgPresentation }}</div>
    </div>
  </div>
</template>

<script lang="ts">
import {
  defineComponent,
  PropType,
  reactive,
  watchEffect,
  computed,
  ref
} from 'vue'
import apis, {
  Presentation,
  ReactionStatistics,
  Comment,
  ReviewStatistics
} from '/@/lib/apis'
import { stampToFileName } from '/@/components/LiveOverlay/reactionRenderer'

interface State {
  reactions: ReactionStatistics | null
  comments: Comment[]
  reviews: ReviewStatistics | null
}

export default defineComponent({
  name: 'PresentationStatics',
  props: {
    presentation: {
      type: Object as PropType<Presentation>,
      required: true
    }
  },
  setup(props) {
    const state = reactive<State>({
      reactions: null,
      comments: [],
      reviews: null
    })
    watchEffect(async () => {
      const id = '' + props.presentation.id
      apis.getPresentationReactions(id).then(res => {
        state.reactions = res.data
      })
      apis.getPresentationComments(id).then(res => {
        state.comments = res.data
      })
      apis.getPresentationReviews(id).then(res => {
        state.reviews = res.data
      })
    })

    const reactions = computed(() =>
      state.reactions?.counts?.map(({ stamp, count }) => ({
        id: stamp,
        imgPath: stampToFileName(stamp),
        count
      }))
    )

    const showComments = ref(false)
    const toggleComments = () => {
      showComments.value = !showComments.value
    }

    return { state, reactions, showComments, toggleComments }
  }
})
</script>

<style lang="scss" module>
.presentationStatics {
}
.commentToggle {
  border: solid 2px #333;
}
.comments {
  background: rgba(0, 0, 0, 0.3);
}
.reactions,
.reviews {
  display: flex;
  justify-content: space-between;
}
.img {
  height: 1.5em;
  width: 1.5em;
  vertical-align: bottom;
}
</style>
