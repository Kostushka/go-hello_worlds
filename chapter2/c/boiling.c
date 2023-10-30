#include <stdio.h>

const int boiling = 212.0;

int main(void) {
	
	float f = boiling;
	float c = (f - 32) * 5 / 9;

	printf("Температура кипения = %gF или %gC\n", f, c);
	
	return 0;
}
