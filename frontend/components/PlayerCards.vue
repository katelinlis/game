<template>

    <Carousel  items-to-show="3" :wrap-around="groupedCards.length>3" touch-drag mouse-drag mouse-wheel   height="150px">
      <Slide
          v-for="card in groupedCards"
          :key="card.name"
      >
      <StackedCard
          :card="card"
          :stack-count="card.count"
      />
      </Slide>
    </Carousel>

</template>

<script setup>


import 'vue3-carousel/carousel.css'
import StackedCard from './StackedCard.vue'
import {Carousel, Slide} from "vue3-carousel";

const props = defineProps({ cards: Array })

// Группировка
const groupedCards = computed(() => {
  const grouped = {}
  for (const card of props.cards) {
    if (!grouped[card.name]) {
      grouped[card.name] = { ...card, count: 1 }
    } else {
      grouped[card.name].count++
    }
  }
  return Object.values(grouped).sort((a, b) => {
    const aVal = parseInt(a.range.toString().split('-')[0])
    const bVal = parseInt(b.range.toString().split('-')[0])
    return aVal - bVal
  })
})
</script>