<template>
  <div
      v-if="card"
      class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50"
      @click.self="$emit('close')"
  >
    <div class="bg-white rounded-xl p-6 w-80 shadow-lg">
      <h2 class="text-xl font-bold mb-2">{{ card.name }}</h2>
      <p class="text-gray-700 text-sm mb-1">Тип: {{ card.type }}</p>
      <p class="text-sm mb-2">🎲 Активация: {{ card.range }}</p>
      <p class="text-sm mb-4">💰 Доход: {{ card.income }}</p>
      <p class="text-xs text-gray-500 mb-4">{{ card.description }}</p>

      <button
          class="w-full bg-green-500 text-white py-2 rounded-lg font-semibold"
          @click="$emit('buy', card)"
      >
        Купить за {{ card.cost }} 💸
      </button>
    </div>
  </div>
</template>

<script setup>
defineProps({ card: Object })
defineEmits(['close', 'buy'])
</script>