# rust-playground

- Generate project `cargo new`
- Run `cargo run`
- Add dependency
    - Edit `Cargo.toml`
    - Run `cargo build`

## The book
- Using a `!` means that you're calling a macro instead of a normal function.
- compile
    - `rustc main.rs`
    - exectable `main` will be generated
- Cargo
    - build system and package manager
    - `cargo build`
        - `./target/debug/hello-cargo` will be generated
    - `cargo run` to compile then run the resulting exectable
    - `cargo check`
    - `cargo build --release` will create an exectable in `target/release` dir instead of `target/debug`
- guessing game
    - `use std:io` to bring the library into scope.
    - `let` to create a variable.
    - `let foo = 5` immutable
    - `let mut foo = 5` mutable
    - `::`
        - `String::new` -> new is the associated function of the String type.(static method)
    - `io::stdin()` returns an instance of `std::io::Stdin`
    - `.read_line(&mut guess)` to take whatever the user types into standard input and place that into a string.
        - the `&` indicaes that this argument is a reference.
    - `.expect("Failed to read line")` to handle potential failure

