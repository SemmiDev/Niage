package ddp2;

import java.util.Scanner;

public class Lab05 {
    public static void main(String[] args) {

        // default array
        int[] list = new int[10];

        // class scanner untuk listen input dari console
        Scanner scan = new Scanner(System.in);

        // menampilkan menu-menu untuk interaksi user
        printMenu();

        // menerima pilihan user berdasarkan nomor menu
        int choice = scan.nextInt();

        // while loop untuk secara terus-menerus menampilkan menu kecuali pilihan user = 0
        while (choice != 0) {

            // deklarasi loc variable untuk menampung index pencarian oada selection sort
            int loc;

            // switch pilihan user
            switch (choice) {
                case 1 -> {
                    System.out.println("Berapa banyak elemen?");
                    int size = scan.nextInt();
                    // random angka dari 1-1000 sebanyak element yg di assign di variable size
                    list = randomize(new int[size]);
                }
                case 2 -> selectionSort(list);
                case 3 -> {
                    System.out.print("Masukan nilai pencarian: ");
                    // asssign loc variable melalui console
                    loc = linearSearch(list, scan.nextInt());
                    if (loc != -1) {
                        System.out.println("Ditemukan pada indeks ke-" + loc);
                    } else {
                        System.out.println("Tidak ditemukan");
                    }
                }
                case 4 -> printList(list);
                default -> System.out.println("Maaf, angka di luar batas");
            }
            printMenu();
            choice = scan.nextInt();
        }
        System.out.println("Bye!");
    }

    // method untuk menampilkan menu
    public static void printMenu() {
        System.out.println("\n Menu");
        System.out.println("======");
        System.out.println("0: Keluar");
        System.out.println("1: Membuat list baru berisi elemen acak");
        System.out.println("2: Urutkan list menggunakan selection sort");
        System.out.println("3: Mencari elemen di list menggunakan pencarian linear");
        System.out.println("4: Cetak list");
        System.out.print("\nMasukan angka menu yang dipilih: ");
    }

    // method helper untuk keperluan random number
    public static int[] randomize(int[] list) {
        // temporary array
        int[] limit = new int[list.length];
        // Generates `limit.length` Random Numbers in the range 1 - 1000
        for (int i = 0; i < limit.length; i++) {
            limit[i] = (int) (Math.random() * 1000 + 1);
        }
        return limit;
    }

    // menampilkan list/array, index -> value
    public static void printList(int[] list) {
        for (int i = 0; i < list.length; i++)
            System.out.println("[" + i + "]:\t" + list[i]);
    }

    public static int linearSearch(int[] list, int target) {
        return linearSearchRec(list, target, 0);
    }

    // implementasi liniar search recursion
    private static int linearSearchRec(int[] list, int target, int lo) {
        // jika 0 lebih besar dari ukuran array, return -1
        if (lo > list.length-1) {
            return -1;
        }

        // jika elemen pertama array sama dengan nilai yg ingin dicari, return lo/indeks element
        if (list[lo] == target) {
            return lo;
        }

        // rekursif
        return linearSearchRec(list, target, lo+1);
    }

    // method untuk selection sort
    public static void selectionSort(int[] list) {
        int minIndex;
        for (int i = 0; i < list.length - 1; i++) {
            // carilah elemen terkecil mulai dari indeks ke-i
            minIndex = i;
            for (int j = i + 1; j < list.length; j++) {
                // cek apakah list index ke j lebih kecil dari list dengan minIndex
                if (list[j] < list[minIndex]) {
                    // kalau iya, minIndex set dengan index / j
                    minIndex = j;
                }
            }
            // swapping
            int smallerNumber = list[minIndex];
            list[minIndex] = list[i];
            list[i] = smallerNumber;
        }
    }
}
