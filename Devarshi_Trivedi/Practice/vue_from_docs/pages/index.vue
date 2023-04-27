
import { BaseMainModule } from '~/.nuxt/components';

<template>
    <!-- <p>{{ runtimeConfig }}</p>
    <p>{{ appConfig }}</p> -->
    <p>{{ msg }}</p>
    <ClientOnly>
        <BaseMainModule msg="Oyeoye" />
    </ClientOnly>
    <LazyBaseMainModule msg="show" v-if="s"/>
</template>

<script setup>
// const runtimeConfig = useRuntimeConfig();
// const appConfig = useAppConfig();


// appConfig.theme.dark=false;
const msg = ref('Hiiiiiiiii');
const s = ref(true);
</script>