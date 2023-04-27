<template>
    <p>{{ data }}</p>

</template>

<script setup>
const { data } = await useAsyncData('users',()=>$fetch('https://fakestoreapi.com/users'))
// const { pending, data: count } = useLazyAsyncData('count', () => $fetch('/api/count'))
console.log("Hiiiiiiiiii")
console.log(1234)

console.log(1)
setTimeout(()=>console.log(2),5000);
console.log(3)
</script>