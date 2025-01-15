# @param {Integer[]} nums
# @param {Integer} k
# @return {Boolean}
def contains_nearby_duplicate(nums, k)
  ans = false
  nums.each_with_index do |n, i|
    k.times do |kk|
      ans = true if n == nums[i + kk]
    end
  end

  ans
end
