#include <stdio.h>
#include <stdlib.h>

int solve1(int m) {
	int n, c = 0;
	int* a;

    scanf("%d", &n);

    char* s = (char *)malloc(n * sizeof(char));
    scanf("%s", s);

    // inverse digits into an array
	a = (int *)malloc(sizeof(int)*n);
	for (int i = 0; i < n; i++) {
		a[i] = s[n-i-1] - '0';
	}

	int j, p, d, x;
	for (int i = 1; i < 64; i++) {
		j = i;
		p = 1;
		d = 0;
		x = 0; // generated number
		while (j > 0 && d < n) {
			// a[d] digit is included ?
			if (j%2 == 1) {

				x = a[d]*p + x;

				if (p == 1 && a[d]%2 == 1) {
					// last digit is odd
					break;
				}

				p = p * 10;

			}

			d++;
			j = j / 2;
		}

		//  check generated number
		if (j == 0 && x%8 == 0) {
			c++;
		}

	}

	printf("%d\n", c);
}

int main() {
	

    return 0;
}