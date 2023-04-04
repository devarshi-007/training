<script setup>
import { ref, reactive, computed } from 'vue'

let id = 0;
const titleClass = ref('my title')
const toggleState = reactive({
  count: 0
})
const showAll = ref(true)
const todos = ref([
  { id: id++, text: 'Learn HTML', done: true },
  { id: id++, text: 'Learn JavaScript', done: false },
  { id: id++, text: 'Learn Vue', done: false }
])
const filteredTodos = computed(() => {
  return showAll ? todos.value :
    todos.value.filter((val) => !val.done)
})

console.log(filteredTodos)
// const s = setTimeout(() => {
//   titleClass.value = 'new title';
//   clearTimeout(s)
// }, 0)
function changeCase() {
  if (toggleState.count == 0) {
    titleClass.value = titleClass.value.toUpperCase()
    toggleState.count = 1
  } else if (toggleState.count == 1) {
    titleClass.value = titleClass.value.toLowerCase()
    toggleState.count = 2
  } else {
    let str = titleClass.value.toLowerCase().split(' ');
    for (var i = 0; i < str.length; i++) {
      str[i] = str[i].charAt(0).toUpperCase() + str[i].slice(1);
    }

    titleClass.value = str.join(' ');
    toggleState.count = 0;
  }
  toggle.value = smile()
}

function onInput(e) {
  titleClass.value = e.target.value;
  toggleState.count = 0
}

function smile() {
  return Math.round(Math.random() * 10) % 2 == 0
}

function addNew() {
  todos.value.push(
    { id: id++, text: titleClass.value, done: false }
  )
}

function removeTodo(todo) {
  todos.value = todos.value.filter((val) => val != todo)
}

function done(todo) {
  todo.done = !todo.done
  console.log(todos.value, filteredTodos.value, showAll.value)
}


</script>

<template>
  <main>
    <input type="text" @input="onInput">
    <button @click="changeCase" class="dynamicId">change case</button>
    <button @click="addNew" class="dynamicId">Add todo</button>
    <h1>{{ titleClass.valueOf() }}</h1>
    <!-- <div v-bind:id="dynamicId"></div> -->

    <h1 v-if="smile()">Vue is awesome! 😀</h1>
    <h1 v-else>Oh no 😶</h1>

    <label>
      <input type="checkbox" v-model="showAll" />
      <span :class="{ done: showAll, remaining: !showAll }" class="dynamicId">{{ showAll ? "show all" : "show remaining"
      }}</span>
    </label>

    <ol>
      <li v-for="todo in filteredTodos" :key="todo.id">
        <span>id: {{ todo.id }},</span>
        <span :class="{ done: todo.done, remaining: !todo.done }" @click="done(todo)">text: {{ todo.text }},</span>
        <button @click="removeTodo(todo)">X</button>
      </li>
    </ol>

  </main>
</template>

<style scoped>
.done {
  color: green;
  text-decoration: line-through;
}

.remaining {
  color: blue;
}

.dynamicId {
  margin-left: 5px;
}

li span {
  margin: 5px;
  user-select: none;
}
</style>