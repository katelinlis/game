<template>
  <div class="fixed bottom-4 right-4 flex flex-col gap-3 z-40">
    <button @click="logout" class="bg-white p-3 rounded-full shadow-md">⚙️</button>
    <button @click="lobby.scriptModal = !lobby.scriptModal" class="bg-white p-3 rounded-full shadow-md">📜</button>
    <button class="bg-white p-3 rounded-full shadow-md">🏆</button>
  </div>
</template>
<script setup>
import {useLobbyStore} from "~/store/lobby.js";

const lobby = useLobbyStore()
const logout = () => {
  localStorage.clear()
}
</script>