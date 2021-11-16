import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.PrintWriter;
import java.util.*;

public class Lab6 {
    private static InputReader in;
    private static PrintWriter out;

    public static void main(String[] args) {
        InputStream inputStream = System.in;
        in = new InputReader(inputStream);
        OutputStream outputStream = System.out;
        out = new PrintWriter(outputStream);

        List<Integer> tanah = new ArrayList<Integer>();
        StringBuilder output = new StringBuilder();

        int N = in.nextInt();
        for (int i = 0; i < N; i++) {
            int height = in.nextInt();
            tanah.add(height);
        }
        int x,y,min,max,indexMin,size,first,second,third,tempMax,t;
        int Q = in.nextInt();
        while(Q-- > 0) {
            first = -1;
            second = -1;
            third = -1;
            size = tanah.size();
            
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
                first = tanah.get(indexMin);
                t = size-1;

                if (indexMin > 0) {
                    second = tanah.get(indexMin-1);
                    if (second > first) {
                        first = second;
                    }
                }
                if (indexMin < t) {
                    third = tanah.get(indexMin+1);
                    if (third > first) {
                        first = third;
                    }
                }

                if (second != -1) {
                    tanah.set(indexMin-1, first);
                }
                if (third != -1) {
                    tanah.set(indexMin+1, first);
                }

                tanah.set(indexMin, first);
                output.append(first).append(" ").append(indexMin).append("\n");
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
