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
        StringBuilder output = new StringBuilder("\n");

        int N = in.nextInt();
        for (int i = 0; i < N; i++) {
            int height = in.nextInt();
            tanah.add(height);
        }

        int Q = in.nextInt();
        while(Q-- > 0) {
            String query = in.next();
            if (query.equals("A")) {
                int y = in.nextInt();
                tanah.add(y);
            } else if (query.equals("U")) {
                int x = in.nextInt();
                int y = in.nextInt();
                tanah.set(x,y);
            } else {                
                int min = Collections.min(tanah);
                int indexMin = tanah.indexOf(min);
                int max = 0;

                if (indexMin == 0) {
                    List<Integer> temp = new ArrayList<Integer>();
                    temp.add(tanah.get(indexMin));
                    temp.add(tanah.get(indexMin+1));

                    max = Collections.max(temp);
                    tanah.set(indexMin, max);
                    if (indexMin < N-1) {
                        tanah.set(indexMin + 1, max);
                    }
                }
                
                if (indexMin > 0) {
                    List<Integer> temp = new ArrayList<Integer>();
                    temp.add(tanah.get(indexMin-1));
                    temp.add(tanah.get(indexMin));
                    temp.add(tanah.get(indexMin+1));
                    max = Collections.max(temp);
                    
                    tanah.set(indexMin-1, max);
                    tanah.set(indexMin, max);
                    if (indexMin < N-1) {
                        tanah.set(indexMin + 1, max);
                    }
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
