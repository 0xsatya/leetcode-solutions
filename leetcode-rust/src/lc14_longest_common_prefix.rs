// PROBLEM NO - 14. Longest Common Prefix
// NAME -

// Description:
// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

//MY SOLUTION
pub fn longest_common_prefix1(strs: Vec<String>) -> String {
    let mut lcp = strs[0].clone();
    let i = 1;
    for i in 1..strs.len() {
        let next = &strs[i];
        let mut j = 0;
        let mut temp_lcp = String::from("");
        while j < lcp.len() && j < next.len() {
            if lcp.chars().nth(j) == next.chars().nth(j) {
                match lcp.chars().nth(j) {
                    Some(char_val) => temp_lcp.push(char_val),
                    None => break,
                }
            } else {
                break;
            };
            j += 1;
        }
        lcp = temp_lcp;
    }
    lcp.to_string()
}

//BEST SOULUTIONS
pub fn longest_common_prefix2(strs: Vec<String>) -> String {
    fn lcp_inplace(mut s1: String, s2: &str) -> String {
        let mut i = 0;
        for (c1, c2) in s1.chars().zip(s2.chars()) {
            if c1 != c2 {
                break;
            }
            i += 1;
        }
        s1.truncate(i);
        s1
    }
    if strs.len() > 0 {
        strs.iter()
            .skip(1)
            .fold(strs[0].clone(), |acc, x| lcp_inplace(acc, &x))
    } else {
        String::from("")
    }
}
//--------
pub fn longest_common_prefix3(strs: Vec<String>) -> String {
    match strs.is_empty() {
        true => "".to_string(),
        _ => strs.iter().skip(1).fold(strs[0].clone(), |acc, x| {
            acc.chars()
                .zip(x.chars())
                .take_while(|(x, y)| x == y)
                .map(|(x, _)| x)
                .collect()
        }),
    }
}

//--------
pub fn longest_common_prefix4(strs: Vec<String>) -> String {
    let first = &strs[0];
    let mut max_ind = first.len();

    for next in strs.iter().skip(1) {
        if let Some(((last_ind, _), _)) = first[..max_ind]
            .char_indices()
            .zip(next.chars())
            .take_while(|((_, c1), c2)| c1 == c2)
            .last()
        {
            max_ind = last_ind + 1;
        } else {
            return "".to_owned();
        }
    }
    first[..max_ind].to_owned()
}
//----

pub fn longest_common_prefix5(strs: Vec<String>) -> String {
    if strs.is_empty() {
        return String::from("");
    }

    let mut answer = strs[0].clone();
    for i in 1..strs.len() {
        while !strs[i].starts_with(&answer) {
            answer = answer.chars().take(answer.len() - 1).collect();
            if answer.len() == 0 {
                return String::from("");
            }
        }
    }
    return answer;
}
//------
pub fn longest_common_prefix6(strs: Vec<String>) -> String {
    strs.into_iter()
        .reduce(|prev, current| {
            prev.chars()
                .zip(current.chars())
                .take_while(|(ch1, ch2)| ch1 == ch2)
                .map(|pair| pair.0)
                .collect::<String>()
        })
        .unwrap()
}

pub fn main() {
    println! {"longest common prefix called"};
}
