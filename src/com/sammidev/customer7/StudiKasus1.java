package com.sammidev.customer7;
import javax.swing.*;

/**
 * Created by sam on 16/03/21.
 */
public class StudiKasus1 {

    static String[] namaPesawatList = {"Airbus", "Garuda"};
    static String[] kelasList = {"A", "B"};
    static Double[] biayaTiketList = {2000000d,3000000d};

    public static void main(String[] args) {

        boolean isNext = true;
        while (isNext) {
            Double biayaTiket    = 0.0;
            String kodeTiket     = JOptionPane.showInputDialog("Masukkan Kode tiket", null);
            String namaPesawat   = JOptionPane.showInputDialog("Masukkan no pesawat\n 1. Airbus\n 2. Garuda\n", null);
            String kelas         = JOptionPane.showInputDialog("Masukkan jenis Kelas\n E. Ekonomi\n B. Bisnis", null);
            String keberangkatan = JOptionPane.showInputDialog("Masukkan Keberangkatan\nContoh: Senin, 20 april 2021)", null);
            String tujuan        = JOptionPane.showInputDialog("Masukkan Tujuan", null);

            if (namaPesawat.equalsIgnoreCase("1")) {
                biayaTiket = biayaTiketList[0];
            }else if (namaPesawat.equalsIgnoreCase("2")) {
                biayaTiket = biayaTiketList[1];
            }else {
                System.exit(0);
            }

            Pesawat pesawat = new Pesawat(kodeTiket,namaPesawat,kelas,keberangkatan,tujuan, biayaTiket);
            pesawat.PPN();
            pesawat.Diskon();
            pesawat.TotalBayar();

            Object[] options = {"Y", "T"};
            int n = JOptionPane.showOptionDialog(null,
                    "Mau coba lagi?",
                    "",
                    JOptionPane.YES_OPTION,
                    JOptionPane.NO_OPTION,
                    null,     //do not use a custom Icon
                    options,  //the titles of buttons
                    options[0]); //default button title

            if (n == 0) {
                isNext = true;
            }else if(n == 1){
                JOptionPane.showMessageDialog(null, "TERIMA KASIH");
                isNext = false;
            }else {
                JOptionPane.showMessageDialog(null, "SALAH INPUT");
                isNext = false;
            }
        }
    }
}

abstract class TiketPesawat {
    public String KodeTiket;
    public String NamaPesawat;
    public String Kelas;
    public String Keberangkatan;
    public String Tujuan;
    public Double BiayaTiket;

    public TiketPesawat(String kodeTiket,
                        String namaPesawat,
                        String kelas,
                        String keberangkatan,
                        String tujuan,
                        Double biayaTiket) {

        KodeTiket = kodeTiket;
        NamaPesawat = namaPesawat;
        Kelas = kelas;
        Keberangkatan = keberangkatan;
        Tujuan = tujuan;
        BiayaTiket = biayaTiket;
    }

    abstract public double PPN();
    abstract public double Diskon();
    abstract public void TotalBayar();
}

class Pesawat extends TiketPesawat{

    public Pesawat(String kodeTiket, String namaPesawat, String kelas, String keberangkatan, String tujuan, Double biayaTiket) {
        super(kodeTiket, namaPesawat, kelas, keberangkatan, tujuan, biayaTiket);
    }

    @Override
    public double PPN() {
        return BiayaTiket * 0.02;
    }

    @Override
    public double Diskon() {
        if (Kelas.equalsIgnoreCase("B")) {
            return  0.1 * BiayaTiket;
        }else if (Kelas.equalsIgnoreCase("E")) {
            return 0.05 * BiayaTiket;
        }
        return 0.0d;
    }

    @Override
    public void TotalBayar() {
        Double totalBayar =  BiayaTiket - PPN() - Diskon();
        JOptionPane.showMessageDialog(null, "Total Bayar : " + totalBayar);
    }
}