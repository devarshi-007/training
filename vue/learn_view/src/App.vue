<script setup>
import { ref, reactive } from 'vue'

const titleClass = ref('my title')
const toggle = reactive({
  count: 0
})
const s = setTimeout(() => {
  titleClass.value = 'new title';
  clearTimeout(s)
}, 0)

function changeCase() {
  if (toggle.count == 0) {
    titleClass.value = titleClass.value.toUpperCase()
    toggle.count = 1
  } else if (toggle.count == 1) {
    titleClass.value = titleClass.value.toLowerCase()
    toggle.count = 2
  } else {
    let str = titleClass.value.toLowerCase().split(' ');
    for (var i = 0; i < str.length; i++) {
      str[i] = str[i].charAt(0).toUpperCase() + str[i].slice(1);
    }

    titleClass.value = str.join(' ');
    toggle.count = 0;
  }
}

function onInput(e) {
  titleClass.value = e.target.value;
  toggle.count = 0
}

</script>

<template>
  <main>
    <input type="text" @input="onInput">
    <button @click="changeCase" class="dynamicId">change case</button>
    <h1>{{ titleClass.valueOf() }}</h1>
    <!-- <div v-bind:id="dynamicId"></div> -->
  </main>
</template>

<style scoped>
.dynamicId {
  margin: 5px;
}
</style>