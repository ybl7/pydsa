def two_sum(arr: list[int], tgt: int) -> list[int]:
    # We'll use a map to store the element as the key and the index as the value, setting the element to the key facilitates fast lookups of the index
    el_to_idx = {}

    # Python for loop only gives the value, we want the idx to so use enumerate
    for idx, x in enumerate(arr):
        comp = tgt - x

        if comp in el_to_idx:
            return idx, el_to_idx[comp]

        el_to_idx[x] = idx

def main():
    print(two_sum([2,7,11,15],9))
    print(two_sum([3,2,4],6))
    print(two_sum([3,3],6))

if __name__ == "__main__":
    main()

# Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
# You may assume that each input would have exactly one solution, and you may not use the same element twice.
# You can return the answer in any order. We can't assume that the array is ordered. 

# Here's the thing, for any given number e we are looking for the complement n = target - e exists in the array
# Now we need to be able to quickly identify if it exists in the array and if so - at which index, this lends itself to a hash table, if we iterate through the array once, we can have a fully populated hash table of all elements
# Then we can do a second pass over the array and for each element e we can quickly lookup if the complement in the hash table exists.

# Except, we don't even need to do a second pass. Why? Because t = a + b necessarily means t - a = b and t - b = a, that means that as long as a complimentary pair exists
# it really doesn't matter which order we find them in. So if we miss the compliment of say a, i.e. we populate our map with a but we haven't found b yet, when we do get to b - if it exists - we will already have a
# This is true for any complementary pair, we get two tries at finding the pair, and we will necessarily always have one of the pair in the map before we find the second

# O(n) time as we iterate the arr one time
# O(n) space as we instantiate a map of size n