async function getData() {

    fetch('https://api.github.com/users')
    .then ((response)=>{
        return response.json();
    })
    .then((response)=>{
       const res = console.log(response)
       if(!Object.keys(res.data).length){
        console.log("no data found");
        }
    });
}
//console.log(JSON.stringify(response));
//console.log(response)
// console.log(response.status)
// console.log(response.statusText); 

// if (response.status === 200) {
//     let data = response.text();
// }
// .then(
//     console.log(response)
// )
// }
getData();
