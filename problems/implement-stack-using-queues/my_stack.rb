class MyStack
    def initialize()
        @stack = []
    end


=begin
    :type x: Integer
    :rtype: Void
=end
    def push(x)
        @stack << x
    end


=begin
    :rtype: Integer
=end
    def pop()
        @stack.deleted_at(-1)
    end


=begin
    :rtype: Integer
=end
    def top()
        return @stack[-1]
    end


=begin
    :rtype: Boolean
=end
    def empty()
        return @stack.empty?
    end


end

# Your MyStack object will be instantiated and called as such:
# obj = MyStack.new()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()
