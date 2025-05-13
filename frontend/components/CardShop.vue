<template>
  <div class="overflow-x-auto flex-wrap place-content-around flex gap-2 p-2 border-t">
    <div
        v-for="card in shopCards"
        :key="card.id"
        :class="colorClass(card.type)"
        class="w-[120px]   border rounded-lg p-2 shadow cursor-pointer"
        @click="$emit('select', card)"
    >
      <h4 class="text-sm font-semibold">{{ card.name }}</h4>
      <p class="text-xs text-gray-500">💰 {{ card.cost }}</p>
      <p class="text-xs">🎲 {{ card.range }}</p>
    </div>
  </div>
</template>

<script setup>
defineProps({ shopCards: Array })
defineEmits(['select'])

const colorClass = (type) => {
  switch (type) {
    case 'blue': return 'bg-blue-50'
    case 'green': return 'bg-green-50'
    case 'red': return 'bg-red-50'
    case 'purple': return 'bg-purple-50'
    default: return 'bg-gray-50'
  }
}
</script>