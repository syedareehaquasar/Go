def predictOutput(s str) -> str:
    r = s.count("R")
    p = s.count("P")
    s = s.count("S")
    if r > max(p, s):
        return "R" * len(s)
    elif p > s:
        return "P" * len(s)
    else:
        return "S" * len(s)

tc = int(input())
while tc:
    s = input()
    predictOutput(s)
    tc -= 1