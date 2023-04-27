export default defineNuxtRouteMiddleware((to,from)=>{
    const route = useRoute();
    if(route.params.top_id.length > 0){
        console.log("middleware 1");
    }
    else{
        return navigateTo('/');
    }
})