#include <iostream>
#include <vector>

using namespace std;

int main() {
    int y, x;
    cin >> y >> x;
    vector<vector<int>> coins(y, vector<int>(x));
    for (int i = 0; i < y; i++) {
        for (int j = 0; j < x; j++) {
            cin >> coins[i][j];
        }
    }
    vector<vector<int>> maxCoins(y, vector<int>(x));
    maxCoins[0][0] = coins[0][0];
    for (int i = 0; i < y; i++) {
        for (int j = 0; j < x; j++) {
            if (i == 0 && j == 0) {
                continue;
            }
            if (i == 0) {
                maxCoins[i][j] = maxCoins[i][j - 1] + coins[i][j];
            } else if (j == 0) {
                maxCoins[i][j] = maxCoins[i - 1][j] + coins[i][j];
            } else {
                maxCoins[i][j] = max(maxCoins[i - 1][j], maxCoins[i][j - 1]) + coins[i][j];
            }
        }
    }
    cout << maxCoins[y - 1][x - 1] << endl;
    vector<string> path;
    int i = y - 1;
    int j = x - 1;
    for (; i > 0 || j > 0;) {
        if (i == 0) {
            path.emplace_back("R");
            j--;
        } else if (j == 0) {
            path.emplace_back("D");
            i--;
        } else {
            if (maxCoins[i - 1][j] > maxCoins[i][j - 1]) {
                path.emplace_back("D");
                i--;
            } else {
                path.emplace_back("R");
                j--;
            }
        }
    }
    for (int i = path.size() - 1; i >= 0; i--) {
        cout << path[i];
    }
    return 0;
}