# @param {Integer[]} nums
# @param {Integer} k
# @return {Boolean}
def contains_nearby_duplicate(nums, k)
  ans = false
  nums.each_with_index do |n, i|
    ans = nums[i+1...i+k].include?(n)
  end

  ans
end
