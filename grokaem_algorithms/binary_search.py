def binary_search(array, elem):
    
    low = 0
    high = len(array) - 1

    while low <= high:
        mid = (low + high) // 2
        guess = array[mid]
        if guess == elem:
            return mid
        if guess > elem:
            high = mid - 1
        else:
            low = mid + 1
    return None

if __name__ == '__main__':
    array = [1, 2, 3, 4, 5, 7, 10, 77]

    print(binary_search(array, 77))
    print(binary_search(array, -1))
