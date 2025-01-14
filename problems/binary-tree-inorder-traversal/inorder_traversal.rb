# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val = 0, left = nil, right = nil)
#         @val = val
#         @left = left
#         @right = right
#     end
# end

# @param {TreeNode} root
# @return {Integer[]}
def inorder_traversal(root)
  ans = []
  def traversal(root, ans)
    return if root.nil?
    traversal(root.left, ans)
    ans << root.val
    traversal(root.right, ans)
  end
  traversal(root, ans)
  ans
end
