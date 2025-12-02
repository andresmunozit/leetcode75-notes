# Euclidean Algorithm

The Euclidean algorithm calculates the GCD (Greatest Common Divisor) of two numbers.

**Key insight:** `gcd(a, b) = gcd(b, a % b)`

Repeatedly replace the larger number with the remainder until one becomes zero. The other number
is the GCD.

**Example:** `gcd(48, 18)`
- `gcd(48, 18)` → `gcd(18, 48 % 18 = 12)` → `gcd(12, 18 % 12 = 6)` → `gcd(6, 12 % 6 = 0)` → **GCD = 6**
