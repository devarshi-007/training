
async function sum() {
    let D = Promise.all([geta(),getb()]).then(
       D => console.log(D)
    );
}

async function geta() {
    const a = await aget();
    return a;
}

async function getb() {
    const b = await bget();
    return b;
}

function aget(){
    return new Promise(resolve=>setTimeout(resolve, 2000, 1));
}

function bget(){
    return new Promise(resolve => setTimeout(resolve, 5000, 2));
}
sum();