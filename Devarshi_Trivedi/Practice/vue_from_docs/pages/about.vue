<template>
    <p>{{ runtimeConfig }}</p>
    <p>{{ appConfig }}</p>
</template>

<script setup>
const runtimeConfig = useRuntimeConfig();
const appConfig = useAppConfig();
appConfig.theme.dark=false;
</script>