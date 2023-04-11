import { print } from "../shared/common.js";

let promise = new Promise((res, err) => setTimeout(err, 100, "hi"));

export default function () {
  print(promise);
  promise
    .then((val) => print("ok: ", val))
    .catch((err) => {
        print("catch: ", err)
        return err + " handled"
    })
    .then((val) => print("ok 1: ", val))
    .catch((err) => print("catch 1: ", err));

  let j = 0;
  print(1);
    let k = setInterval(() => {
      if (j == 2) {
        clearInterval(k);
      }
      j++;
      print(5);
    }, 2000);
    print(10);
}

const p1 = new Promise((res, rej) => {
  let k = setTimeout(() => {
    clearInterval(k);
    res(true);
  }, 5000);
});

const p2 = new Promise((res, rej) => {
  let k = setTimeout(() => {
    clearInterval(k);
    res(false);
  }, 5000);
});

async function nom() {
  let k = Date.now();

  Promise.all([p1, p2])
    .then((result) =>
      result.every((val) => console.log(Math.abs(Date.now() - k), val, "ok***"))
    )
    .catch((err) => console.log(Math.abs(Date.now() - k), err, "err***"));

}

nom();
console.log(typeof null == typeof undefined);
