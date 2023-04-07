<script setup>
import Nav from './nav.vue'
const runtimeConfig = useRuntimeConfig()
const appConfig = useAppConfig()
</script>

<template>
    <Head>
        <Title>hello</Title>
        <link rel="shortcut icon" type="image/ico" href="/public/favicon.ico" />
    </Head>
    <Nav> </Nav>
    <div>
        <h1>Welcome to the homepage</h1>
        <AppAlert>
            This is an auto-imported component
        </AppAlert>
        <img src="@/assets/nuxt.png" alt="Discover Nuxt 3" />
    </div>
</template>
  