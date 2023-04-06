<script setup>
// module import
import { ref, reactive, onMounted, onUpdated, onUnmounted } from 'vue'

// component import
import TodoComp from './components/Todo.vue';

// states
const titleClass = ref('my title')
const todoOpr = ref(null)
const toggleState = reactive({
  count: 0
})

// element reference
const mountIt = ref(null)

// props
const strtingId = ref(5)
const randomSmile = () => {
  let arr = ['😀', '😃', '😆', '😆', '😊', '😋', '😛', '🤣', '😂', '🤗']
  return arr[Math.floor(Math.random() * arr.length)]
}

// emits
const confirmId = ref(0)
const emitFunc = ref(() => { })

// mount
// onMounted(() => {
//   alert(`new value: "${mountIt.value.textContent ? 'show all' : 'remaining'}" mounted!`)
// })
const myp = ref(5)
onMounted(() => {
  console.log(`mounted ${Date.now}`)
})
onUpdated(() => {
  console.log(`updated ${Date.now}`)
})

onUnmounted(() => {
  console.log(`unmounted ${Date.now}`)
})

let g = setInterval(() => {
  myp.value++;
}, 1000)

let gg = setTimeout(() => {
  clearInterval(g)
  clearTimeout(gg)
}, 10000)



// functions
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

const callEmitFunc = () => {
  console.log(emitFunc.value(), "**********", typeof emitFunc, typeof emitFunc.value, emitFunc.value)
}

</script>

<template>
  <main>

    <p v-if="myp % 2 == 0">{{ myp }}</p>

    <!-- <input type="text" @input="onInput">
                    <button @click="changeCase" class="dynamicId">change case</button>
                    <button @click="todoOpr.addNew(titleClass)" class="dynamicId">Add todo</button>
                    <h1 :title="titleClass" ref="mountIt">{{ titleClass.valueOf() }}</h1>

                    <h1 v-if="smile()">Vue is awesome! 😀</h1>
                    <h1 v-else>Oh no 😶</h1>

                    <p>{{ confirmId }}</p>
                    <TodoComp ref="todoOpr" :id="strtingId" :randomSmile="randomSmile" @newID="(desidedID) => confirmId = desidedID"
                      @getNumber="(func) => emitFunc = func">* {{ emitFunc() }} *</TodoComp>

                    <button @click="callEmitFunc()" class="dynamicId">emit function call</button>
                  -->
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