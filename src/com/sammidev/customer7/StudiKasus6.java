package com.sammidev.customer7;


import javax.swing.*;

/**
 * Created by sam on 16/03/21.
 */
public class StudiKasus6 {
    public static void main(String[] args) {
        boolean isNext = true;
        while (isNext) {

            int noPlat, LamaSewa;
            Double BiayaSewa;
            String NamaRental, jenis, NamaPenyewa, NoKTP, Tujuan;

            noPlat = Integer.parseInt(JOptionPane.showInputDialog(null, "No plat"));
            jenis = JOptionPane.showInputDialog(null, "Jenis\nM =  MiniBUS\nB = Bus");
            NamaRental = JOptionPane.showInputDialog(null, "Nama rental");
            NamaPenyewa = JOptionPane.showInputDialog(null, "Nama penyewa");
            NoKTP = JOptionPane.showInputDialog(null, "No KTP");
            LamaSewa = Integer.parseInt(JOptionPane.showInputDialog(null, "Lama Sewa (per hari)"));
            BiayaSewa = Double.parseDouble(JOptionPane.showInputDialog(null, "Biaya Sewa"));
            Tujuan = JOptionPane.showInputDialog(null, "Tujuan");

            if (jenis.equalsIgnoreCase("M")) {
                Minibus minibus = new Minibus(noPlat, "Mini Bus", NamaRental, NamaPenyewa, NoKTP, LamaSewa, BiayaSewa, Tujuan);
                minibus.Total();

            } else if (jenis.equalsIgnoreCase("B")) {
                Bus bus = new Bus(noPlat, "Mini Bus", NamaRental, NamaPenyewa, NoKTP, LamaSewa, BiayaSewa, Tujuan);
                bus.Total();
            }else {
                isNext = true;
            }

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

class Kendaraan {
    private int noPlat;
    private String jenis;
    private String NamaRental;
    private String NamaPenyawai;
    private String NoKTP;
    private int LamaSewa;
    private double BiayaSewa;
    private String Tujuan;

    public Kendaraan(int noPlat,
                     String jenis,
                     String namaRental,
                     String namaPenyawai,
                     String noKTP,
                     int lamaSewa,
                     double biayaSewa,
                     String tujuan) {
        this.noPlat = noPlat;
        this.jenis = jenis;
        NamaRental = namaRental;
        NamaPenyawai = namaPenyawai;
        NoKTP = noKTP;
        LamaSewa = lamaSewa;
        BiayaSewa = biayaSewa;
        Tujuan = tujuan;
    }

    public int getNoPlat() {
        return noPlat;
    }

    public String getJenis() {
        return jenis;
    }

    public String getNamaRental() {
        return NamaRental;
    }

    public String getNamaPenyawai() {
        return NamaPenyawai;
    }

    public String getNoKTP() {
        return NoKTP;
    }

    public int getLamaSewa() {
        return LamaSewa;
    }

    public double getBiayaSewa() {
        return BiayaSewa;
    }

    public String getTujuan() {
        return Tujuan;
    }
}

class Minibus extends Kendaraan{
    public Minibus(int noPlat, String jenis, String namaRental, String namaPenyawai, String noKTP, int lamaSewa, double biayaSewa, String tujuan) {
        super(noPlat, jenis, namaRental, namaPenyawai, noKTP, lamaSewa, biayaSewa, tujuan);
    }

    public Double PotonganHarga() {
        if (getLamaSewa() >= 7) {
            return (0.1 * getLamaSewa()) * getBiayaSewa();
        }
        return (0.025 * getLamaSewa()) * getBiayaSewa();
    }
    public void Total() {
        double result = (getBiayaSewa() * getLamaSewa()) - PotonganHarga();
        JOptionPane.showMessageDialog(null, "Total : " + result);
    }
}

class Bus extends Kendaraan {
    public Bus(int noPlat, String jenis, String namaRental, String namaPenyawai, String noKTP, int lamaSewa, double biayaSewa, String tujuan) {
        super(noPlat, jenis, namaRental, namaPenyawai, noKTP, lamaSewa, biayaSewa, tujuan);
    }

    public Double PotonganHarga() {
        if (getLamaSewa() >= 7) {
            return (0.1 * getLamaSewa()) * getBiayaSewa();
        }
        return (0.025 * getLamaSewa()) * getBiayaSewa();
    }
    public void Total() {
        double result = (getBiayaSewa() * getLamaSewa()) - PotonganHarga();
        JOptionPane.showMessageDialog(null, "Total : " + result);}
}
