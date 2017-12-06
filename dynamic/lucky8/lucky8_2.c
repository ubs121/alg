#include <stdio.h>
#include <stdlib.h>

int n, c;
int* a;

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


int main() {

    scanf("%d", &n);

    char* s = (char *)malloc(n * sizeof(char));
    scanf("%s", s);

    // inverse digits into an array
	a = (int *)malloc(sizeof(int)*n);
	for (int i = 0; i < n; i++) {
		a[i] = s[n-i-1] - '0';
	}

	// last digit
	for (int i = 0; i < n; i++) {

		// check only even numbers
		if (a[i]%2 == 0) {
			check(i, 0, 1);
		}
	}

	printf("%d\n", c);

    return 0;
}