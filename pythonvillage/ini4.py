#Problem
#Given: Two positive integers a and b (a<b<10000).
#Return: The sum of all odd integers from a through b, inclusively.
#Sample Dataset
#100 200
#Sample Output
#7500

# IO ##################################
file = r"rosalind_ini4.txt"
data = open(file, "rt").read().split()
#######################################

def odd(x):
    return x%2==1

idx = list(map(int, data))
num = [x for x in range(idx[0],idx[1]+1) if odd(x)]
print(sum(num))



