# @param {Integer[]} nums
# @return {Integer}
def single_number(nums)
  return nums.find{ |n| nums.count(n) == 1 }
end

p single_number([2,2,1])
p single_number([4,1,2,1,2])
p single_number([1])
