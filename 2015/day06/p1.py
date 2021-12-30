f = open("input.txt")
lines = f.readlines()

count = 0
matrix = [[0 for x in range(1000)] for y in range(1000)]
for line in lines:
	values = line.split()
	if values[0] == "turn":
		pair1 = [int(n) for n in values[2].split(",")]
		pair2 = [int(n) for n in values[4].split(",")]
		if values[1] == "on":
			for i in range(pair1[0], pair2[0]+1):
				for j in range(pair1[1], pair2[1]+1):
					matrix[i][j] = 1
		else:			
			for i in range(pair1[0], pair2[0]+1):
					for j in range(pair1[1], pair2[1]+1):
						matrix[i][j] = 0			

	else:
		pair1 = [int(n) for n in values[1].split(",")]
		pair2 = [int(n) for n in values[3].split(",")]
		for i in range(pair1[0], pair2[0]+1):
					for j in range(pair1[1], pair2[1]+1):
						if matrix[i][j] == 1:
							matrix[i][j] = 0
						else:
							matrix[i][j] = 1	


for i in range(len(matrix)):
	for j in range(len(matrix[i])):
		if matrix[i][j] == 1:
			count+=1

print(count)