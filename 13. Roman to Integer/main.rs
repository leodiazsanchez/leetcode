fn main() {
    println!("{}", Solution::roman_to_int(String::from("III")));
}

struct Solution {}

impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let mut sum = 0;
        let mut prev = 0;

        for c in s.chars().rev() {
            let value = roman_value(c);

            if prev > value {
                sum -= value
            } else {
                sum += value
            }

            prev = value;
        }

        sum
    }
}

fn roman_value(c: char) -> i32 {
    match c {
        'I' => 1,
        'V' => 5,
        'X' => 10,
        'L' => 50,
        'C' => 100,
        'D' => 500,
        'M' => 1000,
        _ => 0,
    }
}
