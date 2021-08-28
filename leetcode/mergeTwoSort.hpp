#ifndef MERGE_TWO_SORT_HPP
# define MERGE_TWO_SORT_HPP


struct	ListNode
{
	int			val;
	ListNode	*next;
	ListNode() : val(0), next(nullptr) {}
	ListNode(int x) : val(x), next(nullptr) {}
	ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class	Solution
{
public:
	ListNode*		mergeTwoLists(ListNode* l1, ListNode* l2)
	{
		ListNode	*curr;
		ListNode	*prev;
		ListNode	*ins;

		curr = l1;
		ins = l2;
		while (l1)
		{
			prev = curr;
			if (curr->val >= ins.val)

		}
    }
};


#endif
