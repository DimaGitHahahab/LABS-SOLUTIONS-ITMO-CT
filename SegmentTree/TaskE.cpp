#include <iostream>
#include <vector>

using namespace std;

class SegmentTree {
public:
    vector<int> tree;

    void build(int x, int xl, int xr, vector<int> &arr) {
        if (xl == xr - 1) {
            tree[x] = arr[xl];
            return;
        }
        int xm = (xl + xr) / 2;
        build(2 * x + 1, xl, xm, arr);
        build(2 * x + 2, xm, xr, arr);
        tree[x] = tree[2 * x + 1] + tree[2 * x + 2];
    }

    void set(int i, int v, int x, int xl, int xr) {
        if (xl == xr - 1) {
            tree[x] = v;
            return;
        }
        int xm = (xl + xr) / 2;
        if (i < xm) {
            set(i, v, 2 * x + 1, xl, xm);
        } else {
            set(i, v, 2 * x + 2, xm, xr);
        }
        tree[x] = tree[2 * x + 1] + tree[2 * x + 2];
    }

    int sum(int l, int r, int x, int xl, int xr) {
        if (xl >= r || l >= xr) {
            return 0;
        }
        if (xl >= l && xr <= r) {
            return tree[x];
        }
        int xm = (xl + xr) / 2;
        int sl = sum(l, r, 2 * x + 1, xl, xm);
        int sr = sum(l, r, 2 * x + 2, xm, xr);
        return sl + sr;
    }
};

int main() {
    int n, inp;
    cin >> n;
    vector<int> data1(n);
    vector<int> data2(n);
    for (int i = 0; i < n; i++) {
        cin >> inp;
        if (i % 2 == 0) {
            data1[i] = inp;
            data2[i] = -1 * inp;
        } else {
            data1[i] = -1 * inp;
            data2[i] = inp;
        }
    }
    SegmentTree st1;
    st1.tree.resize(4 * n);
    SegmentTree st2;
    st2.tree.resize(4 * n);
    st1.build(0, 0, n, data1);
    st2.build(0, 0, n, data2);
    int m, cmd, x, y;
    cin >> m;
    for (int i = 0; i < m; i++) {
        cin >> cmd >> x >> y;
        x--;
        switch (cmd) {
            case 0:
                if (x % 2 == 0) {
                    st1.set(x, y, 0, 0, n);
                    st2.set(x, -1 * y, 0, 0, n);
                } else {
                    st1.set(x, -1 * y, 0, 0, n);
                    st2.set(x, y, 0, 0, n);
                }
                break;
            case 1:
                if (x % 2 == 0) {
                    cout << st1.sum(x, y, 0, 0, n) << endl;
                } else {
                    cout << st2.sum(x, y, 0, 0, n) << endl;
                }
                break;
        }
    }
}