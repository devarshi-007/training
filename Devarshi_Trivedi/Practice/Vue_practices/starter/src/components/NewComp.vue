<template>
    <h1>This is {{ msg }}</h1>
</template>

<script setup>
import { defineProps } from 'vue';

const props = defineProps({
  foo: String
})
</script>