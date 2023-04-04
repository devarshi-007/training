let timeoutId = setTimeout(function () {
  alert("after a second");
  clearTimeout(timeoutId);
}, 2000);

console.log("hii");

console.log(window.location, window.location.href);

console.log(document.getElementsByClassName("row"));

let xhttp = new XMLHttpRequest();

xhttp.onreadystatechange = function () {
  if (this.readyState == 4 && this.status == 200) {
    console.log(JSON.parse(this.responseText));
  }
};

xhttp.open("GET", "https://list.ly/api/v4/users/1", true);
xhttp.send();

let promiss = $.get("https://list.ly/api/v4/users/1");

promiss.then(
  (data) =>
    (document.getElementById("here").textContent =
      ("success: ", JSON.stringify(data))),
  (err) =>
    (document.getElementById("here").textContent =
      ("error: ", JSON.stringify(err)))
);

let form = document.getElementsByTagName("form")[0];
document.addEventListener("submit", (e) => {
  let field = form.elements["just"];
  console.log(
    "submitted",
    field.value,
    field.innerText,
    document.getElementById("here").innerText
  );
  if (field.value.trim() == "") {
    field.style.borderColor = "red";
    field.focus();
    e.preventDefault();
  }
});
