#ifndef PALINDROME_HPP
# define PALINDROME_HPP

class	Solution
{
public:
    bool	isPalindrome(int x)
	{
		long		num1;
		long		num2;
		int			count;
		long 		d;

		num1 = x;
		d = 1;
		if (num1 < 0)
			return false;

		num2 = num1;

		count = _count(num1);
		for (int i = 0; i < count - 1; ++i)
			d *= 10;

		for (; num1 > 0 && num2 > 0; )
		{
			if (num1 / d != num2 % 10)
				return false;
			num1 = num1 % d;
			num2 = num2 / 10;
			d /= 10;
		}
		return true;
	}
private:
	int		_count(long num)
	{
		int	count;


		count = 0;
		while (num > 0)
		{
			++count;
			num = num / 10;
		}
		return count;
	}
};

#endif // PALINDROME_HPP

