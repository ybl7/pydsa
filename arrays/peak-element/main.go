// We are told that we can take arr[-1] = -inf and arr[len(arr)] = +inf
// this means that relative to our "off the array" endpoints, we are guaranteed at least one peak
// unless we have a pathalogical case where every element is exactly the same
// Now a peak is any index i such that arr[i-1] < arr[i] < arr[i+1]
// The naive approach would be to iterate the array and look for an i that satisfies this adn would be O(n).
// But there is a more optimal way. Suppose we just pick a the midpoint m at random of the array.
// There are a few cases:
// case 1: arr[m-1] < arr[m] > arr[m+1], we've found a peak and can exit
// case 2: arr[m-1] < arr[m] < arr[m+1], we are guaranteed a peak in the interval arr[m] to arr[len(arr)] due to our -inf at the end
// case 3: arr[m-1] > arr[m] > arr[m+1], we are guaranteed a peak in the interval arr[0] to arr[m]
// Why are we guaranteed peaks, because even if there are no other peaks and the array is monotonically increasing in case 2
// and monotonically decreasing in case 3, arr[m] acts as a bound between it an -inf, guaranteeing that arr[m+1] (case 2) and
// arr[m-1] (case 3) are peaks. Even if the array keeps getting bigger to the right and left respectively until the endpoints
// the -inf off the arrays will make the endpoints of the arrays peaks.
// So the question just becomes a case of chasing ever increasing values and bounding them with either -inf, or a tighter
// bound that we find as we bisect the array
