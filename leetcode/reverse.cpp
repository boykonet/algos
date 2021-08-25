#include <iostream>

#define PRINT(x) std::cout << x
#define MAX_INT		2147483647
#define MIN_INT		-2147483648

class Solution
{
public:
    int reverse(int x)
    {
		long	num;
        long	reverse_num = 0;
		int		count_num = 0;
        int		sign = 1;
		long	d = 1;
        
		num = x;
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
			d /= 10;
			reverse_num = reverse_num + (num % 10) * d;
			if (reverse_num > MAX_INT || reverse_num < MIN_INT)
				return 0;
			num /= 10;
		}
        return reverse_num * sign;
    }

private:
	int	_count(int num)
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

int     main()
{
    Solution    s;
    int         num = 0;
    
    std::cin >> num;
    PRINT(s.reverse(num));
    
    return (0);
}
