<template>
    <div>
        <ul v-for="product in products">
            <NuxtLink :to="`/products/${product.id}`">{{ product.title }}</NuxtLink>
        </ul>
    </div>
</template>

<script setup>
    definePageMeta({
        layout: `products`
    })

    const {data: products} = await useFetch(`https://fakestoreapi.com/products`)
</script>

<style scoped>

</style>