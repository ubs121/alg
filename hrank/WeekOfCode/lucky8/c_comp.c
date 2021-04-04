#include <stdio.h>
#include <stdlib.h>

int n, c;
int* a;

void init_arr(int m) {
	n = 0;
	int temp = m;
	while (temp > 0) {
		temp = temp / 10;
		n++;
	}

	// inverse digits into an array
	a = (int *)malloc(sizeof(int)*n);
	int t = 0;
	while (m > 0 && t < n) {
		a[t] = m %10;
		m = m / 10;
		t++;
	}
}


/**
*	i - digit position
* 	x - previuos value
* 	t - which tenth
 */
void check(int i, int x, int t) {
	x = a[i]*t + x; // concate a number at position 'i'

	if (x%8 == 0) {
		c++;
	} else {
		if (t == 100) {
			// don't go further if last 3 digits are not divisible by 8
			return;
		}
	}

	for (int ii = i + 1; ii < n; ii++) {
		check(ii, x, t*10);
	}
}

int solve1(int m) {
	init_arr(m);

	c = 0;

	for (int i = 0; i < n; i++) {

		// check only even numbers
		if (a[i]%2 == 0) {
			check(i, 0, 1);
		}
	}

	return c;
}

int solve2(int m) {
	init_arr(m);

	c = 0;


	int j, p, d, x;

	// 6 position => 2^6 = 64 variants
	for (int i = 1; i < 64; i++) { 

		j = i;
		p = 1;
		d = 0;
		x = 0; // generated number

		while (j > 0 && d < n) {
			// a[d] digit is included ?
			if (j & 1) {

				x = a[d]*p + x;

				if (p == 1 && a[d]%2 == 1) {
					// last digit is odd
					break;
				}

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

	return c;
}




int main() {
	int s1, s2;
	for (int i=1; i<300000;i++) {
		s1 = solve1(i);
		s2 = solve2(i);

		if (s1!=s2) {
			printf("%d %d\n", s1, s2);
		}
	}
    

    return 0;
}