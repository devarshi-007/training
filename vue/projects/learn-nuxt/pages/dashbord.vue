<script setup>
import { ref } from 'vue'
definePageMeta({
    middleware: 'auth'
})
const title = ref('My title')
let countdown = 1
useSeoMeta({
    title,
    description: () => `description: ${title.value}`
})

let k = setInterval(() => {
    title.value = countdown++
}, 1000)

let kk = setTimeout(() => {
    title.value = "my title"
    clearInterval(k)
    clearTimeout(kk)
}, 6000)

</script>

<template>
    <h1>Welcome to your dashboard</h1>
</template>
