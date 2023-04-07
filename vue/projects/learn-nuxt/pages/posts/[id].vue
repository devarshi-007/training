<script setup>
import { ref } from 'vue'

const route = useRoute()

// When accessing /posts/1, route.params.id will be 1

definePageMeta({
    validate: async (route) => {
        // Check if the id is made up of digits
        return /^\d{5}$/.test(route.params.id)
    }
})

const title = ref('Hello World')

console.log(route.params.id)
</script>

<template>
    <Head>
        <Title>{{ title }}</Title>
        <Meta name="description" :content="title" />
        <Style type="text/css" children="body { background-color: green; }" />
    </Head>

    <p>in vue</p>
    <p>{{ $route.params.id }}</p>
</template>