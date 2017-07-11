#Problem
#Given: A file containing at most 1000 lines.
#Return: A file containing all the even-numbered lines from the original file. Assume 1-based numbering of lines.

# IO ##################################
file = r"rosalind_ini5.txt"
data = open(file, "rt").read() #readlines()
#######################################

def even(x):
    return x%2==0

lines = data.split("\n")
evens = [lines[x] for x in range(len(lines)) if even(x+1)] 
print("\n".join(evens))


# f.close() is called when with block ends!
#[1::2] starting at line 1:to end:inc by 2
#with open('rosalind_ini5.txt','r') as f:
#    print(''.join(f.readlines()[1::2]))
