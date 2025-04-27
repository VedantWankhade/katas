let o = {
    a: 2,
    "my var": "hello"
}

console.log(typeof o)
console.log(o);
console.log(o.a);
console.log(o["my var"]);
delete o.a
console.log(o);
delete o['my var']
console.log(o);



// references
let x = {
    a: 2,
    b: false
}
let y = x; // copies the reference - does not clone the object
console.log(x)
console.log(y)

y.b = true;
console.log(x)
console.log(y);

// deep copy
let a = {
    a: 22,
    b: false
}

let b = Object.assign({}, a)
console.log("a", a)
console.log("b", b);
b.b = true;
console.log("a", a)
console.log("b", b);

let c = { ...a }
console.log("a", a)
console.log("c", c);
c.b = true;
console.log("a", a)
console.log("c", c);