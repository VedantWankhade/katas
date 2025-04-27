let arr = ["sunday", 1, 2, "wednesday"]
console.log(typeof arr)
console.log(arr);

arr[4] = "index 4"
console.log(arr)


arr[7] = "index 7" // index 5 and 6 will be empty - undefined
console.log(arr)

console.log(arr[5])
console.log(arr[6])


delete arr[1] // makes index 1 item as undefined
console.log(arr[1]);


// methods
console.log(arr.find(e => typeof e == 'number'))
console.log(arr.find(e => e == 'wednesday'))