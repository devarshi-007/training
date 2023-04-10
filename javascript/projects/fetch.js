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

async function request(users) {
  const fetchArr = users.map((user) =>
    fetch(`https://api.github.com/users/${user}`, {
      headers: {
        authorization: "token ~~~",
      },
    })
  );

  const resArray = await Promise.allSettled(fetchArr).then((res) => res.json());

  return resArray;
}

let state = "loading";
let len = state.length;
function loading() {
  return setInterval(() => {
    console.log(state.padStart(len++, state));
  }, 1000);
}

function getResults(resArr) {
  const result = {
    totalUsers: 0,
    userAvalable: 0,
    resList: [],
    takenTime: null,
  };

  return result;
}

// // loading start
// const loadingId = loading();

// request array
const resArr = request(userArr);
console.log(resArr, "here");
// // calculate result
// const result = request(resArr);

// // loading end
// clearInterval(loadingId);

// // printing result
// console.log(result);
