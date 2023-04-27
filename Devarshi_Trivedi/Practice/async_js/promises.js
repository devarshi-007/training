//promises

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

// createPost({title:'post 3',body:"this is post 3"})
// .then(getPosts)
// .catch(err => console.log(err));

const promise1 = Promise.resolve('Hello world!');
const promise2 = 10;
const promise3 = new Promise((resolve,reject)=>setTimeout(resolve, 2000, 'GoodBye'));
const promise4 = fetch('https://jsonplaceholder.typicode.com/posts').then(res=>res.json());

Promise.all([promise1, promise2, promise3,promise4]).then((values) => console.log(values));