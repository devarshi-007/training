<template>
  <div v-if="pending">
    Loading ...
  </div>
  <div v-else>
    <div v-for="post in posts">
      <!-- do something -->
      {{ post }}
    </div>
  </div>
</template>

<script setup>
   const { pending, data: posts }  = useLazyFetch('https://fakestoreapi.com/users')
</script>