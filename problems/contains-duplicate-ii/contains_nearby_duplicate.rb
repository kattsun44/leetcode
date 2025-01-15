# @param {Integer[]} nums
# @param {Integer} k
# @return {Boolean}
def contains_nearby_duplicate(nums, k)
  h = {}
  nums.each_with_index do |n, i|
    return true if h[n] != nil && (h[n] - i).abs <= k
    h[n] = i
  end

  false
end
