export default defineNuxtRouteMiddleware((to, from)=>{
    const route = useRoute()
    if(route.params.id <= 3){
        console.log('abbe saale');
        return navigateTo('/')
    }
})