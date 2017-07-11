# Print substrings of given string & index pairs.

# IO ##################################
data = open(r"rosalind_ini3.txt", "rt").read() #readlines()
#######################################

# Show data
print("RAW DATA:\n",data)
data = data.split()
print("\nLIST of str:\n", data)

# Separate and parse ints
txt = data[0]
idx = list(map(int, data[1:]))
print("INDICES: ",idx)

ans = [txt[idx[x]:idx[x+1]+1] for x in range(0,len(idx),2)]
print("ANS: ",ans)
print(" ".join(ans))
#######################################
# Functional
def ss(txt,beg,end):
	return txt[int(beg):int(end)+1]
d = data    
ans = list(map(ss, d[:1]*2, *zip(d[1:3],d[3:])))
print(' '.join(ans))

txt = d[0]
ind = d[1:]
ans = list(map(ss, [txt]*2, *zip(ind[:3],d[3:])))
print(' '.join(ans))

##def subs(txt,beg,end):
##    return [txt,beg,end]#return(txt[beg:end+1])

##def subs(args):
##    txt = args[0]
##    s,e = args[1][0],args[1][1]
##    return txt[s:e+1]
##
##data = open(r"rosalind_ini3.txt", "rt").read().split()
##
##s = data[0]
##print(s)
##n = list(map(int, data[1:])) #[22, 27, 97, 102]
##print(n)
##p = [n[i:i+2] for i in range(0,len(data)-1,2)]
##print(p)
##z = zip([s]*2, p)
##a = list(z)
##print(a)
##print(list(map(subs,a)))




