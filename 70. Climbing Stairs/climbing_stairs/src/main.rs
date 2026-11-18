use std::collections::HashMap;

fn main() {
    println!("{}", Solution::climb_stairs(45));
}

struct Solution {}


impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {

        let mut cache: HashMap<i32, i32> = HashMap::new();

        Self::helper(n, &mut cache)
    }

    fn helper(n: i32, cache: &mut HashMap<i32, i32>) -> i32 {
        if n <= 2 {
            return n;
        }

        if let Some(&ans) = cache.get(&n) {
            return ans;
        }

        let ans = Self::helper(n - 1, cache)
                + Self::helper(n - 2, cache);

        cache.insert(n, ans);

        ans
    }
}