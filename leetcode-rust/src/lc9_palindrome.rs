// https://leetcode.com/problems/palindrome-number/
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Palindrome Number

// Given an integer x, return true if x is palindrome integer.
// An integer is a palindrome when it reads the same backward as forward.
// For example, 121 is a palindrome while 123 is not.

pub fn is_palindrome(x: i32) -> bool {
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false;
    }

    let (mut x, mut rev) = (x, 0);
    while x > rev {
        rev = rev * 10 + x % 10;
        x /= 10;
    }
    x == rev || x == rev / 10
}

//BEST SOLUTIONS
pub fn is_palindrome2(x: i32) -> bool {
    x.to_string() == x.to_string().chars().rev().collect::<String>()
}

pub fn main() {
    println!("{}", is_palindrome(121));
    println!("{}", is_palindrome2(121));
}
