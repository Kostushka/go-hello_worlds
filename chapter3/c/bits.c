#include <stdio.h>
#include <string.h>

#define LEN 8

void int_to_bits(unsigned int x, char *arr);

int main(void) {

	unsigned int x = 1<<1 | 1<<5;
	unsigned int y = 1<<1 | 1<<2;

	char arr_x[LEN];
	memset(arr_x, 0, sizeof(arr_x));
	char arr_y[LEN];
	memset(arr_y, 0, sizeof(arr_y));	
	char arr_xy[LEN];
	memset(arr_xy, 0, sizeof(arr_xy));
	
	int_to_bits(x, arr_x);
	for (int i = 0; i < LEN; i++) {
		printf("%d", arr_x[i]);
	}
	printf(" \\\\x");
	printf("\n");
	
	int_to_bits(y, arr_y);
	for (int i = 0; i < LEN; i++) {
		printf("%d", arr_y[i]);
	}
	printf(" \\\\y");
	printf("\n");

	int_to_bits(x&y, arr_xy);
	for (int i = 0; i < LEN; i++) {
		printf("%d", arr_xy[i]);
	}
	printf(" \\\\x&y");
	printf("\n");

	memset(arr_xy, 0, sizeof(arr_xy));
	int_to_bits(x|y, arr_xy);
	for (int i = 0; i < LEN; i++) {
		printf("%d", arr_xy[i]);
	}
	printf(" \\\\x|y");
	printf("\n");

	memset(arr_xy, 0, sizeof(arr_xy));
	int_to_bits(x^y, arr_xy);
	for (int i = 0; i < LEN; i++) {
		printf("%d", arr_xy[i]);
	}
	printf(" \\\\x^y");
	printf("\n");
	
	return 0;
}

void int_to_bits(unsigned int x, char *arr) {

	for (int i = LEN - 1; x; x >>= 1, i--) {
		arr[i] = x & 1;
	}
}
