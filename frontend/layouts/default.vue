<script setup lang="ts">

</script>

<template>

    <div class="gWrapper">
      <slot />
    </div>

</template>

<style scoped>
.gWrapper {

   padding-right: var(--tg-safe-area-inset-right, 0px);
   padding-bottom: var(--tg-safe-area-inset-bottom, 0px);
   padding-left: var(--tg-safe-area-inset-left, 0px);

}
</style>