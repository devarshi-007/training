let data =[];
function removeTodo(todo){    
    data = data.filter((t)=>t!==todo);
}
export default {data, removeTodo};
