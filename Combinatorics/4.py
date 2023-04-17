n = int(input())
current = "0" * n
ans = {current}
while True:
    px = current[1:n + 1]
    if px + "1" not in ans:
        current = px + "1"
    elif px + "0" not in ans:
        current = px + "0"
    else:
        break
    ans.add(current)

for i in ans:
    print(i)
