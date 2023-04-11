function mul(a, b, callback) {
  setTimeout(() => {
    callback(a * b);
  }, 500);
}

function makeCal(...args) {
  let b = [];
  args.forEach((elm, index) => {
    if (index % 2 == 0) {
      if (index + 1 < args.length) b.push(elm * args[index + 1]);
      else b.push(elm);
    }
  });
  console.log(b);

  let sum = 0;

  b.forEach((elm) => {
    sum += elm;
  });

  console.log(sum);
}

makeCal(1, 2, 3, 4, 5);
