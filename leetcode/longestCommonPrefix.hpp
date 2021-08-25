#ifndef LONGEST_COMMON_PREFIX_HPP
# define LONGEST_COMMON_PREFIX_HPP

# include <iostream>
# include <string>
# include <vector>
# include <algorithm>

class Solution
{
public:
	std::string		longestCommonPrefix(std::vector< std::string > &strs)
	{
		std::string		s;
		std::string		b;
		std::string		e;

		s = "";

		std::sort(strs.begin(), strs.end());
		
		if (strs.size() == 1)
			return *strs.begin();

		b = *strs.begin();
		e = *(strs.end() - 1);

		for (int i = 0; b[i] && e[i]; ++i)
		{
			if (b[i] == e[i])
				s += b[i];
			else
				break ;
		}
		return s;
    }
};

#endif
