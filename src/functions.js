function speak(msg) {
    console.log(msg);
}

speak("hello")
speak()

function speak1(msg = "hello") {
    console.log(msg);
}

speak1("hello")
speak1()

// even if the definition does not expect parameters, you still have access to them through arguments object
function speak2() {
    console.log(arguments)
    console.log(arguments[0])
}

speak2("yo what", "really?")