#![allow(unused)]
// #![allow(warnings)]

// mod lc14_longest_common_prefix;
// mod sn3_lc1_two_sum;
// mod lc21_merge_two_sorted_lists;
// mod lc83_remove_duplicates_from_sorted_list;
// mod lc9_palindrome;
// mod tests;
mod nc1_array {
    pub mod sn1_lc217_contains_duplicates;
    pub mod sn2_lc242_valid_anagram;
}

mod tests;
use crate::tests::tests::main as tests_main;

fn main() {
    // Running test code
    println!("Running Rust main()");
    // tests::main();
    tests_main();

    // println!("Running leet code challenge!");
    // println!("\n --lc1 ------------");
    // lc1_two_sum::main();

    // println!("\n --lc9 ------------");
    // lc9_palindrome::main();

    // println!("\n --lc14 ------------");
    // lc14_longest_common_prefix::main();

    // println!("\n --lc21 ------------");
    // lc21_merge_two_sorted_lists::main();

    // println!("\n --lc83 ------------");
    // lc83_remove_duplicates_from_sorted_list::main();
}
