#include <iostream>
using namespace std;
int main(){
    int s,luas,keliling,pilihan;

    cout << "Pilih rumus tersedia:" << endl; 
    cout << "1. Persegi" << endl; 
    cout << "2. Persegi Panjang" << endl; 
    cout << "pilihan anda no = ";
    cin >> pilihan; 

    if(pilihan == 1){
       cout << "Masukan sisi persegi = ";
       cin >> s;
    
       luas = s*s;
       cout<<"Luas persegi adalah "<< luas << endl;
    
       keliling = 4 * s;
       cout << "Keliling persegi adalah "<< keliling << endl << endl;     
    }else if(pilihan == 2) {
       int panjang,lebar;
       cout << "Masukan panjang persegi panjang = ";
       cin >> panjang;
       
       cout << "Masukan lebar persegi panjang = ";
       cin >> lebar;
    
       int luasPersegiPanjang = panjang * lebar;
       int kelilingPersegiPanjang = 2 * (panjang + lebar);
    
       cout<<"Luas persegi panjang adalah "<< luasPersegiPanjang << endl;
       cout<<"Keliling persegi panjang adalah "<< kelilingPersegiPanjang << endl;
    

    }else {
       cout << "pilihan salah";
    }

    
    
    
    return 0;
}
