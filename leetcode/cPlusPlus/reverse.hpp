#ifndef REVERSE_HPP
# define REVERSE_HPP

#define MAX_INT		2147483647
#define MIN_INT		-2147483648

class	Solution
{
public:
	Solution()
	{ }
	Solution(Solution const &r)
	{
		*this = r;
	}
	~Solution()
	{ }
	Solution&		operator=(Solution const &r)
	{
		if (this != &r)
			;
		return *this;
	}
    int				reverse(int x)
    {
		long	num;
        long	reverse_num;
		int		count_num;
        int		sign;
		long	d;
        
		num = x;
		reverse_num = 0;
		count_num = 0;
		sign = 1;
		d = 1;

        if (num < 0)
		{
			num *= (-1);
			sign = -1;
		}

		count_num = _count(num);

		for (int i = 0; i < count_num; ++i)
			d = d * 10;
		while (num > 0)
		{
			reverse_num = reverse_num + (num % 10) * (d / 10);
			if (reverse_num > MAX_INT || reverse_num < MIN_INT)
				return 0;
			d /= 10;
			num /= 10;
		}
        return reverse_num * sign;
    }

private:
	int				_count(int num)
	{
		int	count = 0;

		while (num > 0)
		{
			++count;
			num /= 10;
		}
		return count;
	}
};


#endif
