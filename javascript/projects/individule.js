import fetch from "node-fetch";

const userArr = [
  "devarshi-007",
  "devar-007",
  "chintansakhiya",
  "chintansakiya",
  "Ami-Kalola",
  "aniketgohelimp",
  "Ankit Jilka",
  "Annavar-satish",
  "Bhautik",
  "Bhoomiz01",
  "VatsaL",
  "Rakshit Menpara",
  "dhruvjoshi2000",
  "abcd",
  "Disha-Kothari",
  "Denish Makadiya",
  "JAY PONDA IMPROWISED",
];

let arr = [];
let counter = 0;
async function getUsers(user) {
  const response = await fetch(`https://api.github.com/users/${user}`, {
    headers: {
      authorization: "token ~~~",
    },
  });
  const name = await response.json();
  arr.push(name.login);
  console.log(arr);
  console.log(counter++);
}

userArr.forEach((elm) => {
  getUsers(elm);
});

// console.log(hello);
