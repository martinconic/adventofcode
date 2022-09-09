f = open("input.txt")
lines = f.readlines()

sum = 0

for line in lines:
	g1 = False
	g2 = False
	for i in range (0, len(line)-2):
		if line[i] == line[i+2]:
			g1 = True

		if len(line) > 3 and (line.count(line[i:i+2]) > 1):
			g2 = True
	
	if g1 and g1 == g2:
		sum+=1			

print(sum)