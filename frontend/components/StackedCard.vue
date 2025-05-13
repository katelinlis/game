<template>
  <div class="relative w-[100px] h-[120px]">
    <div
        v-for="(item, index) in stackCount"
        :key="index"
        class="absolute w-full transition-transform duration-300"
        :style="{ top: `${index * 4}px`, zIndex: index }"
    >
      <div
          class="rounded-xl p-2 border shadow-sm text-xs"
          :class="colorClass(card.type)"
      >
        <h3 class="font-semibold text-sm">{{ card.name }}</h3>
        <p class="text-[11px] text-gray-600">🎲 {{ card.range }}</p>
        <p class="text-sm">💰 {{ card.income }}</p>
      </div>
    </div>
    <div
        v-if="stackCount > 1"
        class="absolute z-10 -top-2 -right-2 bg-yellow-400 text-white text-xs font-bold rounded-full w-5 h-5 flex items-center justify-center shadow"
    >
      {{ stackCount }}
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  card: Object,
  stackCount: Number
})

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