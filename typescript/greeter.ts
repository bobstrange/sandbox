function greeter(person: string) {
    return `Hello, $(person)`;
}

const user = "Jane User";
const error = [1, 2, 3];
document.body.innerHTML = greeter(user);
