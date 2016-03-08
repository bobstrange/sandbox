function onReady(fn) {
    const readyState = document.readyState;
    if (readyState === 'interactive' || readyState === 'complete') {
        fn();
    } else {
        window.addEventlistner('DOMContentLoaded', fn);
    }
}
onReady(() => {
    console.log('DOM fully loaded and parsed');
});

console.log('== STARTING ==');
