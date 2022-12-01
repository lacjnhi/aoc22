f = open("day1/data_day1.txt", "r")
lines = f.readlines()

res = 0
cur = 0
h = []

for l in lines:
    if len(l) > 1:
        cur += int(l)
    else:
        h.append(cur)
        res = max(res, cur)
        cur = 0

if cur != 0:
    h.append(cur)

print(res)
print(sorted(h, reverse=True))
print(sum([i for i in sorted(h, reverse=True)[:3]]))
