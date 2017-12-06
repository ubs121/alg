#include <stdio.h>
#include <stdlib.h>

int main() {
	int n, c = 0;
	int* a;

    scanf("%d", &n);

    char* s = (char *)malloc(512000 * sizeof(char));
    scanf("%s", s);

    // inverse digits into an array
	a = (int *)malloc(sizeof(int)*n);
	for (int i = 0; i < n; i++) {
		a[i] = s[n-i-1] - '0';
	}

	int j, p, d, x;

	// 6 position => 2^6 = 64 variants
	for (int i = 1; i < 64; i++) { 

		j = i; // marker
		p = 1; // 10 multiples
		d = 0; // array index, digit
		x = 0; // generated number

		while (j > 0 && d < n) {
			

			// a[d] digit is included ?
			if (j & 1) {
				if (p == 1 && a[d]%2 == 1) {
					// last digit is odd
					break; // break while
				}

				x = a[d]*p + x;
				p = p * 10;
			}

			d++;
			j = j >> 1;
		}

		//  check generated number
		if (j == 0 && x%8 == 0) {
			c++;
		}

	}

	printf("%d\n", c);

    return 0;
}