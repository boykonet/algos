#ifndef ROMAN_TO_INT_HPP
# define ROMAN_TO_INT_HPP

# include <iostream>
# include <string>
# include <map>

class	Solution
{
public:
	Solution() : _keys()
	{
		_keys["I"] = 1;
		_keys["V"] = 5;
		_keys["X"] = 10;
		_keys["L"] = 50;
		_keys["C"] = 100;
		_keys["D"] = 500;
		_keys["M"] = 1000;
		_keys["IV"] = 4;
		_keys["IX"] = 9;
		_keys["XL"] = 40;
		_keys["XC"] = 90;
		_keys["CD"] = 400;
		_keys["CM"] = 900;
	}

	~Solution()
	{
		_keys.clear();
	}

    int	romanToInt(std::string s)
	{
		int								num;
		std::string						sub;
		size_t							count;

		num = 0;
		count = 0;
		while (s.length())
		{
			count = 2;
			auto	it = _keys.find(s.substr(0, count));

			if (it == _keys.end())
			{
				count = 1;
				it = _keys.find(s.substr(0, count));
			}
			count = (s.length() < count) ? s.length() : count;
			num = num + it->second;
			s = s.substr(count, s.length());
		}

		return num;
    }
private:
	std::map< std::string, int >	_keys;
};

#endif // ROMAN_TO_INT_HPP
