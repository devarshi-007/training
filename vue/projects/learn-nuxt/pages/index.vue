<script setup>
import Nav from './nav.vue'
const runtimeConfig = useRuntimeConfig()
const appConfig = useAppConfig()

</script>

<template>
    <Nav>
        <div>
            <h1>Welcome to the homepage</h1>
            <AppAlert>
                This is an auto-imported component
            </AppAlert>
        </div>
    </Nav>
</template>
  