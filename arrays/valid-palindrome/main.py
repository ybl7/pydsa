def valid_palindrome(s):
    i, j = 0, len(s) - 1

    while i < j:
        # Increment i or j until we get an alpha numeric character, this needs to be a while loop since we want to increment/decrement
        # until we get to the next alphanumeric character, the only problem is that this might not be a palindrome and we increment/decrement i/j past the midpoint
        # so in this case we need to add the while check since we can easily violate the outer while loop in the process of incrementing/decrementing
        while not s[i].isalnum():
            i += 1
        while not s[j].isalnum():
            j -= 1
        # In this case both i and j should be alphanumeric, then we insist that they are the same
        if s[i].lower() != s[j].lower():
            return False
        # Increment in the case where they are both alphanumeric and were enqual
        i += 1
        j -= 1
    
    # If we haven't returned false and i is no longer less than j, it must mean we have a palindrome
    return True

def main():
    
    test_cases = [
        ("A man, a plan, a canal: Panama"),
        ("race a car"),
        ("1A@2!3 23!2@a1"),
        ("No 'x' in Nixon"),
        ("12321"),
    ]

    for i in test_cases:
        print("\tstring: ", i)
        result = valid_palindrome(i)
        print("\n\tResult:", result)
        print("-"*100)

if __name__ == "__main__":
    main()

# The idea here is to iterate from both the start and end of the array, skipping spaces and punctuation since we only care about characters
# i will be always left of j and we just need to compare if string[i] = string[j] each time we incriment i and decrement j