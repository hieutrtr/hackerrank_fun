def readInput():
    str = input()
    num = int(input().strip())
    return str , num

def findCharPositions(str, char):
    res = []
    for i, c in enumerate(str):
        if c == char:
            res.append(i)
    return res

if __name__ == '__main__':
    s, n = readInput()
    print("{}, {}\n".format(s, n))
    count = 0
    positions = findCharPositions(s, 'a')
    print("{}".format(positions))
    steps = int(n/len(s))
    mod = n%len(s)
    for pos in positions:
        count += steps
        if mod > pos:
            count += 1
    print("{}".format(int(count)))