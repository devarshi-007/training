<template>
    <p>{{ count }}</p>
</template>

<script setup>
    const {data: count} = await useFetch('https://fakestoreapi.com/users')
</script>