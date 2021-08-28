#include <string>

class	Solution
{
public:
	int		strStr(std::string haystack, std::string needle)
	{
		size_t		indx_h;
		int			indx;

		indx_h = 0;
		if (needle.length() == 0)
			return (0);
		if (needle.length() > haystack.length())
			return (-1);
		while (haystack[indx_h])
		{
			indx = 0;
			while (haystack[indx_h + indx]
			&& needle[indx]
			&& (haystack[indx_h + indx] == needle[indx]))
				++indx;
			if (indx == needle.length())
				return static_cast<int>(indx_h);
			if ((haystack.length() - indx_h) < needle.length())
				break ;
			++indx_h;
		}
		return (-1);
	}
};
