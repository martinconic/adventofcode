f = open("input.txt")
str = f.read()

i = 0
count = 0
for c in str:
	i+=1
	if c == '(':
		count+=1
	if c == ')':
		count-=1

	if count == -1:
		print(i)
		break		