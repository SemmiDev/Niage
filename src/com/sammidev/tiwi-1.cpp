#include <iostream>
using namespace std;

int main(){
    cout << "Program deret angka kuadrat\n";
    cout << "-------------------\n";
    
    int jumlah;
    cout << "Masukkan jumlah : ";
    cin >> jumlah;
    
    cout << endl;
    for (int i = 0; i < jumlah; i++) {
        cout << i*i <<" ";
    }
    cout << endl;

    return 0;
}
