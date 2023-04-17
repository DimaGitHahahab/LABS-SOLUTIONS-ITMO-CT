def inverse(str_binary):
    result = ""
    for symbol in str_binary:
        if symbol == "0":
            result += "1"
        else:
            result += "0"
    return result


def triplesys(value):
    result = ""
    while value > 0:
        result += str(value % 3)
        value //= 3
    result = result[::-1]
    return result


def move(str_binary):
    result = ""
    for symbol in str_binary:
        if symbol == "0":
            result += "1"
        elif symbol == "1":
            result += "2"
        elif symbol == "2":
            result += "0"
        else:
            print("Error")
    return result


length = int(input())
for i in range(3 ** (length - 1)):
    triple = triplesys(i)
    if len(triple) < length:
        triple = '0' * (length - len(triple)) + triple
    print(triple)
    for i in range(2):
        triple = move(triple)
        print(triple)
