line = input()
s = sorted(set(line))
alph = [""] * len(s)
out = [0] * len(line)
i = 0
for j in s:
    alph[i] = j
    i += 1
for i in range(len(line)):
    out[i] = alph.index(line[i]) + 1
    alph.pop(alph.index(line[i]))
    alph.insert(0, line[i])


print(*out)
