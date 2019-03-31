// переписал туже прогу на плюсах
// , хотя тут и не нужен такой подход, как на го, чисто для сравнения
// ассемблер го  https://go.godbolt.org/z/QZh59R
// ассемблер с++ https://godbolt.org/z/piyf0o
#include "iostream"

namespace {

struct Comparable {
    virtual bool Less(Comparable&) = 0;
};

struct Doubler {
    virtual void Double() = 0;
};

struct ComparableAndDoubler : Comparable, Doubler {};

struct S : ComparableAndDoubler {
    int a;
    S(int a) : a{a} {}
    bool Less(Comparable& other) override {
        return this->a < dynamic_cast<S&>(other).a; 
    }
    void Double() override {
        a *= 2;
    }
};

Comparable& GetMax(Comparable& _1, Comparable& _2) {
    return _1.Less(_2) ? _2 : _1;
};

void FindMaxAndDouble (ComparableAndDoubler& _1, ComparableAndDoubler& _2) {
	dynamic_cast<ComparableAndDoubler&>(GetMax(_1,_2)).Double();
}

} // namespace {

int main() {
    auto a1 = S{1};
    auto a2 = S{2};

    auto maxCopy = dynamic_cast<S&>(GetMax(a1, a2));
    auto& maxPointer = dynamic_cast<S&>(GetMax(a1, a2));

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