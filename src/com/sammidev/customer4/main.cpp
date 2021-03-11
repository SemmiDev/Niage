#include <iostream>
using namespace std;
int main(){
    int s,luas,keliling;

    cout << "Masukan sisi persegi = ";
    cin >> s;
    
    luas = s*s;
    cout<<"Luas persegi adalah "<< luas << endl;
    
    keliling = 4 * s;
    cout << "Keliling persegi adalah "<< keliling << endl << endl;
    
    int panjang,lebar;
    cout << "Masukan panjang persegi panjang = ";
    cin >> panjang;
    
    cout << "Masukan lebar persegi panjang = ";
    cin >> lebar;
    
    int luasPersegiPanjang = panjang * lebar;
    int kelilingPersegiPanjang = 2 * (panjang + lebar);
    
    cout<<"Luas persegi panjang adalah "<< luasPersegiPanjang << endl;
    cout<<"Keliling persegi panjang adalah "<< kelilingPersegiPanjang << endl;
    
    
    return 0;
}
