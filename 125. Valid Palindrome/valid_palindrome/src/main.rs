fn main() {
    println!("{}", Solution::is_palindrome("A man, a plan, a canal: Panama".to_string()));
}

struct Solution {}

impl Solution {
    pub fn is_palindrome(s: String) -> bool {

        let chars: Vec<char> = s
            .chars()
            .filter(|c| c.is_ascii_alphanumeric())
            .map(|c| c.to_ascii_lowercase())
            .collect();

        let (mut left, mut right) = (0, chars.len());

        while left < right {
            right -= 1;

            if chars[left] != chars[right] {
                return false;
            }

            left += 1;
   
        }

        true
    }
}