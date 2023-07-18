const println = console.log.bind(this)

// strings

let str = 'this' + " is" + 
"\n `something" + `
 new\`
`
let str1 = "this is\
\n `something\
\n new`\
\n"

println(str, typeof str, str == str1, str === str1)

// numeric

let p = 10
let k
println(p++, p)
println(++p, p)

println(p--, p)
println(--p, p)

println(3 + 3 * 3)

k = 10.3
println(k+p)

println(5.0 == 5, "5.0" === 5)
println(5.0 == 5, "5.0" !== 5)

if (0 || undefined || null || NaN || false) {
 println(0)
}
println(0 ,"as false")

if (+(1.1 + 10.333).toFixed(2) === 11.43){
    println("ok")
}

if(undefined || "false"){
    println("true")
}

function name5(newname = null){
    
    return String(newname).toUpperCase()
}
let myFun = name5

println(myFun(1), name5(10))

let symbol = Symbol()

let obj = {
    name: "k",
    "full name": "j k",
    [symbol]: "hii"
}

println(obj, obj["full name"], obj.name)
println(obj, obj["full name"], obj.symbol)

obj["full name"]

// document.getElementById("main").setAttribute("data-new", "new-data")


function  makeObj(name = "", surname = "") {
    if ((name && name.trim() != "")  && (surname && surname.trim() != ""))
    return {
        name: name,
        surname: surname,
        getFullName: function(){
            return `${this.name} ${this.surname}`
        }
    };

    throw Error("name and surname should be string")
}

obj = makeObj("jay", "ponda")
println(obj.getFullName(), obj.name , !obj, !!obj)

println(eval("1 + 2 * 3 ** 2"))

const arr = [1, 2]

println(typeof arr, Array.isArray(arr), arr.length, arr.values().next(), arr.values(), arr.values().next())