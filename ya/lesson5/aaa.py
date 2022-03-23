n = int(input())

a1 = input()
a2 = input()
a3 = input()

x1 = list(map(int,a1.split()[1:]))
x2 = list(map(int,a2.split()[1:]))
x3 = list(map(int,a3.split()[1:]))

index = {}
for i in range(len(x3)):
    if x3[i] not in index:
        index[x3[i]] = i

def summ(x1,x2,x3,n,ind):
    for i in range(len(x1)):
        for j in range(len(x2)):
            if n - x1[i] - x2[j] in ind:
                return [i,j,ind[n - x1[i] - x2[j]]]

    return [-1]

print(*summ(x1,x2,x3,n,index))
