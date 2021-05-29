#include <stdlib.h>
#include <iostream>
#include <string>
#include <algorithm>
using namespace std;

const double POP_MIE = 10000;
const double LEMONILO = 13000;
const double SUPER_BUBUR = 20000;
const double SEBLAK_ACHI = 17000;
const double AQUA = 5000;
const double FANTA = 8000;
const double GOOD_DAY = 10000;
const double TEH_BOTOL = 6000;

// struct Vending Machine
typedef struct VendingMachine{
    int noMenu;
    int jumlahMenu;
    double hargaMenu;
    double uangKembalian;
    double totalHarga;
} VENDING_MACHINE;

// struct Base
typedef struct Base{
    int noMenu;
    string namaMenu;
    double hargaMenu;
} BASE;

// array vending menu
const int sizeArrMenu = 4;
Base menuMinuman[sizeArrMenu];
Base menuMakanan[sizeArrMenu];

// array pesanan
VENDING_MACHINE pesanan;
string namaMenu;

void header();
void inputPesanan();
void kalkulasiPesanan();
void details();
void repeat();
void start();
void searchingMatchMakananAtauMinuman();
void normal();
void sortingBerdasarkanHargaASC();
void sortingBerdasarkanHargaDSC();

void isiMenu();

int main() {
    start();
    return 0;
}

void start() {
    isiMenu();

    int pilihan = 1;

    cout << "Pilihan Tampilan" << endl;
    cout << "1. Normal" << endl;
    cout << "2. Ascending berdasarkan harga menu" << endl;
    cout << "3. Descending berdasarkan harga menu" << endl;
    cout << "4. Cari Menu" << endl;
    cout << "   Pilihan anda : ";
    cin >> pilihan;


    switch (pilihan) {
        case 1:
            normal();
            break;
        case 2:
            sortingBerdasarkanHargaASC();
            break;
        case 3:
            sortingBerdasarkanHargaDSC();
            break;
        case 4:
            searchingMatchMakananAtauMinuman();
            break;
        default:
            normal();
    }

    cout << endl << endl;
    inputPesanan();
}

void header() {
    cout<<"=======================================================================" <<endl;
    cout<<"                 SELAMAT DATANG DI VENDING MECHINE CITA RASA           " <<endl;
    cout<<"                 silahkan memilih menu makanan, dan minuman            " <<endl;
    cout<<"=======================================================================" <<endl;
    cout << endl;
}

void searchingMatchMakananAtauMinuman() {
    cout<<"=======================================================================" <<endl;
    cout<<"                 SELAMAT DATANG DI VENDING MECHINE CITA RASA           " <<endl;
    cout<<"                                   cari menu                           " <<endl;
    cout<<"=======================================================================" <<endl;
    cout << endl;

    string matchName;

    cout << "Masukkan nama makanan/minuman yang ingin anda cari : ";
    cin >> matchName;

    cout << "===================================================================" << endl;
    cout << "MATCH WITH " + matchName << " |MINUMAN|" << endl;
    cout << "===================================================================" << endl;

    for(int x = 0; x < sizeArrMenu; x++){
        if (menuMinuman[x].namaMenu.find(matchName, 0) != string::npos){
            cout<<"     " << menuMinuman[x].noMenu << ". " << menuMinuman[x].namaMenu  << " Rp." << menuMinuman[x].hargaMenu << endl;
        }
    }

    cout << "===================================================================" << endl;
    cout << "MATCH WITH " + matchName << " |MAKANAN|" << endl;
    cout << "===================================================================" << endl;

    for(int x = 0; x < sizeArrMenu; x++){
        if (menuMakanan[x].namaMenu.find(matchName, 0) != string::npos){
            cout<<"     " << menuMakanan[x].noMenu << ". " << menuMakanan[x].namaMenu  << " Rp." << menuMakanan[x].hargaMenu << endl;
        }
    }
}

void normal() {
    header();

    cout<<">>   MAKANAN "<< endl;
    for (int i = 0; i < sizeArrMenu; ++i) {
        cout<<"     " << menuMakanan[i].noMenu << ". " << menuMakanan[i].namaMenu  << " Rp." << menuMakanan[i].hargaMenu << endl;
    }

    cout << endl;

    cout<<">>   MINUMAN "<< endl;
    for (int i = 0; i < sizeArrMenu; ++i) {
        cout<<"     " << menuMinuman[i].noMenu << ". " << menuMinuman[i].namaMenu  << " Rp." << menuMinuman[i].hargaMenu << endl;
    }
}

void sortingBerdasarkanHargaDSC() {
    header();

    cout << "===================================================================" << endl;
    cout << "SORTING MINUMAN DESC BY HARGA " << endl;
    cout << "===================================================================" << endl;

    double hargaMenu[sizeArrMenu];

    for(int i=0; i < sizeArrMenu; i++){
        hargaMenu[i] = menuMinuman[i].hargaMenu;
    }

    sort(hargaMenu, hargaMenu + sizeArrMenu, greater<double>());

    for (int i = 0; i < sizeArrMenu; i++){
        for (int j = 0; j < sizeArrMenu; j++){
            if (hargaMenu[i] == menuMinuman[j].hargaMenu){
                cout<<"     " << menuMinuman[j].noMenu << ". " << menuMinuman[j].namaMenu  << " Rp." << menuMinuman[j].hargaMenu << endl;
                break;
            }
        }
    }

    cout << endl << endl;

    cout << "===================================================================" << endl;
    cout << "SORTING MAKANAN DESC BY HARGA " << endl;
    cout << "===================================================================" << endl;

    double hargaMenu2[sizeArrMenu];

    for(int i=0; i < sizeArrMenu; i++){
        hargaMenu2[i] = menuMakanan[i].hargaMenu;
    }

    sort(hargaMenu2, hargaMenu2 + sizeArrMenu, greater<double>());

    for (int i = 0; i < sizeArrMenu; i++){
        for (int j = 0; j < sizeArrMenu; j++){
            if (hargaMenu2[i] == menuMakanan[j].hargaMenu){
                cout<<"     " << menuMakanan[j].noMenu << ". " << menuMakanan[j].namaMenu  << " Rp." << menuMakanan[j].hargaMenu << endl;
                break;
            }
        }
    }
}

void isiMenu() {
    menuMakanan[0].noMenu = 1;
    menuMakanan[0].namaMenu= "POP MIE";
    menuMakanan[0].hargaMenu= 10000;

    menuMakanan[1].noMenu = 2;
    menuMakanan[1].namaMenu= "LEMONILO";
    menuMakanan[1].hargaMenu= 13000;

    menuMakanan[2].noMenu = 3;
    menuMakanan[2].namaMenu= "SUPER BUBUR";
    menuMakanan[2].hargaMenu= 20000;

    menuMakanan[3].noMenu = 4;
    menuMakanan[3].namaMenu= "SEBLAK ACHI";
    menuMakanan[3].hargaMenu= 17000;

    menuMinuman[0].noMenu = 5;
    menuMinuman[0].namaMenu= "AQUA";
    menuMinuman[0].hargaMenu= 5000;

    menuMinuman[1].noMenu = 6;
    menuMinuman[1].namaMenu= "FANTA";
    menuMinuman[1].hargaMenu= 8000;

    menuMinuman[2].noMenu = 7;
    menuMinuman[2].namaMenu= "GOOD DAY";
    menuMinuman[2].hargaMenu= 10000;

    menuMinuman[3].noMenu = 8;
    menuMinuman[3].namaMenu= "TEH BOTOL";
    menuMinuman[3].hargaMenu= 6000;
}

void inputPesanan() {
    cout<<"====== input pesanan ======"<<endl;
    cout << "Masukkan nomor makanan atau minuman: ";
    cin >> pesanan.noMenu;

    cout << "Masukkan jumlah pesanan: ";
    cin >> pesanan.jumlahMenu;

    cout << endl;

    kalkulasiPesanan();
}

void kalkulasiPesanan() {
    switch (pesanan.noMenu) {
        case 1:
            pesanan.hargaMenu = POP_MIE;
            namaMenu = "POP MIE";
            break;
        case 2:
            pesanan.hargaMenu = LEMONILO;
            namaMenu = "LEMONILO";
            break;
        case 3:
            pesanan.hargaMenu = SUPER_BUBUR;
            namaMenu = "SUPER BUBUR";
            break;
        case 4:
            pesanan.hargaMenu = SEBLAK_ACHI;
            namaMenu = "SEBLAK ACHI";
            break;
        case 5:
            pesanan.hargaMenu = AQUA;
            namaMenu = "AQUA";
            break;
        case 6:
            pesanan.hargaMenu = FANTA;
            namaMenu = "FANTA";
            break;
        case 7:
            pesanan.hargaMenu = GOOD_DAY;
            namaMenu = "GOOD DAY";
            break;
        case 8:
            pesanan.hargaMenu = TEH_BOTOL;
            namaMenu = "THE BOTOL";
            break;
        default:
            pesanan.hargaMenu = 0;
            namaMenu = "TIDAK DIKETAHUI";
    }

    pesanan.totalHarga = pesanan.jumlahMenu * pesanan.hargaMenu;

    cout << "====== Detail Pesanan ======";
    cout << "\nNo Menu        : " << pesanan.noMenu << endl;
    cout << "Nama Menu      : " << namaMenu << endl;
    cout << "Jumlah pesanan : " << pesanan.jumlahMenu << endl;
    cout << "Harga menu     : Rp. " << pesanan.hargaMenu << endl;
    cout << "Total harga    : Rp. " << pesanan.totalHarga << endl << endl;

    details();
}

void details() {
    double uang;
    cout << "====== Pembayaran ======" << endl;
    cout << "Masukkan uang anda : Rp. ";
    cin >> uang;

    cout << endl << endl;

    if (uang >= pesanan.totalHarga) {
        if (uang == pesanan.totalHarga) {
            pesanan.uangKembalian = 0;
        }else {
            pesanan.uangKembalian = uang - pesanan.totalHarga;
        }

        cout << "====== Struk Pesanan ======";
        cout << "\nNo Menu        : " << pesanan.noMenu << endl;
        cout << "Nama Menu      : " << namaMenu << endl;
        cout << "Jumlah pesanan : " << pesanan.jumlahMenu << endl;
        cout << "Harga menu     : Rp. " << pesanan.hargaMenu << endl;
        cout << "Total harga    : Rp. " << pesanan.totalHarga << endl;
        cout << "Uang anda      : Rp. " << uang << endl;
        cout << "Uang kembalian : Rp. " << pesanan.uangKembalian << endl;
    }else {
        cout << "Uang anda tidak cukup"<< endl;
    }

    repeat();
}

void repeat() {
    char status;

    cout << "\n\n====== Ingin Beli Lagi ======" << endl;
    cout << "y/n : ";
    cin >> status;
    if (status == 'y') {
        start();
    }else if(status == 'n') {
        exit(0);
    }
}


void sortingBerdasarkanHargaASC() {
    header();

    cout << "===================================================================" << endl;
    cout << "SORTING MINUMAN ASC BY HARGA " << endl;
    cout << "===================================================================" << endl;

    double hargaMenu[sizeArrMenu];

    for(int i=0; i < sizeArrMenu; i++){
        hargaMenu[i] = menuMinuman[i].hargaMenu;
    }

    sort(hargaMenu, hargaMenu + sizeArrMenu);

    for (int i = 0; i < sizeArrMenu; i++){
        for (int j = 0; j < sizeArrMenu; j++){
            if (hargaMenu[i] == menuMinuman[j].hargaMenu){
                cout<<"     " << menuMinuman[j].noMenu << ". " << menuMinuman[j].namaMenu  << " Rp." << menuMinuman[j].hargaMenu << endl;
                break;
            }
        }
    }

    cout << endl << endl;

    cout << "===================================================================" << endl;
    cout << "SORTING MAKANAN ASC BY HARGA " << endl;
    cout << "===================================================================" << endl;
    double hargaMenu2[sizeArrMenu];
    for(int i=0; i < sizeArrMenu; i++){
        hargaMenu2[i] = menuMakanan[i].hargaMenu;
    }
    sort(hargaMenu2, hargaMenu2 + sizeArrMenu);

    for (int i = 0; i < sizeArrMenu; i++){
        for (int j = 0; j < sizeArrMenu; j++){
            if (hargaMenu2[i] == menuMakanan[j].hargaMenu){
                cout<<"     " << menuMakanan[j].noMenu << ". " << menuMakanan[j].namaMenu  << " Rp." << menuMakanan[j].hargaMenu << endl;
                break;
            }
        }
    }
}
