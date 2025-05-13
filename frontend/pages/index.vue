<template>
  <div class="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-blue-100 to-purple-100 px-6">
    <div  class="max-w-md w-full bg-white rounded-xl shadow-lg p-6">
      <h1 class="text-2xl font-bold text-center text-purple-700 mb-4">🎲 Добро пожаловать в&nbsp;CivilCraft</h1>

      <p class="text-gray-600 text-sm text-center mb-6">
        Построй свой город быстрее других игроков! Бросай кубики, активируй здания и собирай монеты, чтобы стать победителем.
      </p>

      <label class="block text-sm font-medium text-gray-700 mb-1" for="username">
        Имя пользователя
      </label>
      <input
          v-model="lobby.name"
          @keydown.enter="enterGame"
          id="username"
          type="text"
          placeholder="Введите ваше имя"
          class="w-full px-4 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-purple-400 focus:outline-none"
      />

      <button

          @click="enterGame"
          :disabled="!lobby.name"
          class="w-full mt-4 bg-purple-600 text-white font-semibold py-2 px-4 rounded-md shadow hover:bg-purple-700 transition disabled:opacity-50 disabled:cursor-not-allowed"
      >
        Войти
      </button>
    </div>
  </div>
</template>

<script setup>
import { jwtDecode } from "jwt-decode";
import { ref } from 'vue'
import {useLobbyStore} from "~/store/lobby.ts";
import {domain} from "~/store/game.ts";
import { TaddyWeb } from 'taddy-sdk-web'
const username = ref('')


const lobby = useLobbyStore()
onMounted( async ()=>{

  if(window && window.Telegram && window.Telegram.WebApp && window.Telegram.WebApp.initDataUnsafe && window.Telegram.WebApp.initDataUnsafe.user){
    if(telemetree){
      const telemetreeBuilder = telemetree({
        projectId: "e6c3e8dc-4187-4b19-98ee-af7cbf9eba15",
        apiKey: "370f46bb-1a6e-4d62-b3a9-ef6c9ddd5adc",
        isTelegramContext: true, // use false, if a website is not in Telegram Web App context
        logLevel: 'info', // set log level to debug if you need to. Default is info. (options: error, warn, info, debug)
        trackGroup: "medium" // set group to low if you need to. Default is medium. (options: "high", "medium", "low", false)
      });
    }

    const taddy = new TaddyWeb('miniapp-db5b3bab7ac75518afcc64b9');

    if (window.miniAppOpened) {
      Telegram.WebApp.close(); // Закрыть дубликат
    } else {
      window.miniAppOpened = true;
    }

    lobby.name = window.Telegram.WebApp.initDataUnsafe.user.username
    if(localStorage.getItem("name")){
      await fetch(`${domain}/lobby/auth?${window.Telegram.WebApp.initData}`,{method: "POST"})
          .then((data)=>data.json())
          .then((data)=>{

            lobby.userID = jwtDecode(data).user

            localStorage.setItem("token",data)
            navigateTo(
                {
                  name: 'lobby'
                })
            return true

          })
    }

  }





})



async function enterGame() {
  if (lobby.name) {
    localStorage.setItem("name",lobby.name)
      await fetch(`${domain}/lobby/auth?${window.Telegram.WebApp.initData}`,{method: "POST"})
          .then((data)=>data.json())
          .then((data)=>{
            localStorage.getItem(lobby.name)
              lobby.userID = jwtDecode(data).user

             localStorage.setItem("token",data)
            navigateTo(
                {
                  name: 'lobby'
                })

            })



  }
}
</script>