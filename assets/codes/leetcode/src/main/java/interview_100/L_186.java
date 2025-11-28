package interview_100;

import java.util.Arrays;

public class L_186 {
    public void reverseWords(char[] s) {
        int start = 0, end = s.length - 1;
        while (start <= end) {
            char temp = s[start];
            s[start] = s[end];
            s[end] = temp;
            start++;
            end--;
        }
        start = 0;
        end = 0;
        while (start < s.length) {
            while (end < s.length && s[end] != ' ') end++;
            int tmp= end + 1;
            end--;
            while (start < end) {
                char temp = s[start];
                s[start] = s[end];
                s[end] = temp;
                start++;
                end--;
            }
            start = tmp;
            end = tmp;
        }
    }

    public static void main(String[] args) {

    }
}
