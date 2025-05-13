<template>
  <div id="main" class="min-h-screen bg-gradient-to-b from-purple-100 to-blue-100 p-4">
    <div class="max-w-md mx-auto bg-white rounded-xl shadow p-5 space-y-6">

      <div style="display: flex; justify-content: space-between;">
        <h2 class="text-xl font-bold text-purple-700 text-center">🧩 Лобби игры </h2>
        <span @click="lobby.getLobby">Обновить</span>
      </div>


       <!-- Создание лобби -->
       <div class="flex gap-2">
         <input
             v-model="newLobbyName"
             type="text"
             placeholder="Название нового лобби"
             class="flex-1 px-3 py-2 border border-gray-300 rounded focus:ring-2 focus:ring-purple-400"
         />
         <button
             @click="createLobby"
             class="bg-purple-600 text-white px-4 py-2 rounded hover:bg-purple-700 transition"
         >
           ➕
         </button>
       </div>

       <!-- Список лобби -->
       <div v-if="lobby.lobbyList">
         <h3 class="text-sm font-semibold text-gray-600 mb-1">Доступные лобби:</h3>
         <ul class="space-y-2">
           <li
               v-for="lobby in lobby.lobbyList"
               :key="lobby.id"
               class="flex justify-between items-center border border-gray-200 p-3 rounded hover:bg-purple-50 transition"
           >
             <div>
               <div class="font-semibold">{{ lobby.name }}</div>
               <div class="text-xs text-gray-500">Игроков: {{ lobby.playerList.length }}/{{ 4 }}</div>
             </div>
             <button
                 @click="joinLobby(lobby.id)"
                 class="text-purple-600 font-medium hover:underline text-sm"
             >
               Подключиться
             </button>
           </li>
         </ul>
       </div>
       <div v-else class="text-center text-gray-500 text-sm">Нет активных лобби</div>



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
import { ref } from 'vue'
import {useLobbyStore} from "~/store/lobby.ts";
import FloatingControls from "~/components/FloatingControls.vue";
import ScriptModal from "~/components/ScriptModal.vue";

const newLobbyName = ref('')

const lobby = useLobbyStore()

onMounted(() => {
  lobby.name = localStorage.getItem("name")
  lobby.getLobby()

})

// Состояние блокировки
let creatingLobby = ref(false);



 const createLobby = async() => {
  const name = newLobbyName.value
  if (!name) {
    return;
  }

  // Проверяем состояние блокировки
  if (creatingLobby.value) {
    console.log("Лобби уже создается, подождите...");
    return;
  }

  try {
    // Устанавливаем блокировку
    creatingLobby.value = true;

    await lobby.createLobby(name);
    newLobbyName.value = '';


  }catch (error) {
    console.error("Ошибка при создании лобби:", error);
  } finally {
    // Снимаем блокировку
    creatingLobby.value = false;
  }

}

function joinLobby(id) {
  lobby.connectLobby(id)
}
</script>