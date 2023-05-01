//parallel run excercise
//for get uuid
async function getUuid(){
    console.time("getUuid Time")
    const response = await fetch("https://httpbin.org/uuid")
    console.timeEnd("getUuid Time")
    const uuid = await response.json();
    return uuid.uuid;
}

// getdelay 
async function delay(delayData){
    console.time("Delay Time");
    const response = await fetch("https://httpbin.org/delay/5")
    console.timeEnd("Delay Time");
    const delay = await response.json();
    return delay;
}

//post 
async function postAnything(uuid){
    const response = await fetch(`https://httpbin.org/anything`, {
        method: "POST",
        body: uuid,
      });
    const data = await response.json()
    return data
}

async function main() {
    console.time("TotalTime:")
    
    let uuid = getUuid();
    let delayRes = delay();
    console.time("Post time:")
    let anything = await postAnything(await uuid);
    console.log("data", anything.data);
    console.timeEnd("Post time:")

    await delayRes
    console.timeEnd("TotalTime:")
  }
  main();

