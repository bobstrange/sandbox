const req = require('request');

function getURL(URL) {
    return new Promise((resolve, reject) => {
        req(URL, (error, response, body) => {
            if (!error && response.statusCode === 200) {
                resolve(body);
            } else {
                reject(error);
            }
        });
    });
}

const URL = 'http://httpbin.org/get';
getURL(URL).then(value => {
    console.log(value);
}).catch(error => {
    console.log(error)
});
