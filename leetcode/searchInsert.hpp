#ifndef SEARCH_INSERT_HPP
# define SEARCH_INSERT_HPP

# include <algorithm>
# include <vector>

class	Solution
{
public:
	int			searchInsert(std::vector<int>& nums, int target)
	{
		int		left = 0;
		int		right = nums.size() - 1;
		int		elem;

		if (target < nums[left])
			return 0;
		else if (target > nums[right])
			return static_cast<int>(nums.size());
        while (left < right)
        {
            elem = (left + right) / 2;
            if (nums[elem] <= target)
                right = elem;
            else
                left = elem + 1;
        }
        return left;
    }
};

#endif // SEARCH_INSERT_HPP
