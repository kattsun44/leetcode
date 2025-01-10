# @param {Integer[]} nums
# @return {Integer}
def find_max_consecutive_ones(nums)
  res, count = 0, 0
  nums.each do |n|
    count = n == 1 ? count + 1 : 0
    res = [res, count].max
  end
  return res
end

p find_max_consecutive_ones([0])
p find_max_consecutive_ones([1])
p find_max_consecutive_ones([1, 0])
p find_max_consecutive_ones([1,1,0,1,1,1])
p find_max_consecutive_ones([1,0,1,1,0,1])
