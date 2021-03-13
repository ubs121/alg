
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 *
 * @author ub
 */
public class CountDistinct {

    public static void main(String[] args) {
        List<String> words = CountDistinct.extract2("ulaanbaatar", 3);
        
    }

    // using a map as counter
    private static List<String> extract1(String str, int k) {
        List<String> ret = new ArrayList<>();
        Map<Character, Integer> seg = new HashMap<>();
        Character c, l;

        for (int i = 0; i < k; i++) {
            c = str.charAt(i);
            if (seg.containsKey(c)) {
                seg.put(c, seg.get(c) + 1);
            } else {
                seg.put(c, 1);
            }
        }

        for (int i = 0; i + k < str.length(); i++) {
            // str[i:i+k] is distinct
            if (seg.size() == k) {
                // TODO: put into 
                // ret.add(str.substring(i, i+k));
                System.out.println(str.substring(i, i + k));
            }

            // add last char
            c = str.charAt(i + k);
            if (seg.containsKey(c)) {
                seg.put(c, seg.get(c) + 1);
            } else {
                seg.put(c, 1);
            }

            // remove first char
            l = str.charAt(i);
            seg.put(l, seg.get(l) - 1);
            if (seg.get(l) == 0) {
                seg.remove(l);
            }
        }
        return null;
    }
    
    // using regular array as counter
    private static List<String> extract2(String str, int k) {
        List<String> ret = new ArrayList<>();

        // count initial segment
        int[] seg = new int[25];
        int dist = 0;
        for (int i = 0; i < k; i++) {
            if (seg[str.charAt(i)-'a'] == 0) {
                dist++;
                seg[str.charAt(i)-'a'] = 1;
            } else {
                seg[str.charAt(i)-'a']++;
            }
        }

        
        for (int i = 0; i + k < str.length(); i++) {
            // str[i:i+k] is distinct
            if (dist == k) {
                // TODO: put into 
                // ret.add(str.substring(i, i+k));
                System.out.println(str.substring(i, i + k));
            }

            // add last char
            if (seg[str.charAt(i+k)-'a'] == 0) {
                dist++;
                seg[str.charAt(i+k)-'a'] = 1;
            } else {
                seg[str.charAt(i+k)-'a']++;
            }

            // remove first char
            seg[str.charAt(i)-'a']--;
            if (seg[str.charAt(i)-'a'] == 0) {
                dist--;
            }
        }
        return null;
    }
    
    // TODO: "extract3" using Hashing
}
