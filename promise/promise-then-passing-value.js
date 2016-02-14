function doubleUp(value) {
    return value * 2;
}

function increment(value) {
    return value + 1; // returnした値が引数に設定される
}

function output(value) {
    console.log(value);
}

const promise = Promise.resolve(1);
promise
    .then(increment)
    .then(doubleUp)
    .then(output)
    .catch(error => {
        console.error(error);
    });
