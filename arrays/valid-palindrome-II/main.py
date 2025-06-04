def valid_palindrome_ii(s):
    def valid_palindrome_rec(i, j, count):
        if count < 0:
            return False

        while i < j:
            if s[i].lower() != s[j].lower():
                return valid_palindrome_rec(i+1, j, count-1) or valid_palindrome_rec(i, j-1, count-1)

            # if we don't enter the if statement it means s[i] == s[j] so we can move onto the next characters
            i += 1
            j -= 1
         # if at the end of the while loop we still haven't returned false, then it means that this is a valid palindrome
        return True
    return valid_palindrome_rec(0, len(s)-1, 1)



def main():
    print(valid_palindrome_ii("madame"))    # True
    print(valid_palindrome_ii("dead"))      # True
    print(valid_palindrome_ii("abca"))      # True
    print(valid_palindrome_ii("tebbem"))    # False
    print(valid_palindrome_ii("eeccccbebaeeabebccceea"))    # False


if __name__ == "__main__":
    main()

# We can take a very similar approach to the original valid palindrome question, except we can tolerate one out of place character
# Suppose we have an palendrome like "ABCBA", now  let's state the valid cases
# "ABCBA" OR "XABCBA" OR "AXBCBA" OR "ABXCBA" OR "ABCXBA" OR "ABCBXA" OR "ABCBXA"
# Suppose we are at case "XABCBA" initially str[i] == X and str[j] == A, now at this point we can't tell if it's A or X that is wrong, we can't really assume that
# The best that we can do is is to run our function against "XABCB" and "ABCBA" recursively, 
# Since we already hit the first failure, we decrement the count variable
# Here is the clever part, we aren't actually checking for a palindrome like the first question (see valid-palindrome in this repo)
# We are checking for violation of palindrome more than once, which is an easier condition to check for recursively
# In the next recursive call "XABCB" will immediately return false since we hit another invalid case, but "ABCBA" will return true
# We need either of these to return true
# It's important to note that since we want to remove invalid characters, this question is also slightly different from the first one since we can't tolerate
# non-alphanumeric characters in this case, they actually matter
