# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end

# @param {ListNode} head
# @return {ListNode}
def delete_duplicates(head)
  res = []
  return res if head.nil?
  val = head.val
  res << val
  last = head
  while last != nil do
    if last.val > val
      res << last.val
    end
    val = last.val
    last = last.next
  end
  return res
end

p delete_duplicates(nil)
p delete_duplicates(ListNode.new(1, ListNode.new(1, ListNode.new(2, nil))))
p delete_duplicates(ListNode.new(1, ListNode.new(1, ListNode.new(2, ListNode.new(3, ListNode.new(3, nil))))))
p delete_duplicates(ListNode.new(1, ListNode.new(1, ListNode.new(1, ListNode.new(1, ListNode.new(2, nil))))))
