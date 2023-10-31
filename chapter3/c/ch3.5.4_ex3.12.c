#include <stdio.h>

int anagram(char *s, char *t);
int get_len(char *str);

int main(void) {

	char s[] = "hello";
	char t[] = "ellho";
	if (anagram(s, t)) {
		printf("true\n");
	} else {
		printf("false\n");
	}

	char s1[] = "hello";
	char t1[] = "elho";
	if (anagram(s1, t1)) {
		printf("true\n");
	} else {
		printf("false\n");
	}
	
	return 0;
}

int anagram(char *s, char *t) {
	int len_s = get_len(s);
	int len_t = get_len(t);
	if (len_s != len_t) {
		return 0;
	}
	for (int i = 0; i < len_s; i++) {
		for (int j = 0; j < len_t; j++) {
			if (s[i] == s[j]) {
				s[j] = '-';
				break;
			}
			if (j == len_t - 1) {
				return 0;
			}
		}
	}
	return 1;
}

int get_len(char *str) {
	int len = 0;
	for (int i = 0; str[i] != '\0'; i++) {
		++len;
	}
	++len;
	return len;
}
