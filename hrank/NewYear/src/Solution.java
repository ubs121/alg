// https://www.hackerrank.com/challenges/newyear-game/problem

import java.util.Scanner;

/**
 *
 * @author ub
 */
public class Solution {

    /**
     *
     * @param s score, |sb - sk| % 3 (!!! abs(sb-sk))
     * @param r remainder, r[0], r[1], r[2]
     * @param b whose turn, true = balsa, false = koca
     * @return
     */
    private int run(int s, int[] r, boolean b) {
        if (s == 0) {
            return 0; // Koca
        }

        int rsum, v;
        // r[0]+r[1]+r[2] > 1
        for (int i = 0; i < 3; i++) {
            if (r[i] > 0) {
                r[i]--; // take r[i]

                rsum = 0;
                if (b) {
                    // add if Balsa's move
                    rsum = (s + i) % 3; // (A + B) mod C = (A mod C + B mod C) mod C
                } else {
                    // subtract if Koca's move
                    rsum = (s - i) % 3; // (A - B) mod C = (A mod C - B mod C) mod C
                    
                    if (rsum < 0) {
                        rsum += 3;
                    }
                }

                v = run(rsum, r, !b); // recursive call
                if (v == 0) {
                    return 0; // Koca
                }

                // TODO: memorize (rsum, r, !b) -> 'v'
                r[i]++; // reverse, put it back
            }
        }

        // no more number in 'r[]', r[0]+r[1]+r[2] == 0
        return 1; // Balsa
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int T = in.nextInt();

        // solution model
        Solution agent = new Solution();

        for (int t = 0; t < T; t++) {
            int n = in.nextInt();

            int[] r = new int[3];
            for (int i = 0; i < n; i++) {
                r[in.nextInt() % 3]++;
            }

            // debug
            for (int i = 0; i < r.length; i++) {
                System.out.printf("%d ", r[i]);
            }
            System.out.printf("\n");

            if (agent.run(0, r, true) == 0) {
                System.out.println("Koca");
            } else {
                System.out.println("Balsa");
            }
        }

    }

}
