f = open("input.txt")
lines = f.readlines()

sum = 0
for line in lines:
	occurances = [each for each in line if each in ['a', 'e', 'i', 'o', 'u']]
	if len(occurances) < 3:
		continue
	if any(each in line for each in ["ab", "cd", "pq", "xy"]):
		continue
	for i in range(0, len(line)-1):
		if line[i] == line[i+1]:
			sum+=1
			break	
print(sum)