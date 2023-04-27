<script setup>
import { stringifyExpression } from '@vue/compiler-core';

const props=defineProps({
    msg:String,
});
const emits = defineEmits(['tommo']);
</script>

<template>
    <p>{{ props.msg }}</p>
    <button @click="$emit('tommo')">Hello</button>
</template>