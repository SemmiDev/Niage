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
        List<Integer> temp;
        StringBuilder output = new StringBuilder();

        int N = in.nextInt();
        for (int i = 0; i < N; i++) {
            int height = in.nextInt();
            tanah.add(height);
        }

        int x,y,min,max,indexMin,size;
        int first;
        int second;
        int third;

        int Q = in.nextInt();
        while(Q-- > 0) {
            first = -1;
            second = -1;
            third = -1;
            
            String query = in.next();
            if (query.equals("A")) {
                y = in.nextInt();
                tanah.add(y);
                // show(tanah);
            } else if (query.equals("U")) {
                x = in.nextInt();
                y = in.nextInt();
                size = tanah.size();
                if (x >= 0 && x < size) {
                    tanah.set(x,y);
                }
                // show(tanah);
            } else {
                min = Collections.min(tanah);
                indexMin = tanah.indexOf(min);
                max = 0;

                temp = new ArrayList<Integer>();
                try {
                    first = tanah.get(indexMin);
                    temp.add(first);
                    // show(tanah);
                }catch(Exception e) {}

                try {
                    second = tanah.get(indexMin-1);
                    temp.add(second);
                    // show(tanah);
                }catch(Exception e) {}

                try {
                    third = tanah.get(indexMin+1);
                    temp.add(third);
                    // show(tanah);
                }catch(Exception e) {}

                max = Collections.max(temp);

                if (first != -1) {
                    tanah.set(indexMin, max);
                    // show(tanah);
                }
                if (second != -1) {
                    tanah.set(indexMin-1, max);
                    // show(tanah);
                }
                if (third != -1) {
                    tanah.set(indexMin+1, max);
                    // show(tanah);
                }

                output.append(max).append(" ").append(indexMin).append("\n");
            }
        }
        System.out.println(output);
        out.flush();
    }

    static void show(List<Integer> tanah) {
        for (int x : tanah) {
            System.out.print(x);
        }
        System.out.println();
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
