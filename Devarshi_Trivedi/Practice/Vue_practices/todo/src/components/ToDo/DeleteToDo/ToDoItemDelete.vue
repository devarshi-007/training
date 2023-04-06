<script setup>
import removeTodo from '../Data.js'

const props = defineProps(['todoItem']);
function removedTodo(todo) {
    removeTodo(todo);
}
</script>

<template>
    <button class="buttonsee" @click="removedTodo(props.todoItem)">Remove</button>
</template>