package interview_100;

import java.util.HashMap;
import java.util.Map;

public class L_159 {
    public int lengthOfLongestSubstringTwoDistinct(String s) {
        Map<Character, Integer> map = new HashMap<>();
        int res = 0;
        int left = 0;
        for (int i = 0; i < s.length(); i++) {
            map.put(s.charAt(i), map.getOrDefault(s.charAt(i), 0) + 1);
            if (map.size() <= 2) {
                res = Math.max(res, i - left + 1);
            }
            while (map.size() > 2) {
                map.put(s.charAt(left), map.get(s.charAt(left)) - 1);
                if (map.get(s.charAt(left)) <= 0) {
                    map.remove(s.charAt(left));
                }
                left++;
            }
            res = Math.max(res, i - left + 1);
        }
        return res;
    }

    public static void main(String[] args) {
        L_159 l159 = new L_159();
        System.out.println(l159.lengthOfLongestSubstringTwoDistinct("eceba"));
        System.out.println(l159.lengthOfLongestSubstringTwoDistinct("ccaabbb"));
        System.out.println(l159.lengthOfLongestSubstringTwoDistinct(""));
        System.out.println(l159.lengthOfLongestSubstringTwoDistinct("a"));
        System.out.println(l159.lengthOfLongestSubstringTwoDistinct("ab"));
        System.out.println(l159.lengthOfLongestSubstringTwoDistinct("abc"));
    }
}
