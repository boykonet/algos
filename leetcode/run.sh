#!/bin/bash

ARGS=$1

if [[ $# == 1 ]]
then
	clang++ -g -fsanitize=address -Wall -Wextra -Werror --std=c++17 $ARGS -o prog
	./prog
else
	printf "\e[1;31mAdd argument!!!\n\e[0m"
fi
