<template>
    <div>
        <p>Product details for {{ product.id }}</p>
        <p>{{ product.title }}</p>
    </div>
</template>

<script setup>
const { id } = useRoute().params
const uri = `https://fakestoreapi.com/products/`+id

const {data: product} = await useFetch(uri,{key: id});
definePageMeta({
        layout: `products`
    })
</script>

<style scoped>

</style>