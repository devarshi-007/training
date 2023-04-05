<script setup>
import { ref, reactive, computed, onMounted, watch, onUpdated } from 'vue'
import TodoComp from './components/Todo.vue';

const titleClass = ref('my title')
const todoOpr = ref(null)
const toggleState = reactive({
  count: 0
})

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
}

function onInput(e) {
  titleClass.value = e.target.value;
  toggleState.count = 0
}

function smile() {
  return Math.round(Math.random() * 10) % 2 == 0
}

const mountIt = ref(null)
onMounted(() => {
  alert(`new value: "${mountIt.value.textContent ? 'show all' : 'remaining'}" mounted!`)
})

</script>

<template>
  <main>
    <input type="text" @input="onInput">
    <button @click="changeCase" class="dynamicId">change case</button>
    <button @click="todoOpr.addNew(titleClass)" class="dynamicId">Add todo</button>
    <h1 ref="mountIt">{{ titleClass.valueOf() }}</h1>
    <!-- <div v-bind:id="dynamicId"></div> -->

    <h1 v-if="smile()">Vue is awesome! 😀</h1>
    <h1 v-else>Oh no 😶</h1>

    <TodoComp ref="todoOpr" />
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