let timeoutId = setTimeout(function () {
  alert("after a second");
  clearTimeout(timeoutId);
}, 2000);

console.log("hii");

console.log(window.location, window.location.href)

console.log(document.getElementsByClassName('row')) 