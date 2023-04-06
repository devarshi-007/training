<template>
  <button v-on:[attr]="clik">Show</button>
  <div v-if="disable">
    <span>Message is: {{ disable.value ? 'This is not be seen': 'This will be seen' }}</span>
    <span v-html="rawhtml"></span>
  </div>
  <button @click="increment">count: {{ state.count }}, {{ count }}</button>
  <span v-show="true"> Value is: {{ value }},{{ dateN}}</span>
  <div :class="{'active':isActive, 'error':hasError}">dfssdfdsfds</div>
  <div>
  <ul>
    <li v-for="value in myObject" :key="value.title">
      {{ value }}
    </li>
  </ul>
  </div>
  <input @key.left="keyCouer.value++"/>
  <span>{{ keyCouer }}</span>
  <hr />

</template>

<script setup>
import {reactive, ref, computed} from 'vue'
const myObject = reactive({
  title: 'How to do lists in Vue',
  author: 'Jane Doe',
  publishedAt: '2016-04-10'
})
let attr = 'click'
const disable=ref(true);
const isActive=ref(true);
const hasError=ref(false);
let a=ref(1),b=ref(2);
const value = computed(()=>{
  return a.value+b.value;
})
const rawhtml = "<br /><span style=\"color:red;\">This should be red</span>";
a.value=10;
function clik(){
  disable.value=!disable.value;
}
let keyCouer = ref(0);
const state = reactive({ count: 10 })

function increment() {
  state.count++
}
//const dateN = ref(Date.now())
//const date = computed(()=>dateN);

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.active{
  color: green;
}
.error{
  color: gray;
}
</style>
