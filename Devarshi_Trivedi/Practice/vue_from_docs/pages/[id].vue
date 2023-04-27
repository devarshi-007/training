<template>
    <p>{{ route }}</p>
    <button @click="HeyChangeLayout">Click me</button>
    <NuxtLayout name="main"></NuxtLayout>
</template>

<script setup>
definePageMeta({
    validate: async (route) => {
    // Check if the id is made up of digits
    return /^\d+$/.test(route.params.id)
    },
    middleware: 'auth',
    layout: "main",
});

function HeyChangeLayout(){
    setPageLayout('marulayout');
}

const route = ref('orreo');
    // const route = useRoute();
    // console.log(route.params.id);
</script>
