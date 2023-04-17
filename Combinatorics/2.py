value = int(input())

for i in range(2 ** value):
    binary = i ^ (i // 2)
    print('0' * (value - len(bin(binary)[2:])) + bin(binary)[2:])
