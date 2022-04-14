/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2)
    {
        ListNode    *head = nullptr;
        ListNode    *curr = nullptr;
        
        if (!l1 && !l2)
            return head;
        if (l1 && !l2)
        {
            head = new ListNode(l1->val);
            l1 = l1->next;
        }
        else if (!l1 && l2)
        {
            head = new ListNode(l2->val);
            l2 = l2->next;
        }
        else
        {
            if (l1->val < l2->val)
            {
                head = new ListNode(l1->val);
                l1 = l1->next;
            }
            else
            {
                head = new ListNode(l2->val);
                l2 = l2->next;
            }
        }
        
        curr = head;
        
        while (1)
        {
            if (!l1 && !l2)
                break ;
            if (l1 && !l2)
            {
                curr->next = new ListNode(l1->val);
                l1 = l1->next;
            }
            else if (!l1 && l2)
            {
                curr->next = new ListNode(l2->val);
                l2 = l2->next;
            }
            else
            {
                if (l1->val < l2->val)
                {
                    curr->next = new ListNode(l1->val);
                    l1 = l1->next;
                }
                else
                {
                    curr->next = new ListNode(l2->val);
                    l2 = l2->next;
                }
            }
            curr = curr->next;
        }
        return head;
    }
};
