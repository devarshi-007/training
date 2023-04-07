<script setup lang="ts">
import { ref } from 'vue';
const p = ref("hello from about")
useHead({
    title: p,
    meta: [
        { name: 'description', content: 'My amazing site.' }
    ],
    bodyAttrs: {
        class: 'test'
    },
    script: [{ innerHTML: 'console.log(\'Hello world\')' }],
})

</script>

<template>
    <Head>
        <title>{{ p }}</title>
        <link rel="shortcut icon" type="image/png" href="/img/nuxt.png" />
    </Head>
    <section>
        <p>This page will be displayed at the /about route.</p>
        <p>{{ p }}</p>
    </section>
</template>
  