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

let state = "loading";
let len = state.length;
const result = {
  totalUsers: 0,
  userAvalable: 0,
  takenTime: null,
  resultArr: null,
};

function loading() {
  return setInterval(() => {
    console.log(state.padEnd(len++, "*"));
  }, 1000);
}

async function request(users, loadingID, callback) {
  const start = Date.now();

  const fetchArr = users.map((user) =>
    fetch(`https://api.github.com/users/${user}`, {
      headers: {
        authorization: "token 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",
      },
    })
  );

  await Promise.all(fetchArr)
    .then((res) => Promise.all(res.map((elm) => elm.json())))
    .then((res) => callback(res, start, loadingID));
}

function getResults(resArr, starttime, loadingID) {
  clearInterval(loadingID);
  result.takenTime = Date.now() - starttime;
  result.totalUsers = resArr.length;
  result.resultArr = resArr.map((element) => {
    if (element.hasOwnProperty("login")) {
      result.userAvalable++;
      return element.login;
    }
    return {};
  });

  result.resultArr.forEach((element) => {
    console.log(element);
  });
  console.log(result.takenTime, result.userAvalable, result.totalUsers);
}

// loading start
const loadingId = loading();

// request array
request(userArr, loadingId, getResults);
