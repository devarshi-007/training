<script setup>
    import data from '../Data.js'
    const props = defineProps(['item']);
    const emits =defineEmits(['resetText'])
    let id=0;
    function TodoItemAdd(){
        let k = {}
        k.id=++id;;
        k.item=props.item;
        data.push(k);
        emits('resetText');
        console.log(data)
    }
</script>
<template>
    <button class="Add" @click="TodoItemAdd">Add</button>    
</template>