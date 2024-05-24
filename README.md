# Generate Password
This Go program generates a random password based on a given pattern. 

## breakdown of the code:
1. Inside `main()`, the program seeds the random number generator using the current time `rand.Seed(time.Now().UnixNano())`.

2. Three anonymous functions are defined:
- `c()` returns a random lowercase letter (excluding 'i', 'l', 'o', and 'q') as a string.
- `d()` returns a random digit (2, 3, 4, 5, 6, 7, 8, or 9) as a string.
- `cc(s string)` takes a string s as input and replaces all occurrences of 'c' with a random lowercase letter and all occurrences of 'd' with a random digit using regular expressions. It returns the modified string.

The pattern should consist of `'c'` and `'d'` characters, where `'c'` represents a lowercase letter (excluding 'i', 'l', 'o', and 'q') 
and `'d'` represents a digit (2, 3, 4, 5, 6, 7, 8, or 9). If no argument is provided, the default pattern `"dddcccdddcccddd"` is used.
For example, running the program with the argument `"cccdddcccddd"` will generate a password like `"bpv2345hkm6789"`.