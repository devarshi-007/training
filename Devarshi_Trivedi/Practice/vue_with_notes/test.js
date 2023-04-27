

async function add(){
    let a = new Promise(resolve =>setTimeout(() => resolve(5),5000))
   
    let b = new Promise(resolve =>setTimeout(() => resolve(10),5000))

    let k = Date.now()

    let a1 =   a;
    let b1 =  b;
    console.log("Tom and Jerry");
    console.log(await a1,await b1,(Date.now()-k)/1000);

    // Promise.all([a,b]).then(
    //     _ => console.log((Date.now()-k)/1000)
    // )
    // if (a1 !== 4){
        // console.log((Date.now()-k)/1000);
    // } 
}

add();