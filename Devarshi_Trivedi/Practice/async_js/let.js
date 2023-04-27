//promise

async function tio(){
    return new Promise((resolve,reject)=>{
        setTimeout(resolve, 2000, "I am tio.");
    });
}

async function toyer(){
    return new Promise((resolve, reject)=>{
        setTimeout(resolve, 2000, "I am tior.");
    });
}

async function caller(){
    let p1 = tio();
    let p2 = toyer();
    Promise.all([p1,p2]).then(
        val=>console.log(val),
    )
}

caller();


