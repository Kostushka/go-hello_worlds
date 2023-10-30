#include <stdio.h>

double ftoc(double value);

int main(void) {

	const double boiling = 212.0;
	const double freezing = 32.0;

	printf("%gF = %gC\n", boiling, ftoc(boiling)); 
	printf("%gF = %gC\n", freezing, ftoc(freezing)); 

	return 0;
}

double ftoc(double value) {
	return (value - 32) * 5 / 9;	
}
