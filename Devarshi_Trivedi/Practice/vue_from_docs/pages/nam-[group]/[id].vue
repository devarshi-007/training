<template>
    <p>{{ id  }}{{ group }}</p>

</template>

<script setup>
const route =useRoute();
const id = route.params.id;
const group = route.params.group;
</script>
