import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static class Pair<K, V> {
        private final K key;
        private final V value;

        public Pair(K key, V value) {
            this.key = key;
            this.value = value;
        }

        public K getKey() {
            return key;
        }

        public V getValue() {
            return value;
        }

        public String toString() {
            return "(" + key + "," + value + ")";
        }
    }

    private static boolean isPalindrome(String s) {
        for (int i = 0, j = s.length() - 1; i <= j; i++, j--) {
            if (s.charAt(i) != s.charAt(j)) {
                return false;
            }
        }
        return true;
    }

    public static List<Pair<Integer, Integer>> generatePalindromePairs(String[] words) {
        List<Pair<Integer, Integer>> list = new ArrayList<>();
        for (int i = 0; i < words.length; i++) {
            for (int j = 0; j < words.length; j++) {
                if (i != j) {
                    String joined = words[i] + words[j];
                    if (isPalindrome(joined)) {
                        list.add(new Pair<>(i, j));
                    }
                }
            }
        }
        return list;
    }

    public static String reverseStr(String str) {
        StringBuilder sb = new StringBuilder(str);
        return sb.reverse().toString();
    }

    public static List<Pair<Integer, Integer>> generatePalindromePairs2(String[] words) {
        List<Pair<Integer, Integer>> list = new ArrayList<>();
        HashMap<String, Integer> map = new HashMap<>();
        for (int i = 0; i < words.length; i++) {
            map.put(words[i], i);
        }

        if(map.containsKey("")){
            int blankIdx = map.get("");
            for(int i = 0; i < words.length; i++){
                if(isPalindrome(words[i])){
                    if(i == blankIdx) continue;
                    list.add(new Pair<>(blankIdx, i));
                    list.add(new Pair<>(i, blankIdx));
                }
            }
        }

        map.forEach((k, v) -> {
            String reversed = reverseStr(k);
            Integer i = map.get(reversed);
            if (i != null && !i.equals(v)) {
                list.add(new Pair<>(i, v));
            }
        });

        //find the pair s1, s2 that
        //case1 : s1[0:cut] is palindrome and s1[cut+1:] = reverse(s2) => (s2, s1)
        //case2 : s1[cut+1:] is palindrome and s1[0:cut] = reverse(s2) => (s1, s2)
        for(int i = 0; i < words.length; i++){
            String cur = words[i];
            for(int cut = 1; cut < cur.length(); cut++){
                if(isPalindrome(cur.substring(0, cut))){
                    String cut_r = reverseStr(cur.substring(cut));
                    if(map.containsKey(cut_r)){
                        int found = map.get(cut_r);
                        if(found == i) continue;
                        list.add(new Pair<>(found, i));
                    }
                }
                if(isPalindrome(cur.substring(cut))){
                    String cut_r = reverseStr(cur.substring(0, cut));
                    if(map.containsKey(cut_r)){
                        int found = map.get(cut_r);
                        if(found == i) continue;
                        list.add(new Pair<>(i, found));
                    }
                }
            }
        }
        return list;
    }


    public static void main(String[] args) {
        //TIP Press <shortcut actionId="ShowIntentionActions"/> with your caret at the highlighted text
        // to see how IntelliJ IDEA suggests fixing it.
        String[] input1 = {"code", "edoc", "da", "d"};
        System.out.println(generatePalindromePairs2(input1));

        String[] input2 = {"abcd","dcba","lls","s","sssll"};
        System.out.println(generatePalindromePairs2(input2));
    }
}