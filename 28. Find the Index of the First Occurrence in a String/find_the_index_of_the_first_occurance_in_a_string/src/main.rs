fn main() {
    println!("{}", Solution::str_str("leetcode".to_string(), "leeto".to_string()));
}

struct Solution {}


impl Solution {
    pub fn str_str(haystack: String, needle: String) -> i32 {
        match haystack.find(&needle) {
            Some(x) => x as i32,
            None => -1,
        }
    }
}