import java.io.*;
import java.util.*;

public class Lab1 {
    private static InputReader in = new InputReader(System.in);
    private static PrintWriter out = new PrintWriter(System.out);
    

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
        

        int[] temps = new int[N-1];
        int tempsIndex = 0;
        int sell = 0;
        for (int i = 1; i < M; i++){
            if (sequence.get(i).equals("*")){
                temps[tempsIndex] = sell;
                tempsIndex += 1;
                sell = 0;
            }
            else {
                sell += Integer.parseInt(sequence.get(i));
            }
        }
        int result = temps[0];
        int jumlah = temps[0];    
        for (int j = 1; j < temps.length; j++){
            if (result < 0 && temps[j] > result){
                result = temps[j];
                jumlah = temps[j];
            } else if (result >= 0){
                jumlah += temps[j];
                if (jumlah > result) {
                    result = jumlah; 
                }
            }
        }
        return result;
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
