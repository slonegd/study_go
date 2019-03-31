// ну а тут, как бы я написал на плюсах без оглядки на го
// ассемблер шаблоны с++ https://godbolt.org/z/OZFZCe
#include "iostream"

namespace {

struct S  {
    int a;
    S(int a) : a{a} {}
    bool operator< (S& other) {
        return this->a < other.a;
    }
    void Double() {
        a *= 2;
    }
};

template<class T>
T& GetMax(T& _1, T& _2) {
    return _1 < _2 ? _2 : _1;
};

template<class T>
void FindMaxAndDouble (T& _1, T& _2) {
	GetMax(_1,_2).Double();
}

} // namespace {

int main() {
    auto a1 = S{1};
    auto a2 = S{2};

    auto maxCopy = GetMax(a1, a2);
    auto& maxPointer = GetMax(a1, a2);

    std::cout << "a2.a before maxPointer: " << a2.a << "\n";
    maxPointer.a = 10;
    maxCopy.a = 11;
    std::cout << "maxPointer.a: " << maxPointer.a << "\n";
    std::cout << "a2.a after maxPointer: " << a2.a << "\n";
    std::cout << "maxCopy.a: " << maxCopy.a << "\n";
    a1.Double();

    FindMaxAndDouble(a1,a2);
    std::cout << "FindMaxAndDouble: " << a2.a << "\n";
}