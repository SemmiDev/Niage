import java.io.*;
import java.util.*;

public class Lab1 {
    private static InputReader in = new InputReader(System.in);
    private static PrintWriter out = new PrintWriter(System.out);
    private static List<Integer> hasilPenjumlahan = new ArrayList<Integer>();
    private static List<Integer> highlight= new ArrayList<>();
    private static  int[] results;

    /**
     * The main method that reads input, calls the function 
     * for each question's query, and output the results.
     * @param args Unused.
     * @return Nothing.
     */
    public static void main(String[] args) {
        int N = in.nextInt();   // banyak bintang
        int M = in.nextInt();   // panjang sequence

        List<String> sequence = new ArrayList<String>();
        for (int i = 0; i < M; i++) {
            String temp = in.next();
            sequence.add(temp);
        }

        int maxMoney = getMaxMoney(N, M, sequence);
        out.println(maxMoney);
        out.close();
    }

    public static int getMaxMoney(int N, int M, List<String> sequence) {
        // int result = 0;
        int size = 0;

        // constraint 1
        if ((N >= M) || (N < 2) || (M > 300000) ) {
            System.exit(1);
        }

        // constraint 2
        int value;
        for (int i = 1; i < sequence.size()-1; i++) {
            if (!sequence.get(i).equalsIgnoreCase("*")) {
                value = Integer.parseInt(sequence.get(i));
                if ((value < -1000) || (value > 1000)) {
                    System.exit(1);
                }
            }
        }

        // constraint 3
        boolean starPrefix = sequence.get(0).equalsIgnoreCase("*");
        boolean starSuffix = sequence.get(sequence.size() - 1).equalsIgnoreCase("*");
        if (!(starPrefix && starSuffix)) {
            System.exit(1);
        }

        // constraint 4
        boolean before, after;
        for (int i = 0; i < sequence.size()-1; i++) {
            before = sequence.get(i).equalsIgnoreCase("*");
            after = sequence.get(i+1).equalsIgnoreCase("*");
            if (before && after) {
                System.exit(1);
            }
        }

        // sum of numbers among 2 star(*)
        int lastStar = 0;
        int total = 0;
        boolean isStar;

        for (int i = 1; i < sequence.size(); i++) {
            isStar = sequence.get(i).equalsIgnoreCase("*");
            if (isStar) {
                for (int j = lastStar+1; j < i; j++) {
                    total += Integer.parseInt(sequence.get(j));
                }
                highlight.add(total);
                total = 0;
                lastStar = i;
            }
        }
        hasilPenjumlahan.addAll(highlight);
        int loop = 1;
        for (int i = 0; i < highlight.size()-1; i++) {
            combine(loop, highlight);
            loop++;
        }
        size = hasilPenjumlahan.size();
        results = hasilPenjumlahan.stream().mapToInt(i -> i).toArray();
        if (size <= 1000) {
            Arrays.sort(results);
        }else {
            Arrays.parallelSort(results);
        }
        return results[results.length-1];
    }

    static void combine(int r, List<Integer> numbers) {
        int total = 0;
        int con = r;

        for (int i = 0; i < numbers.size(); i++) {
            for (int j = i; j <= con; j++) {
                if (con >= numbers.size()) {
                    return;
                }
                total += numbers.get(j);
            }
            hasilPenjumlahan.add(total);
            total = 0;
            con++;
        }
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
