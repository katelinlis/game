<template>
  <div class="mt-4 px-4">
    <h3
        class="text-lg font-semibold mb-2 cursor-pointer"
        @click="toggleCollapsed"
    >
      🏰 Постройки
    </h3>
    <div class="grid grid-cols-2 gap-3 transition-all duration-300">
      <div
          v-for="landmark in landmarks"
          :key="landmark.id"
          class="p-3 rounded-lg border text-sm transition cursor-pointer"
          :class="[
          landmark.built
            ? 'bg-green-100 border-green-400 text-green-800'
            : landmark.cost
              ? 'hover:bg-yellow-50 border-yellow-400 text-yellow-900'
              : 'bg-gray-100 border-gray-300 text-gray-400 cursor-not-allowed'
        ]"
          @click="tryBuild(landmark)"
      >
        <div class="font-bold">{{ landmark.name }}</div>
        <div v-show="!collapsed" class="text-xs">{{ landmark.description }}</div>
        <div  v-show="!collapsed" class="mt-1 text-right text-xs font-semibold">
          {{ landmark.built ? '✔️ Построено' : `💸 ${landmark.cost}` }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useGameStore } from '~/store/game.js'

defineProps({ landmarks: Array })
const game = useGameStore()
const collapsed = ref(false)

function toggleCollapsed() {
  collapsed.value = !collapsed.value
}

async function tryBuild(landmark) {
  if (!landmark.built && landmark.cost) {
        game.buildMain(landmark.name).catch().then(()=>{
          landmark.built = true
        })

  }
}
</script>