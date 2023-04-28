//delay
async function getDelay() {
  const response = await fetch("https://httpbin.org/delay/3");
  const delay = await response.json();
  return delay;
}

//uuid
async function getUuid() {
  const response = await fetch("https://httpbin.org/uuid");
  const uuid = await response.json();
  return uuid.uuid;
}

//anything
async function anything(uuid) {
  const response = await fetch(`https://httpbin.org/anything`, {
    method: "POST",
    body: uuid,
  });
  const anything = await response.json();
  return anything;
}

async function main() {
  console.time("time");

  let uuid = await getUuid();
  getDelay();
  let anythingObj = await anything(uuid);
  console.log("data", anythingObj.data);
  console.timeEnd("time");
}

main();
