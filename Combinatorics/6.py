n = int(input())
count = 0
ans = []
for i in range(2 ** n):
    binary = bin(i)[2:]
    if len(binary) < n:
        binary = '0' * (n - len(binary)) + binary
    if binary.count("11") == 0:
        count += 1
        ans.append(binary)
print(count)
for i in ans:
    print(i)
