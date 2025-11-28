package interview_100;

public class L_1056 {
    public boolean confusingNumber(int n) {
        int sum = 0;
        int temp = n;
        while (n != 0) {
            int val = n % 10;
            if (val == 2 || val == 3 || val == 4 || val == 5 || val == 7) {
                return false;
            }
            sum *= 10;
            if (val == 9) {
                sum += 6;
            } else if (val == 0 || val == 1 || val == 8) {
                sum += val;
            } else {
                sum += 9;
            }
            n = n / 10;
        }
        if (sum == temp) {
            return false;
        }
        return true;
    }

    public static void main(String[] args) {
        L_1056 l1056 = new L_1056();
        System.out.println(l1056.confusingNumber(916));
    }
}
