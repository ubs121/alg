#include <stdio.h>
#include <stdlib.h>

int main() {
	int n;
    scanf("%d", &n);

    char* s = (char *)malloc(512 * sizeof(char));
    scanf("%s", s);

    // inverse digits into an array 'd'
	int* d = (int *)malloc(sizeof(int)*n);
	for (int i = 0; i < n; i++) {
		d[i] = s[n-i-1] - '0';
	}


	int r = 0; // counter
	int d1, d2, d3;

	for (int a = 0; a < n; a++) {
		d1 = d[a];

		// skip if odd number
		if (d1%2 == 1) {
			continue;
		}

		if (d1%8 == 0) {
			r++; // single digit number
			printf("%d, %d <- %d\n", r, d1, a);
		}

		for (int b = a+1; b < n; b++) {
			d2 = d[b]*10+d1;
			if (d2%8 == 0) {
				r++; // 2 digit number
				printf("%d, %d <- %d %d\n", d2, r, b, a);
			}

			for (int c = b+1; c < n; c++) {
				d3 = d[c]*100+d2;

				if (d3%8 == 0) {
					r++; // 3 digit number
					
					// any further combinations are divisable by 8
					// other, 2^(n-3)-1 
					switch (n-3) {
						case 1: r = r + 1; break;
						case 2: r = r + 3; break;
						case 3: r = r + 7; break;
						case 4: r = r + 15; break; // just in case
						default: break;
					}

					printf("%d, %d <- %d %d %d\n", d3, r, c, b, a);
				} else {
					// skip to next digit
					continue;
				}
			}
		}
	}

	printf("%d\n", r % (1000000007));

    return 0;
}