<script setup lang="ts">

defineProps<{ msg: string }>()

</script>

<template>
  <div class="flex flex-col items-center">
  
  <img src="../assets/snake.png" alt="Snake Icon" />

  <div class="card">
    <p>
      Welcome to the competitive world of Snake! <br>
      Use the arrow keys to change direction.
    </p>
  </div>
  </div>
</template>
