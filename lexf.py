import itertools as it
# IO ##################################
file = r"rosalind_lexf.txt"
data = open(file, "rt").readlines()
#######################################
dna = data[0].strip().replace(' ','')
n = int(data[1])

p = it.product(dna,repeat=n)
q = list(map(''.join, p))
print('\n'.join(q))


##def index(s):
##    indstr = ''
##    for c in s:
##        indstr += str(dna.index(c))
##
##    #print(indstr)    
##    return int(indstr)
##
##q.sort()
##print(q)
##print(sorted(q,key=index))
##

        

