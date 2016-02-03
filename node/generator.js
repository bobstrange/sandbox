(function() {
    "use strict";

    function* makeNumber() {
        yield 1;
        yield 2;
        yield 3;
    }

    let generator = makeNumber();

    console.log(generator.next());
    console.log(generator.next());
    console.log(generator.next());
    console.log(generator.next());
})
