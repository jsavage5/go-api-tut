import time

def sum_of_squares(n):
    return sum(i * i for i in range(n))

if __name__ == "__main__":
    n = 100_000_000  # Adjust the range for larger computations
    start_time = time.time()
    result = sum_of_squares(n)
    end_time = time.time()
    
    print(f"Sum of squares: {result}")
    print(f"Time taken in Python: {end_time - start_time:.4f} seconds")
