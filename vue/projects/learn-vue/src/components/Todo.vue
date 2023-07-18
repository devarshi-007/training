<script setup>
import { ref, computed, watch } from 'vue'

const props = defineProps({
  id: Number,
  randomSmile: Function
})

const emit = defineEmits(
  ['newID', 'getNumber']
)

let id = props.id ? props.id : 0;
emit('newID', id)
emit('getNumber', () => Math.ceil(Math.random() * 10))

const showAll = ref(true)
const todos = ref([
  { id: id++, text: 'Learn HTML', done: true },
  { id: id++, text: 'Learn JavaScript', done: false },
  { id: id++, text: 'Learn Vue', done: false }
])

const filteredTodos = computed(() => {
  return showAll.value ? todos.value :
    todos.value.filter((val) => !val.done)
})

function addNew(text) {
  todos.value.push(
    { id: id++, text: text, done: false }
  )
}

function removeTodo(todo) {
  todos.value = todos.value.filter((val) => val != todo)
}

function done(todo) {
  todo.done = !todo.done
}

watch(todos, () => {
  let completed = 0

  todos.value.forEach((elm) => {
    if (elm.done) completed++;
  })

  if (completed >= 0) {
    alert(`you complated ${completed} todos`)
  }
}, { deep: true })

defineExpose({
  addNew
});

</script>

<template>
  <label>
    <input type="checkbox" v-model="showAll" ref="mountIt" />
    <span :class="{ done: showAll, remaining: !showAll }" class="dynamicId">{{ showAll ? "show all" : "show remaining"
    }}</span>
  </label>

  <ol>
    <li v-for="todo in filteredTodos" :key="todo.id">
      <span>{{ randomSmile() }}</span>
      <span>id: {{ todo.id }},</span>
      <span :class="{ done: todo.done, remaining: !todo.done }" @click="done(todo)">text: {{ todo.text }},</span>
      <button @click="removeTodo(todo)">X</button>
    </li>
  </ol>
  <slot>* fallback content *</slot>
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