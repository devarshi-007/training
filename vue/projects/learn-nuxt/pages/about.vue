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
        <Title>{{ p }}</Title>
        <Meta name="description" :content="p" />
        <Link rel="shortcut icon" type="image/jpg" href="/img/logo.jpg" />
    </Head>
    <section>
        <p>This page will be displayed at the /about route.</p>
        <p>{{ p }}</p>
        <img src="/img/logo.jpg" />
    </section>
</template>
  