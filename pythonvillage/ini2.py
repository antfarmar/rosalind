# Given: Two positive integers a and b, each less than 1000.
# Return: The integer corresponding to the square of the hypotenuse of the right triangle whose legs have lengths a and b.

# IO ##################################
file = r"rosalind_ini2.txt"
data = open(file, "rt").read() #readlines()
#######################################

def hypSqr(a,b):
	return a**2 + b**2

print("RAW DATA: ",data)
data = data.split()
print("LIST of str: ", data)
nums = list(map(int,data))
print("LIST of int: ", nums)
tup = tuple(nums)
print("TUPL of int: ", tup)
#######################################
print(hypSqr(*map(int,data)))
print(hypSqr(*nums)) # *nums = unpack/flatten nums
print(hypSqr(*tup))
#######################################
print(sum(int(d)**2 for d in data))
print(sum(map(lambda x: int(x) ** 2, data)))
