f = open("input.txt")
dict = {"0,0":1}	
i=0
j=0
while True:
	c = f.read(1)
	if not c:
		break
	if c == '^':
		i+=1
	if c == 'v':
		i-=1
	if c == '>':
		j+=1
	if c == '<':
		j-=1

	key = str(i)+','+str(j)	
	if not key in dict:
		dict[key] = 1
	else:
		dict[key] +=1

print(len(dict))								 	