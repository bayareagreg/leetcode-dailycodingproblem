import java.util.*;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {

    // O(s*w)
    public static List<Integer> findAnagramIndices(String w, String s) {
        List<Integer> list = new ArrayList<>();
        for (int i = 0; i+w.length() <= s.length() ; i++) {
            String subs = s.substring(i, i+w.length());
            if (isAnagram(w, subs)) {
                list.add(i);
            }
        }
        return list;
    }

    // O(n)
    public static boolean isAnagram(String w, String s) {
        if (w.length() != s.length()) {
            return false;
        }

        int[] freq = new int[26];
        for (int i = 0; i < w.length(); i++) {
            freq[w.charAt(i)-'a']++;
        }

        for (int i = 0; i < s.length(); i++) {
            if (freq[s.charAt(i)-'a'] == 0) {
                return false;
            }
            freq[s.charAt(i)-'a']--;
        }
        return true;
    }

    public static List<Integer> findAnagramIndices2(String w, String s) {
        int[] map = new int[26];
        List<Integer> result = new ArrayList<>();

        for( int i = 0; i < w.length(); i++ ) {
            map[w.charAt(i) - 'a']++;
        }

        int start = 0, end = 0;
        while (end < s.length()) {
            if (map[s.charAt(end) - 'a'] > 0) {
                map[s.charAt(end) - 'a']--;
                end++;

                if (end - start == w.length()) {
                    result.add(start);
                }
            } else if (start == end) {
                start ++;
                end ++;
            } else {
                map[s.charAt(start) - 'a']++;
                start++;
            }
        }
        return result;
    }

    public static void main(String[] args) {
        System.out.println(findAnagramIndices("ab", "abxaba"));
    }
}