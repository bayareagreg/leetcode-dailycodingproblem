import java.util.HashSet;
import java.util.Set;
import java.util.TreeSet;

public class Main {
    public static Set<String> findRotatedStrings(String s, int k) {
        HashSet<String> set = new HashSet<>();
        for (int i = 0; i < k; i++) {
            StringBuilder buf = new StringBuilder();
            for (int j = 0; j < i; j++) {
                buf.append(s.charAt(j));
            }
            buf.append(s.substring(i+1));
            buf.append(s.charAt(i));
            set.add(buf.toString());
        }
        return set;
    }


    public static void smallestRotatedStringHelper(String s, int len, int k, TreeSet<String> all) {
        if (len > 0) {
            Set<String> thisTurn = findRotatedStrings(s, k);
            all.addAll(thisTurn);

            for (String s1 : thisTurn) {
                smallestRotatedStringHelper(s1, len-1, k, all);
            }
        }
    }

    public static String smallestRotatedString(String s, int k) {
        TreeSet<String> all = new TreeSet<>();
        smallestRotatedStringHelper(s, s.length(), k, all);
        return all.first();
    }

    public static void main(String[] args) {

        //TIP Press <shortcut actionId="ShowIntentionActions"/> with your caret at the highlighted text
        // to see how IntelliJ IDEA suggests fixing it.
//        System.out.println(smallestRotatedString("daily", 1));
//        System.out.println(smallestRotatedString("daily", 2));
//        System.out.println(smallestRotatedString("daily", 3));

        System.out.println(smallestRotatedString("GEEKSQUIZ", 4));
        //System.out.println(smallestRotatedString("GFG", 1));
        //System.out.println(smallestRotatedString("EEKSFORGEEKSG", 1));
    }
}