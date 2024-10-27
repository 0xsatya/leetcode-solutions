use std::{char, collections::btree_map::IterMut, ops::Index};

fn first_word(s: &String) -> usize {
    let s1 = s.as_bytes();
    for (index, &item) in s1.iter().enumerate() {
        println!("item:{}", item);
        if item == b' ' {
            return index;
        }
    }
    return s.len();
}

pub fn main() {
    /*
    //Testing strings
    let mut s = String::from("foo");
    s.push_str("bar");
    println!("s is {}", s);

    let mut s1 = String::from("foo");
    let s2 = "bar";
    s1.push_str(&s2);
    println!("s1 is {}", s1);
    println!("s2 is {}", s2);

    let mut s = String::from("lo");
    s.push('l');
    println!("s is {}", s);

    let s1 = String::from("Hello, ");
    let s2 = String::from("world!");
    let s3 = s1 + &s2; // Note that s1 has been moved here and can no longer be used
                       // println!("s1 is {}", s1);
    println!("s3 is {}", s3);

    let s1 = String::from("tic");
    let s2 = String::from("tac");
    let s3 = String::from("toe");
    // format is same as println and doesn't take ownership
    let s = format!("{}-{}-{}", s1, s2, s3);

    
    let s1 = String::from("Hello World");
    let index = first_word(&s1);
    println!("Index should be 5: {}", index);
    println!("String was:{}", s1);
    */

    // string literal not dynamic, can't be modified.
    // it must be known at compile time
    let s = "hello world";

    // create String from string literal
    // let s = String::from("Hello World!");
    let s = String::from(s);
    print!("{s}");

    //************************************************ */
    //find first word of the string
    // let s1 = String::from("Hello World");
    // let index = first_word(&s1);
    
    //************************************************ */
    
    //************************************************ */


}
