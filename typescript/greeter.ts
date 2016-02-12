function greeter(person) {
    return `Hello, $(person)`;
}

const user = "Jane User";

document.body.innerHTML = greeter(user);