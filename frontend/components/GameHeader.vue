<template>
  <header class="sticky top-0 flex justify-between items-center p-4 bg-blue-100 shadow-md z-50">
    <div @click="game.logout" class="text-xl font-bold">{{ playerName }}</div>
    <div class="flex items-center gap-2">
      💰 <span class="text-lg font-semibold">{{ coins }}</span>
    </div>
    <div class="text-sm text-gray-600">Ход: {{ currentPlayer }}</div>
  </header>
</template>
<style scoped>
header {
  padding-top: calc(var(--tg-safe-area-inset-top, 0px) + var(--tg-content-safe-area-inset-top));
}
</style>
<script setup>

import {useGameStore} from "~/store/game.js";

const game = useGameStore()

defineProps({
  playerName: String,
  coins: Number,
  currentPlayer: String
})
</script>