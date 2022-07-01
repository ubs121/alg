
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

// NOTE: solve using stack
public class FrameLabeling {

    /*
    * 
    
    2d space solution
    
    aekbakbdefdkdemedm
    -----------------------> pos
 |  a   a
\/   e      e    e e
 y    k  k     k
       b  b
           d  d d   d
             f
                  m  m
    
    Answer: aedm
     */
    private String solve(String s) {
        int[] pos = new int[26];      // 'pos[i]' is position in 'y[]' , pos['a'] = 0, pos['e'] = 1 etc. 
        Arrays.fill(pos, 0, pos.length, -1); // pos[*] == -1 means doesn't exist

        // collect bars and its heights
        List<Label> y = new ArrayList<>(); // 'y[i]' is ordered letters and height of bar
        for (int i = 0; i < s.length(); i++) {
            int p = pos[s.charAt(i) - 'a']; // find position for s[i] in 'y' axis
            if (p == -1) { // new letter
                pos[s.charAt(i) - 'a'] = y.size();
                y.add(new Label(s.charAt(i), i, i + 1));
            } else { // existing
                y.get(p).extendTo(i + 1);
            }

        }

        

        // collect visible bars (from top)
        StringBuilder ret = new StringBuilder();
        ret.append(y.get(0).label); // first letter is always visible
        
        Label last = y.get(0);
        for (int i = 1; i < y.size(); i++) {
            
            if (y.get(i).end > last.end) { // is visible?
                ret.append(y.get(i).label);
                last = y.get(i);
            }
        }

        return ret.toString();
    }

    // helper class
    class Label {

        char label;
        int start, end;

        private Label(char c, int i, int j) {
            label = c;
            start = i;
            end = j;
        }

        public int len() {
            return end - start;
        }

        public void extendTo(int j) {
            end = j;
        }

    }

    public static void main(String[] args) {
        FrameLabeling model = new FrameLabeling();
        // aekbakbdefdkdemedm
        String ret = model.solve("narantuya");
        System.out.println(ret);
    }

}
