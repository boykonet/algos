#include <iostream>
#include <string>

#define PRINT(x) std::cout << x

/* #include "reverse.hpp" */
/* void		reverse() */
/* { */
/* 	Solution	s; */
/*     int			num; */
    
/*     std::cin >> num; */
/*     PRINT(s.reverse(num)); */
/* } */

/* #include "Palindrome.hpp" */
/* void		palindrome() */
/* { */
/* 	Solution	p; */
/* 	int			num; */

/* 	std::cin >> num; */
/* 	PRINT(p.isPalindrome(num)); */
/* } */

/* #include "roman_to_int.hpp" */
/* void	roman_to_int() */
/* { */
/* 	Solution		roman; */
/* 	std::string		s; */

/* 	std::cin >> s; */
/* 	PRINT(roman.romanToInt(s) << '\n'); */
/* } */

/* #include "valid.hpp" */
/* void		valid() */
/* { */
/* 	Solution		v; */
/* 	std::string		s; */

/* 	std::cin >> s; */
/* 	PRINT(v.isValid(s)); */
/* } */

/* #include "strStr.hpp" */
/* void	strStr() */
/* { */
/* 	Solution	a; */

/* 	PRINT(a.strStr("a", "a")); */
/* } */

#include "searchInsert.hpp"

void	searchInsert()
{
	Solution	e;
	std::vector<int>	v = {1};

	PRINT(e.searchInsert(v, 1));
}

int		main()
{
	/* reverse(); */
	/* palindrome(); */
	/* roman_to_int(); */
	/* valid(); */
	/* strStr(); */
	searchInsert();

	return 0;
}
