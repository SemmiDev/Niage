#include <iostream>
using namespace std;

int main(){
cout << "Program deret angka kubik \n";
    cout << "-------------------\n";
    
    int jumlah;
    cout << "Masukkan jumlah : ";
    cin >> jumlah;
    
    cout << endl;
    for (int i = 1; i <= jumlah; i++) {
        cout << i << " " << i*i*i <<"\n";
    }

    return 0;
}
