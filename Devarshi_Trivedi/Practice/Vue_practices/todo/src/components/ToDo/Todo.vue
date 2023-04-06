<script setup>
import { ref } from 'vue'
import data from './Data.js'
import TodoItemAdd from './AddToDo/ToDoItemAdd.vue'
import RemoveTodo from './DeleteToDo/ToDoItemDelete.vue'

const todoItem = ref('')
</script>

<template>
    <div>
        <input type="text" v-model="todoItem" />
        <TodoItemAdd @reset-text="()=>todoItem=''" :item="todoItem"/>
    </div>
    <div>
        <ul class="todoItem" v-for="todo in data" :key="todo.id">
           <div class="textItem">
                <span>{{ todo.item }}</span>
            <RemoveTodo :todoItem="todo" />
           </div>
        </ul>
    </div>
</template>

<style scoped>
.Add{
    margin-left: 5px;
}
.todoItem{
    background-color:lightgray;
    margin: 5px;
    border-radius: 10px;
    height: 40px;
    text-align: start;
}
.textItem{
    display: flex;
    height: 100%;
    width: 95%;
    align-items: center;
    justify-content: space-between;
}
</style>