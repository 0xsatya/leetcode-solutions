// PROBLEM - 1
// Two Sum
// https://leetcode.com/problems/two-sum/
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

use std::collections::HashMap;

// sum using hash function [ complexity O(n)]
pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut m: HashMap<i32, i32> = HashMap::new();
    for (i, v) in nums.iter().enumerate() {
        match m.get(&(target - *v)) {
            Some(&i2) => return vec![i as i32, i2],
            None => m.insert(*v, i as i32),
        };
    }
    vec![]
}

pub fn main1() {
    println!("Two Sum called !");
    let arg1: Vec<i32> = [3, 2, 5, 6].to_vec();
    let list1: Vec<i32> = two_sum(arg1, 8);

    if list1.len() != 0 {
        println!("SUCCESS");
        println!("{:?}", list1);
    } else {
        println!("FAILED");
    }
}
