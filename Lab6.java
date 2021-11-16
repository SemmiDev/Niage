import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.PrintWriter;
import java.util.*;

class Lab6 {
    private static InputReader in;
    private static PrintWriter out;

    public static void main(String[] args) {
        InputStream inputStream = System.in;
        in = new InputReader(inputStream);
        OutputStream outputStream = System.out;
        out = new PrintWriter(outputStream);

        List<Integer> tanah = new ArrayList<Integer>();
        List<Integer> temp;
        StringBuilder output = new StringBuilder();

        int N = in.nextInt();
        int size = N;
        for (int i = 0; i < N; i++) {
            int height = in.nextInt();
            tanah.add(height);
        }

        int x,y,min,max,indexMin;
        int Q = in.nextInt();
        while(Q-- > 0) {
            String query = in.next();
            if (query.equals("A")) {
                y = in.nextInt();
                tanah.add(y);
                size = tanah.size();
            } else if (query.equals("U")) {
                x = in.nextInt();
                y = in.nextInt();
                if (x >= 0 && x < size) {
                    tanah.set(x,y);
                }
            } else {
                min = Collections.min(tanah);
                indexMin = tanah.indexOf(min);
                max = 0;

                int first = -1;
                int second = -1;
                int third = -1;

                temp = new ArrayList<Integer>();
                try {
                    first = tanah.get(indexMin);
                }catch (Exception e) {}
                try {
                   second = tanah.get(indexMin-1);
                }catch (Exception e) {}
                try {
                   third = tanah.get(indexMin+1);
                }catch (Exception e) {}

                if (first != -1) {
                    temp.add(first);
                }
                if (second != -1) {
                    temp.add(second);
                }
                if (third != -1) {
                    temp.add(third);
                }

                max = Collections.max(temp);
                if (first != -1) {
                    tanah.set(indexMin, max);
                }
                if (second != -1) {
                    tanah.set(indexMin-1, max);
                }
                if (third != -1) {
                    tanah.set(indexMin+1, max);
                }

                output.append(max).append(" ").append(indexMin).append("\n");
            }
        }
        System.out.println(output);
        out.flush();
    }

    static class InputReader {
        public BufferedReader reader;
        public StringTokenizer tokenizer;

        public InputReader(InputStream stream) {
            reader = new BufferedReader(new InputStreamReader(stream), 32768);
            tokenizer = null;
        }

        public String next() {
            while (tokenizer == null || !tokenizer.hasMoreTokens()) {
                try {
                    tokenizer = new StringTokenizer(reader.readLine());
                } catch (IOException e) {
                    throw new RuntimeException(e);
                }
            }
            return tokenizer.nextToken();
        }

        public int nextInt() {
            return Integer.parseInt(next());
        }

    }
}
