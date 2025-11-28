package interview_100;

public class L_161 {
    public boolean isOneEditDistance(String s, String t) {
        if (s.length() < t.length()) return isOneEditDistance(t, s);
        if (s.length() - t.length() > 1) return false;
        if (s.isEmpty() && s.equals(t)) return false;
        if (s.equals(t)) return false;

        int len = s.length();
        int i = 0, j = 0;
        if (s.length() != t.length()) {
            while (i < len && j < t.length()) {
                if (s.charAt(i) == t.charAt(j)) {
                    i++;
                    j++;
                } else {
                    i++;
                }
                if (i - j > 1) return false;
            }
        } else {
            int diff = 0;
            while (i < len && j < t.length()) {
                i++;
                j++;
                if (s.charAt(i) != t.charAt(j)) {
                    diff++;
                }
                if (diff > 1) return false;
            }
        }


        return true;
    }

    public static void main(String[] args) {
        L_161 l161 = new L_161();
        System.out.println(l161.isOneEditDistance("ab", "cab"));
    }
}
