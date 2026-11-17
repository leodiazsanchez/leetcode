fn main() {
    println!("{}", Solution::search_insert([1, 3, 5, 6].to_vec(), 5));
}

struct Solution {}

impl Solution {
    pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
        let mut left = 0;
        let mut right = nums.len();

        while left < right {
            let mid = left + (right - left) / 2;
            if nums[mid] < target {
                left = mid + 1
            } else if nums[mid] >= target {
                right = mid
            } else {
                return mid.try_into().unwrap();
            }
        }

        return left.try_into().unwrap();
    }
}
