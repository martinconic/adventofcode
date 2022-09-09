f = open("input.txt")
dict = {"0,0":1}	
i=ii=0
j=jj=0
while True:
	c = f.read(1)
	# print(c)
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

	key1 = str(i)+','+str(j)	
	if not key1 in dict:
		dict[key1] = 1
	else:
		dict[key1] +=1

	# print(dict)	

	c = f.read(1)
	# print(c)
	if not c:
		break
	if c == '^':
		ii+=1
	if c == 'v':
		ii-=1
	if c == '>':
		jj+=1
	if c == '<':
		jj-=1	


	key2 = str(ii)+','+str(jj)	
	if not key2 in dict:
		dict[key2] = 1
	else:
		dict[key2] +=1

	# print(dict)
		
print(len(dict))		