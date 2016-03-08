"use strict";

function asyncFunction() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve('Fulfilled');
        }, 1000);
    });
}

// asyncFunction().then(value => {
//     console.log(value);
// }).catch(error => {
//     console.log(error);
// });

asyncFunction().then(value => {
    console.log(value);
}, error => {
    console.log(error);
})
