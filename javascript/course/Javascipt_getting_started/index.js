const print = console.log.bind(this)

// strings

let str = 'this' + " is" + 
"\n `something" + `
 new\`
`
let str1 = "this is\
\n `something\
\n new`\
\n"

print(str, typeof str, str == str1, str === str1)

// numeric

let p = 10
let k
print(p++, p)
print(++p, p)

print(p--, p)
print(--p, p)

print(3 + 3 * 3)

k = 10.3
print(k+p)