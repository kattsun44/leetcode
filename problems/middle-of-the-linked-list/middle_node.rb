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
def middle_node(head)
  middle = head
  last = head
  while last != nil && last.next != nil do
    middle = middle.next
    last = last.next.next
  end
  return middle
end

p middle_node(ListNode.new(1, ListNode.new(2, ListNode.new(3, ListNode.new(4, ListNode.new(5, nil))))))
p middle_node(ListNode.new(1, ListNode.new(2, ListNode.new(3, ListNode.new(4, ListNode.new(5, ListNode.new(6, nil)))))))
