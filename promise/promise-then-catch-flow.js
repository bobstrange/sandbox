function taskA() {
    console.log('TaskA');
}

function taskB() {
    console.log('TaskB');
}

function onRejected(error) {
    console.log(`Catch Error: A or B ${error}`);
}

function finalTask() {
    console.log('Final Task');
}

var promise = Promise.resolve();
promise
    .then(taskA)
    .then(taskB)
    .catch(onRejected)
    .then(finalTask);
