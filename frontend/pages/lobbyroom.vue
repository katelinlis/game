<template>
  <div id="main" class="min-h-screen bg-gradient-to-b from-purple-50 to-blue-50 p-4">
    <div class="max-w-md mx-auto bg-white rounded-xl shadow p-5 space-y-5">
      <h2 class="text-xl font-bold text-purple-700 text-center">🧩 Лобби: {{ lobby.lobby.name }}</h2>

      <!-- Список игроков -->
      <div>
        <div style="display: flex;justify-content: space-between">
          <h3 class="text-sm font-semibold text-gray-600 mb-2">Игроки:</h3>
          <span @click="lobby.getLobbyThis()">Обновить</span>
        </div>

        <ul class="space-y-1">
          <li
              v-for="player in lobby.lobby.playerList"
              :key="player.id"
              class="flex justify-between items-center py-1 px-2 rounded border border-gray-200"
          >

            <span>{{ player.name }}</span>
            <span
                :class="player.ready ? 'text-green-600' : 'text-gray-400'"
                class="text-xs font-semibold"
            >
              {{player.ready? 'Готов' : 'Ожидает' }}
            </span>
          </li>
        </ul>
      </div>

      <!-- Кнопки управления -->
      <div class="flex flex-col gap-2 mt-4">
        <button
            @click="toggleReady"
            class="bg-purple-600 text-white py-2 px-4 rounded font-semibold hover:bg-purple-700 transition"
        >
          {{ ready ? 'Отменить готовность' : 'Готов' }}
        </button>

        <button
            v-if="isHost"
            :disabled="!allReady"
            @click="startGame"
            class="bg-green-600 text-white py-2 px-4 rounded font-semibold hover:bg-green-700 transition disabled:opacity-50"
        >
          🚀 Начать игру
        </button>

        <button
            @click="leaveLobby"
            class="text-gray-500 text-sm hover:underline mt-2"
        >
          ⬅️ Выйти из лобби
        </button>
      </div>
    </div>
    <FloatingControls />
    <ScriptModal
        v-if="lobby.scriptModal"

        @close="lobby.scriptModal = false"

    />
  </div>
</template>
<style scoped>
#main {
  padding-top: calc(var(--tg-safe-area-inset-top, 0px) + var(--tg-content-safe-area-inset-top));
}
</style>
<script setup>
import {useLobbyStore} from "~/store/lobby.ts";

import { ref } from 'vue'
import FloatingControls from "~/components/FloatingControls.vue";
import ScriptModal from "~/components/ScriptModal.vue";

const ready = ref(false)

const lobby = useLobbyStore()

onMounted(() => {
  lobby.name = localStorage.getItem("name")
  lobby.getLobbyThis()

})

//
const isHost = computed(() => {
  if(lobby.lobby.playerList.length)
  return lobby.lobby.playerList[0].id === lobby.userID

  return false
})
const allReady = computed(() =>
    lobby.lobby.playerList.length > 1 &&
     lobby.lobby.playerList.every(p => p.ready)
 )

function toggleReady() {
  lobby.ready()
//ready.value = !ready.value
}
function startGame() {
  lobby.initGame()
}
const leaveLobby = async ()=> {


  await lobby.disconnect()


}
</script>