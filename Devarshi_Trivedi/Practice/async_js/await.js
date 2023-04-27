
const posts = [
    {title: "post1", body: "This is post 1" },
    {title: "post2", body: "This is post 2" },
];

function getPosts(){
    setTimeout(() => {
        let output='';
        posts.forEach((post,index) =>{
            output+=`<li>${post.title}</li>`; 
        });
        document.body.innerHTML = output;
    },1000);
}

function createPost(post, callback){
    return new Promise((resolve,reject)=>{
        setTimeout(()=>{
            posts.push(post);
            const error = false;

            if(!error){
                resolve();
            } else {
                reject("Something wrong");
            }
        },2000);
    })
}


// async function init(){
//     await createPost({title:'post 3',body:"this is post 3"})
//    getPosts();
// }

// init();


async function fetcher(){
    const res = await fetch('https://jsonplaceholder.typicode.com/posts');
    const data = await res.json();
    console.log(data)
}

fetcher();