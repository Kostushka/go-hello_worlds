#include <stdio.h>
#include <stdlib.h>

struct Slice {
	int len;
	int cap;
	int *arr;	
};

struct Slice make(int len, int cap);
struct Slice append(struct Slice int_slice, int x);

int main(void) {
	struct Slice int_slice = make(0, 4);

	int_slice = append(int_slice, 2);
	int_slice = append(int_slice, 3);
	int_slice = append(int_slice, 4);
	struct Slice b = append(int_slice, 5);
	printf("a: ");
	for (int i = 0; i < int_slice.len; i++) {
		printf("%d ", int_slice.arr[i]);
	}
	printf("\n");
	printf("b: ");
	for (int i = 0; i < b.len; i++) {
		printf("%d ", b.arr[i]);
	}
	printf("\n");
	int_slice.arr[0] = 0;
	printf("a: ");
	for (int i = 0; i < int_slice.len; i++) {
		printf("%d ", int_slice.arr[i]);
	}
	printf("\n");
	printf("b: ");
	for (int i = 0; i < b.len; i++) {
		printf("%d ", b.arr[i]);
	}
	printf("\n");

	for (int i = 0; i < 300; i += 20) {
		int_slice = append(int_slice, i);
		printf("len: %d\tcap: %d\n", int_slice.len, int_slice.cap);
		printf("[ ");
		for (int j = 0; j < int_slice.len; j++) {
			printf("%d ", int_slice.arr[j]);
		}
		printf("]\n");
	}
	
	return 0;
}

struct Slice make(int len, int cap) {
	struct Slice new;
	// выделить память размером в емкость
	new.arr = (int *) calloc(cap, sizeof(int));
	if (new.arr == NULL) {
		perror("calloc");
		return new;
	}
	new.len = len;
	new.cap = cap;
	
	return new;
}

struct Slice append(struct Slice int_slice, int x) {
	// создать новую структуру
	struct Slice new;

	int new_len = int_slice.len + 1;
	// проверить, что емкости хватает для добавления нового элемента
	if (new_len <= int_slice.cap) {
		new = int_slice;
		// обновить длину, добавить элемент по указателю
		new.len = new_len;
		new.arr[new_len - 1] = x;
	} else {
		// увеличить емкость
		new.cap = new_len;
		if (new.cap < 2 * int_slice.len) {
			new.cap = 2 * int_slice.len;
		}
		// перевыделить память
		new.arr = (int *) realloc(int_slice.arr, sizeof(int) * new.cap);
		if (new.arr == NULL) {
			perror("realloc");
			exit(1);
		}
		// обновить длину, добавить элемент по указателю
		new.len = new_len;
		new.arr[new_len - 1] = x;
	}
	return new;
}
