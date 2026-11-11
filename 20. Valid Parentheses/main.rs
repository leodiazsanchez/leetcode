use std::collections::HashMap;

fn main() {
    println!("{}", Solution::is_valid(String::from("()")));
}

struct Solution {}

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let pairs = HashMap::from([('(', ')'), ('[', ']'), ('{', '}')]);

        let mut stack = Vec::new();

        for c in s.chars() {
            if let Some(&closing) = pairs.get(&c) {
                stack.push(closing);
            } else if stack.pop() != Some(c) {
                return false;
            }
        }

        stack.is_empty()
    }
}
