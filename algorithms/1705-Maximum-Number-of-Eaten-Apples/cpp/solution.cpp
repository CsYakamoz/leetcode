#include <algorithm>
#include <iostream>
#include <vector>

struct AppleAndExpiredDate {
    int apple;
    int expiredDate;
};

class Solution {
  public:
    std::vector<AppleAndExpiredDate> heap;

    int eatenApples(std::vector<int> &apples, std::vector<int> &days) {
        auto cmp = [](AppleAndExpiredDate &x, AppleAndExpiredDate &y) {
            return x.expiredDate > y.expiredDate;
        };

        int result = 0;

        for (int d = 0; d < apples.size() || heap.size() != 0; d++) {
            if (d < apples.size() && days[d] != 0) {
                heap.push_back(AppleAndExpiredDate{apples[d], d + days[d]});
                push_heap(heap.begin(), heap.end(), cmp);
            }

            while (heap.size() != 0 && heap.begin()->expiredDate <= d) {
                pop_heap(heap.begin(), heap.end(), cmp);
                heap.pop_back();
            }

            if (heap.size() != 0) {
                result++;

                if (--heap.begin()->apple == 0) {
                    pop_heap(heap.begin(), heap.end(), cmp);
                    heap.pop_back();
                }
            }
        }

        return result;
    }
};
