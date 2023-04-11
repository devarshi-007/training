const loading =  require('loading-cli');
const load = loading("loading...")

const users=[ "jinal-gajjar","fbfdbgrtbcbvx","nilam-singh","JayPonda","chintansakhiya","Ami-Kalola","aniketgohelimp","Ankit Jilka","Annavar-satish","Bhautik","Bhoomiz01","VatsaL","Rakshit Menpara","dhruvjoshi2000","abcd","Disha-Kothari","JAY PONDA IMPROWISED"];

const detail=[];

let notFound = 0;
let found = 0;
async function getData(){
    const timeStart=Date.now();

    for(let user of users){
        userDetail =await fetch(`https://api.github.com/users/${user}`,{
        
    }).then(x => x.json()).then(val=>{
        detail.push(val);
        if (val.message=="Not Found"){
            notFound++
        }
        else{
            found++
        }
    });
}
    const timeEnd = Date.now();
    load.stop()
    const countTime = (timeEnd-timeStart)/1000
    console.log("timing in (ms):"+countTime)
    console.log("Total Found:" + found )
    console.log("Not Found:" + notFound)
}

function Demo(){
    load.start()
    getData()
}
Demo()