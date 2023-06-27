var toastElList = [].slice.call(document.querySelectorAll('.toast'))
var toastList = toastElList.map(function (toastEl) {
  return new bootstrap.Toast(toastEl)
})

toastList.forEach(element => {
    element.show()
});

var collapseElementList = [].slice.call(document.querySelectorAll('.collapse'))
var collapseList = collapseElementList.map(function (collapseEl) {
  return new bootstrap.Collapse(collapseEl)
})

document.querySelectorAll('[data-bs-toggle="popover"]')
.forEach(function(popover){
    new bootstrap.Popover(popover)
})