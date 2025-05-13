<template>
  <div class="min-h-screen bg-gray-100">

   
    <GameHeader
        :playerName="lobby.name"
        :coins="game.coins"
        :currentPlayer="game.currentPlayer.name"
    />
    <PlayerLandmarks :landmarks="game.mainBuild" />
    <PlayerCards :cards="game.cards" />

    <DiceRoller />

    <CardShop :shopCards="game.shop" @select="openCard" />

    <CardModal
        v-if="game.selectedCard"
        :card="game.selectedCard"
        @close="game.selectedCard = null"
        @buy="handleBuy"
    />

    <ScriptModal
        v-if="lobby.scriptModal"

        @close="lobby.scriptModal = false"

    />

    <FloatingControls />
  </div>
</template>


<script setup>
import { useGameStore } from '@/store/game'
import GameHeader from '@/components/GameHeader.vue'
import PlayerCards from '@/components/PlayerCards.vue'
import DiceRoller from '@/components/DiceRoller.vue'
import CardShop from '@/components/CardShop.vue'
import CardModal from '@/components/CardModal.vue'
import FloatingControls from '@/components/FloatingControls.vue'
import {useLobbyStore} from "~/store/lobby.js";
import ScriptModal from "~/components/ScriptModal.vue";

const game = useGameStore()
const lobby = useLobbyStore()

onMounted(() => {
  console.log(import.meta.browser)
  if(import.meta.browser){

    game.initGame()
  }





})

function openCard(card) {
  game.selectedCard = card
}

function handleBuy(card) {
  game.buyCard(card)
  game.selectedCard = null
}
</script>