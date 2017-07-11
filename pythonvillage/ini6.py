#Problem: Count words in text.
#Given: A file containing at most 1000 lines.
#Return: A file containing all the even-numbered lines from the original file. Assume 1-based numbering of lines.

# IO ##################################
#file = r"rosalind_ini6.txt"
#data = open(file, "rt").read() #readlines()
#######################################
data = "This is a test\nand only a test"
words = data.split()
print(data)
#######################################
print("---------METHOD 1:")
d = {} #dict
for w in words:
    if w in d:
        d[w] += 1
    else:
        d[w] = 1

for key, value in d.items():
    print(key,value)
#######################################
print("---------METHOD 2:")
d = {} #dict
for w in words:
    d[w] = d.get(w,0) + 1

[print(k,v) for k,v in d.items()]
#######################################
print("---------METHOD 3:")

#[print(w, words.count(w)) for w in set(words)]
for w in set(words):
    print(w, words.count(w))
#######################################
print("---------METHOD 4:")
from collections import Counter
for k,v in Counter(words).items():
    print(k,v)

