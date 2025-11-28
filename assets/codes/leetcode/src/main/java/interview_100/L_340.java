package interview_100;

import java.util.HashMap;
import java.util.Map;

public class L_340 {

    public int lengthOfLongestSubstringKDistinct(String s, int k) {
        Map<Character, Integer> map = new HashMap<>();
        int res = 0;
        int left = 0;
        for (int i = 0; i < s.length(); i++) {
            map.put(s.charAt(i), map.getOrDefault(s.charAt(i), 0) + 1);
            if (map.size() <= k) {
                res = Math.max(res, i - left + 1);
            }
            while (map.size() > k) {
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

    }
}
