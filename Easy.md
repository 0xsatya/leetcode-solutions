## Easy problems of leetcode with solutions in JS, Typescript & Go






### Problem - Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Javascript
```
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let chash = {}
    for(i = 0; i<nums.length; i++){
        let comp = target-nums[i]
        if(chash[comp] >= 0) return [i, chash[comp]]
        chash[nums[i]] = i;
    }
};
```

RUST
```

use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        use std::collections::HashMap;
        
        let mut m: HashMap<i32, i32> = HashMap::new();
        for (i, v) in nums.iter().enumerate() {
            match m.get(&(target - *v)) {
                Some(&i2) => return vec![i as i32, i2],
                None => m.insert(*v, i as i32),
            };
        }
        vec![]
    }
}
```
### Problem - Palindrome Number

Given an integer x, return true if x is palindrome integer.
An integer is a palindrome when it reads the same backward as forward.
For example, 121 is a palindrome while 123 is not.

Javascript
```
/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    if(x < 0 ) return false;
    
    let palendInt = 0;
    let temp = x;
    while( temp > 0) {
        palendInt = palendInt*10 + temp % 10;
        temp = Math.floor(temp/10);
    }
    if(palendInt === x) return true;
    else return false;  
};
```

BEST SOLUTION
```
function isPalindrome(x) {
  if (x < 0) return false;
  if (x < 10) return true;
  if (x % 10 === 0) return false;

  let rev = 0;
  while (rev < x) {
    rev *= 10;
    rev += x%10;
    x = Math.trunc(x/10);
  }
  return rev === x || Math.trunc(rev/10) === x;
}
```
BY CONVERTING TO STRING
```
var isPalindrome = function(x) {
    if (x < 0) return false;
	
    // reverse the string representation of x
    const reverse = `${x}`.split('').reverse().join('');
    // compare the value regardless of types
    return x == reverse;
};
```

RUST
```
//MY SOLUTION
impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x < 0 || (x % 10 == 0 && x != 0) {
            return false
        }
        
        let (mut x, mut rev) = (x, 0);
        while x > rev {
            rev = rev * 10 + x % 10;
            x /= 10;
        }
        x == rev || x == rev / 10
    }
}
```

BEST SOULUTIONS
```
impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        x.to_string()==x.to_string().chars().rev().collect::<String>()
    }
}

```