length = int(input())
for i in range(2 ** length):
    binary = bin(i)[2:]
    if len(binary) < length:
        binary = '0' * (length - len(binary)) + binary
    print(binary)