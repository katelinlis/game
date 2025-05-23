<template>
  <div
      class="fixed inset-0 bg-opacity-40 flex items-center justify-center z-50"
      @click.self="$emit('close')"
  >
    <div class="bg-white rounded-xl p-6 w-full max-w-lg shadow-lg max-h-[90vh] overflow-y-auto relative">
      <!-- Шапка -->
      <div class="sticky top-0 bg-white z-10 border-b border-gray-200 shadow-sm">
        <!-- Кнопка закрытия -->
        <button
            @click="$emit('close')"
            class="absolute top-2 right-2 text-gray-400 hover:text-gray-600"
        >
          ✖️
        </button>

        <!-- Заголовок -->
        <h2 class="text-2xl font-bold text-purple-700 py-4 text-center">📜 Правила игры</h2>
      </div>

      <!-- Простая цель игры -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-2">🎯 Основная цель</h3>
        <p class="text-gray-600">Задача игрока — быстрее других построить все свои <strong>достопримечательности</strong>.</p>
      </div>

      <!-- Ход игры -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-2">🔄 Как проходит ход</h3>
        <ol class="list-decimal list-inside text-gray-600 space-y-2">
          <li>
            <strong>Бросок кубиков:</strong>
            <ul class="list-disc list-inside ml-4">
              <li>В начале своего хода вы можете бросить 1 кубик.</li>
              <li>После постройки <span class="text-purple-600 font-semibold">Вокзала</span> вы сможете бросать по 2 кубика одновременно.</li>
            </ul>
          </li>
          <li>
            <strong>Эффект карт:</strong>
            Карты, чей номер совпал с итогом кубиков, активируются. В зависимости от их типа доход получает вы или соперники.
          </li>
          <li>
            <strong>Покупка карты:</strong>
            В конце хода вы можете приобрести одну карту (здание или достопримечательность), оплатив её стоимость.
          </li>
        </ol>
      </div>

      <!-- Типы карт -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-2">🎨 Типы карт</h3>
        <p class="text-gray-600 mb-2">Карты разделяются на 4 основных категории:</p>
        <ul class="space-y-2">
          <li><span class="text-blue-600 font-semibold">🟦 Синие:</span> Приносят доход в любой ход.</li>
          <li><span class="text-green-600 font-semibold">🟩 Зелёные:</span> Приносят доход только в ваш ход.</li>
          <li><span class="text-red-600 font-semibold">🟥 Красные:</span> Активируются в ход соперников и забирают монеты у них.</li>
          <li v-if="false"><span class="text-purple-600 font-semibold">🟪 Фиолетовые:</span> Дают уникальные эффекты, активируются только в ваш ход.</li>
        </ul>
      </div>

      <!-- Достопримечательности -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-2">🏛 Уникальные достопримечательности</h3>
        <p class="text-gray-600 mb-2">Каждый игрок начинает с 4 закрытыми достопримечательностями, которые можно открыть за монеты. Каждая из них даёт полезные бонусы:</p>
        <ul class="list-disc list-inside text-gray-600 space-y-2 ml-4">
          <li><strong>Порт:</strong> Позволяет использовать некоторые карты.</li>
          <li><strong>Вокзал:</strong> Позволяет бросать 2 кубика.</li>
          <li><strong>Торговый центр:</strong> Приносит на 1 монету больше с красных и зелёных карт.</li>
          <li><strong>Парк развлечений:</strong> Даёт возможность сделать бросить еще раз кубики, если выбросили дубль.</li>
          <li><strong>Телебашня:</strong> в разработке</li>
          <li><strong>Аэропорт:</strong> в разработке</li>
        </ul>
      </div>

      <!-- Советы -->
      <div class="mb-6">
        <h3 class="text-lg font-semibold text-gray-800 mb-2">🧱 Советы для победы</h3>
        <ul class="list-disc list-inside text-gray-600 space-y-2">
          <li>Стройте экономику, покупая полезные карты, которые приносят доход.</li>
          <li>Комбинируйте эффекты карт и достопримечательностей для максимальной эффективности.</li>
          <li>Не забывайте обороняться от соперников с помощью красных карт.</li>
          <li>Помните: достопримечательности приближают победу!</li>
        </ul>
      </div>

      <!-- Примеры карт -->
      <div>
        <h3 class="text-lg font-semibold text-gray-800 mb-2">📝 Примеры карт</h3>
        <ul class="space-y-2">
          <li><strong>Пекарня:</strong> 🟩 Активна при 2-3. Приносит 1 монету.</li>
          <li><strong>Кафе:</strong> 🟥 Активно при 3. Забирает 1 монету у активного игрока.</li>
          <li><strong>Лесопилка:</strong> 🟦 Активна при 4. Приносит 1 монету в любой ход.</li>
          <li v-if="false"><strong>Стадион:</strong> 🟪 Активен при 6. Берёт по 2 монеты у всех других игроков.</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup>
defineEmits(['close'])
</script>

<style scoped>
/* Анимация и стили для лучшего UX */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: scale(0.95);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}

div[role="dialog"] .bg-white {
  animation: fadeIn 0.3s ease-out;
}
</style>