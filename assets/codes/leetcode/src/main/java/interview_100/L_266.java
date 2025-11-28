package interview_100;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class L_266 {
    public boolean canPermutePalindrome(String s) {
        char[] chars = s.toCharArray();
        int[] ascii = new int[128];
        for (char aChar : chars) {
            ascii[aChar]++;
        }
        if (s.length() % 2 == 0) {
            for (Integer i : ascii) {
                if (i % 2 == 1) {
                    return false;
                }
            }
        } else {
            int count = 0;
            for (Integer i : ascii) {
                if (i % 2 == 1) {
                    count++;
                }
                if (count > 1) {
                    return false;
                }
            }
        }
        return true;
    }

    public static void main(String[] args) {

    }
}
