/* package codechef; // don't place package name! */

import java.util.*;
import java.lang.*;
import java.io.*;

/* Name of the class has to be "Main" only if the class is public. */
class codechef {
    static class FastReader {
        BufferedReader br;
        StringTokenizer st;

        public FastReader() {
            br = new BufferedReader(
                    new InputStreamReader(System.in));
        }

        String next() {
            while (st == null || !st.hasMoreElements()) {
                try {
                    st = new StringTokenizer(br.readLine());
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }

            return st.nextToken();
        }

        int nextInt() {
            return Integer.parseInt(next());
        }

        long nextLong() {
            return Long.parseLong(next());
        }

        double nextdouble() {
            return Double.parseDouble(next());
        }

    }

    public static void main(String[] args) {
        FastReader scan = new FastReader();
        int T = scan.nextInt();
        while (T-- > 0) {
            int X = scan.nextInt();
            int Y = scan.nextInt();
            int Z = scan.nextInt();

            if (X * Y <= Z * 60 * 24)
                System.out.println("YES");
            else
                System.out.println("NO");

        }
    }

}