n, k = [int(x) for x in input().split()]
tempN = n
tempK = k
while tempN != tempK:
    if tempN > tempK:
        tempN -= tempK
    else:
        tempK -= tempN
print(n * k // tempN)
