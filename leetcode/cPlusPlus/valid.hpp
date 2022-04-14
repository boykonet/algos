#ifndef VALID_HPP
# define VALID_HPP

# include <iostream>
# include <string>
# include <stack>

class	Solution
{
public:
    bool	isValid(std::string s)
	{
		std::stack< char >	st;
		char				symb;
		size_t				i;

		i = 0;
		for (; i < s.length(); ++i)
		{
			symb = s[i];

			if (symb == '(' || symb == '{' || symb == '[')
				st.push(symb);
			if (st.empty())
				return false;
			else if (symb == ')' || symb == '}' || symb == ']')
			{
				if (symb == ')')
				{
					if (st.top() == '(')
						st.pop();
					else
						return false;
				}
				else
				{
					if (st.top() == (symb - 2))
						st.pop();
					else
						return false;
				}
			}
		}
		if (st.empty() && i == s.length())
			return true;
		return false;
    }
};

#endif
