<script setup>
const route = useRoute()

// When accessing /posts/1, route.params.id will be 1

definePageMeta({
    validate: async (route) => {
        // Check if the id is made up of digits
        return /^\d{5}$/.test(route.params.id)
    }
})

console.log(route.params.id)
</script>

<template>
    <p>in vue</p>
    <p>{{ $route.params.id }}</p>
</template>