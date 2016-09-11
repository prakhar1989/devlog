---
title: "CPP 11 and 14 - key changes"
brief: "Enter a 2-3 liner here"
date: 2016-09-09
type: blog
---

If you've been on the Python bandwagon lately, you have likely missed out on the amazing updates that C++ has had in 2011 and 2014. Another update is just around the corner in 2017. While I gear up for a new specification release next year, here is a list quick review of the main features of C++11 and C++14.

## Commonly usable changes
These changes improve readability of the source code within a single function. For example, better strings, concise constructs, etc.

### Brief declarations
Iterators in C++98 were long and wordy. This changes to a short notation using the `auto` keyword.

```c++11
vector<int>::const_iterator it = ...;   // C++98
auto it = ...;                          // C++11
```

The type is inferred automatically from the return type on the right side.

### Vectors and arrays
Vectors in C++ have always had the `begin` and `end` member functions. Now they're no longer bound to a particular object and can operate on STL containers and arrays.

```c++11
// C++11
vector<int> v = {-5, 10, 30, -50, 300};
auto first3 = find(begin(v), end(v), 3);

// These also work on arrays
const char names[] = ["Bill", "Mark", "Eric", "Tim", "Jeff"];
auto first4 = find_if(begin(names), end(names), ...);
```

C++14 adds a bunch of new variants:

```c++11
// C++14
cbegin(v); cend(v);     // Const iterator
rbegin(v); rend(v);     // Reverse iterator
crbegin(v); crend(v);   // const reverse iterator
```

### Range based loops

C++11 introduces the idea of [range based for loops](http://stackoverflow.com/a/6963910/387099).

```c++11
// Convenient range based for loops
for(auto &e : names) {
    // ...
}

for(auto &elem : {1, 2, 3, 4, 5}) {
    // ...
}

// Does not work
for(auto e2 : iterator) {
}
```

While we're discussing iterators, there's a new keyword `nullptr` to denote null pointers. `NULL` is [part of the library](http://stackoverflow.com/a/924675/387099) and not a language feature. Use of 0 often caused confusion. The `nullptr` keyword makes the intentions clear.

### The >> issue
If you've worked with OpenCV, you're familiar with the >> issue.

```c++11
vector<vector<int> > v;      // C++98 Need to put an extra space between the two >
vector<vector<int>> v;       // C++11 - just works fine (like it should)
```

### Static asserts
Asserts are great for runtime checks but you [couldn't do compile time asserts](http://stackoverflow.com/a/6765787/387099). This is now possible with `static_assert`.

```c++11
assert(5 > 2);                  // Runtime assert

// Compile time assert
static_assert(sizeof(int) == 4, "Unsupported architecture");
```

The assertion is run during compilation. If the static assert fails, the compilation continues (and looks for other errors).

### Strings
C++11 comes with support for UTF strings. I'm sure some folks find this very useful.

```c++11
string s1 = u8"This is a UTF-8 string";
const char16_t *s2 = u"This is a UTF-16 string";
const char32_t *s3 = U"This is a UTF-32 string";

// Raw strings stard with R"( and end with ")
string raw_1 = R"(This is a very " \badly quoted string)";

// If you want to use the characters ") in your string, create
// a tag for: starts at R"xyz( and ends at )xyz"
string raw_2 = R"xyz(What if I wanted to use)xyz";
```

`long long` is now available. `alignas` and `alignof` are available. There are hooks into the garbage collector.

## Lambdas
With lambdas, you can write much cleaner code. The return type is inferred and the logic is closer to the usage.

```c++11
int main() {
    auto ... = find_if(begin(v), end(v), [](int n) { return n > 0; });
}
```

A lambda internally creates a new class (who's type you never find out unless you use `typeid`). Most of the code in this new class is optimized away but it leads to a problem. You do not have access to variables outside of the lambda.

```c++11
...
int x = 5;
int y = 10;
auto lambda1 = [](int n) { n + x; };     // x is undefined inside the lambda

auto lambda2 = [x](int n) { n + x; };    // x is passed in by value
auto lambda3 = [&x](int n) { ... };      // x is passed in by reference
auto lambda4 = [=](int n) { ... };       // All locals are passed (not data members)
auto lambda4 = [&](int n) { ... };       // All locals are passed by ref
auto lambda5 = [=, &y](int n) { ... };   // All passed by val, except y
auto lambda6 = [&, =y](int n) { ... };   // All passed by ref, except x
```

This is called [capturing variables](http://stackoverflow.com/a/7627218/387099) inside a lambda. You can only capture local variables - not data members of a class. You can, however, capture the `this` local variable.

C++98 does not allow you to define functions inside a function (you can define a class inside a function though). This restriction is somewhat removed with lambdas:

```c++11
auto lambda2 = [x](int n) { ... };    // x is passed in by value
lambda2(4);                           // lambda2 defined inside a function
```

C++14 introduces [generic lambdas](https://en.wikipedia.org/wiki/C%2B%2B14#Generic_lambdas) as well. This makes it templatized. I won't even dare to talk about this craziness.

```c++14
// C++14
[](const auto &p1, const auto &p2) { ... }
```

## Functions

### Inferring return type
You can use the `auto` keyword to infer the return type of a function. In C++11, you need to provide an operation on template types using `decltype`.

```c++11
// C++11
template<typename T, typename U>
auto product(const T &t, const U &u) -> decltype(t*u) {
    // ...
}
```

The `decltype` is required to infer the type of the product of T and U (which might be different for each pair of classes T and U). C++14 makes it even more brief:

```c++11
// C++14
template<typename T, typename U>
auto product(const T& t, const U& u) {
    return t*u;
}
```

Here, the type is inferred from the `return` statement directly. In case of multiple return statements, the compiler expects consistent return types.

> The `auto` keyword uses the template type deduction mechanism while the `decltype` keyword uses the (new) decltype type deduction mechanism.

### Attributes
You can decorate functions. `[[noreturn]]` allows you to flag functions that do not return back to the caller.

```c++11
[[noreturn]]
void something() {
    throw "error";
}
```

For a more complete list of attributes, look at the [GCC page](https://gcc.gnu.org/onlinedocs/gcc/Common-Function-Attributes.html#Common-Function-Attributes) or the [Clang page](http://clang.llvm.org/docs/AttributeReference.html).

### Dynamic exception specification
The `noexcept` operator tells the programmer and the compiler that a particular function will never throw an exception. Things like constructors, assignments, etc might benefit from this.

```c++11
class Base {
    Base() noexcept;
    Base(int i) noexcept;
}
```

This operator also enables the compiler to do specific optimizations - just like the `const` at the end of a function.

### Explicit override flagging
Virtual functions need to be overridden. But sometimes, users of your class may not know that.

```c++11
// C++11
class Base {    
    void sum(int) final;            // Can't override this anymore
    virtual int average();          // Can be overridden
    virtual void print() const;     // Can be overridden
}

class Derived : public Base {
    void sum(int);                  // Error!
    int average() override;         // Okay
    void print();                   // Compiled but whoops, no const.
                                    // Make sure you remember to use `override`
}
```

These keywords are contextual. They do their job only when they appear at the expected places. So your code containing the variable called override is going to be fine.

### Constructors and data-member initialization
This was a given in languages like Python but not in C++98.

```c++11
// C++11
class Animal {
    Animal() : Animal(4) { }
    Animal(int legs) : Animal(legs, 0) { }
    Animal(int l, int h) : legs(l), hands(h) { validate(); }

private:
    int legs = 0;       // Class level defaults
    int hands = 0;
}
```

During code generation, the constructor is given priority to assign default values. If the constructor does not assign a variable, the class-level defaults are used.

It is possible to inherit constructors into a class as well.

```c++11
// C++11
class Cat : public Animal {
    using Animal::Animal;       // Was illegal until C++11
                                // Now you have constructors
                                // Cat(), Cat(4), Cat(4, 0)
}
```

The new standard also allows you to initialize data-members (even ones that aren't the usual `int` based):

```c++11
class Base {
    int id;
    complex<double> count = 1000;
    int a[9] = {1, 2, 3, 4, 5};
}
```

You can initialize an array in the class declaration but you'll need to [specify the size of the array](http://www.stroustrup.com/C++11FAQ.html#member-init).



### Explicit conversion operator
Sometimes, the compiler automatically does type conversion if possible. This can cause hard-to-find bugs in code. The `explicit` keyword forces an explicit cast.

```c++98
// C++98
class Fraction {
    operator double() const;    // The compiler can sometimes convert it implicitly
}

Fraction f;
double d = f;                   // Implicit conversion
```

```c++11
// C++11
class Fraction {
    explicit operator double() const;   // Requires (double)fraction to work.
}

Fraction f;
double d = f;                           // Fails
double dc = (double)f;                  // Compiles!
```

## What's next?
I still haven't talked about a bunch of things - templates, move semantics, RTTI and pointers. Maybe next time?
