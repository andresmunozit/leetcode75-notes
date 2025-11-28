## Binary Search

Binary search is an efficient algorithm for finding a **target value** in a **sorted array**.  
Instead of checking each element one by one, it repeatedly splits the search range in half.

- **Setup**:  
  - Set `left` to the index of the first element (0).  
  - Set `right` to the index of the last element (`n - 1`).

- **Loop** (while `left <= right`):  
  - Compute the middle index:  
    \[
    mid = left + \frac{right - left}{2}
    \]
    (integer division; this avoids overflow compared to `(left + right) / 2`).  
  - Compare `nums[mid]` with the `target`:
    - If `nums[mid] == target`, return `mid`.  
    - If `nums[mid] < target`, discard the **left half** (move `left = mid + 1`).  
    - If `nums[mid] > target`, discard the **right half** (move `right = mid - 1`).

- **If the loop ends without a match**:  
  The target is not in the array.

**Requirements**
- The array must be **sorted** (typically non-decreasing order).

**Complexity**
- Time: \(O(\log n)\) — each step halves the remaining search space.  
- Space: \(O(1)\) for the iterative implementation.

**Examples**
- 34. Find First and Last Position of Element in Sorted Array
