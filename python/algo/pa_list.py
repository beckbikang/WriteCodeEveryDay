#!/opt/homebrew/bin/python3.10


class ListNode:

    def __init__(self, next, val:int) -> None:
        self.next = None
        self.val = val
    

class Solution:
    def  answer(self, head:ListNode)->bool:
        
        pre = None
        slow = fast = head

        while fast and fast.next:
            # 下下一个
            fast = fast.next.next
            # 下一个
            next = slow.next
            # slow的下一个指向pre
            slow.next = pre
            # pre 等于slow
            pre = slow
            # slow = next
            slow = next
        
        #print(slow.val, fast.val)
        if fast:
            fast = fast.next
        while slow :
            if (pre is None) or (slow.val != pre.val):
                return False
            #print(pre.val, slow.val)
            pre = pre.next
            slow = slow.next
        return True

  

if "__main__" == __name__:
    print("start running")
    s = Solution()
    # 1 2  2 1
    n1 = ListNode(None,1)
    n2 = ListNode(None,2)
    n3 = ListNode(None,2)
    n4 = ListNode(None,1)

    n1.next = n2
    n2.next = n3
    n3.next = n4
   
    print(s.answer(n1))
    n5 = ListNode(None,1)
    n4.next = n5
    print(s.answer(n1))


