function add(a, b, callback) {
  callback(a + b);
}

function mul(a, b, callback) {
  callback(a * b);
}

function opr(sum = []) {}

opr(2, 1);
