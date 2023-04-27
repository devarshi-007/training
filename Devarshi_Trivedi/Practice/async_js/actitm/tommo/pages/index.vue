<template>
    <p>{{main}}</p>
</template>

<script setup>
let {data:main} = useLazyAsyncData('count',()=>$fetch('https://api.publicapis.org/entries'));
</script>