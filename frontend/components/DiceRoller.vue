<template>
  <div class="flex justify-center items-center p-4">
    <button
        :disabled="!isPlayerTurn"
        :class="[
        isPlayerTurn
          ? 'bg-yellow-400 hover:bg-yellow-500'
          : 'bg-gray-300 cursor-not-allowed',
        'text-white font-bold py-2 px-4 rounded-full shadow-md'
      ]"
        @click="handleDiceRoll"
    >
      {{ game.hasRolled ? '⏭️ Пропустить ход' : '🎲 Бросить кости' }}
    </button>
    <button
        v-if="hasTrainStation && !game.hasRolled"
        :disabled="!isPlayerTurn"
        :class="[
        isPlayerTurn
          ? 'bg-yellow-400 hover:bg-yellow-500'
          : 'bg-gray-300 cursor-not-allowed',
        'ml-4 text-white font-bold py-2 px-4 rounded-full shadow-md'
      ]"
        @click="rollTwoDice"
    >
      🎲🎲 Кинуть два кубика
    </button>
    <div v-if="dice" class="ml-4 text-2xl">
      {{ dice }}
    </div>
    <DiceModal
        :isVisible="game.modalVisibleDice"
        :diceResult="game.diceResult"
        @close="game.modalVisibleDice = false"
    />

  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useGameStore } from '~/store/game.js'
import { useLobbyStore } from '~/store/lobby.js'
import DiceModal from "~/components/DiceModal.vue";
const dice = ref(0)
const game = useGameStore()
const lobby = useLobbyStore()

const isPlayerTurn = computed(() => game.currentPlayer.id===0 || game.currentPlayer.id === lobby.userID)

const hasTrainStation = computed(() =>
    game.mainBuild.some(landmark => landmark.name === 'Вокзал' && landmark.built)
)

const modalVisible = ref(false)



function handleDiceRoll() {
  if (game.hasRolled) {
    //game.skipTurn() // Пропуск хода
    rollDice()
  } else {
    rollDice()
  }
}

function rollDice() {
  game.dice().then(res => {
    dice.value = res
  })
}

function rollTwoDice() {
  game.dice(true).then(res => {
    dice.value = res
  })
}
</script>
