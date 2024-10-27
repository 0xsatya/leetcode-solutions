## Rust Notes

### Important concepts

- When variable goes out of scope, `drop` function is called.
- Rust calls drop automatically at the closing curly bracket.

`String`

`   let s1 = String::from("hello");`
`let s2 = s1;`

- It is stored in heap, can be mutable.
  ![alt text](https://doc.rust-lang.org/book/img/trpl04-01.svg)
- The **_length_** is how much memory, in bytes, the contents of the String are currently using.
- The **_capacity_** is the total amount of memory, in bytes, that the String has received from the allocator.
- when s2 and s1 go out of scope, they will both try to free the same memory. This is known as a double free error and is one of the memory safety bugs we mentioned previously. Freeing memory twice can lead to memory corruption, which can potentially lead to security vulnerabilities.

`"string".to_owned()`

- .clone() as "I already have a T, and I'd like another one" whereas
- .to_owned() as "I have something borrowed, and I want to make an owned copy so I can mess with it".
- let x: T = ...; let y: T = x.clone();
- let x: &T = ...; let y: T = x.to_owned();.
