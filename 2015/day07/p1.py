f = open("input.txt")
lines = f.readlines()


data = {}
for line in lines:
	values = line.split(" -> ")
	expresion = values[0]

	if "AND" in expresion:
		expresion = expresion.replace("AND", "&")
	if "OR" in expresion:
		expresion = expresion.replace("OR", "|")
	if "LSHIFT" in expresion:
		expresion = expresion.replace("LSHIFT", "<<")
	if "RSHIFT" in expresion:
		expresion = expresion.replace("RSHIFT", ">>")
	if "NOT" in expresion:
		val = expresion.split()
		expresion = val[1] + " ^ 65535"


	data[values[1].strip()] = expresion

b = True
dv = ""

data["b"] = "16076"
print(data["b"])

while b:
	for i,j in data.items():
		values = j.split()

		if len(values) == 1 and j not in data:
			for value in data:
				temp = data[value].split() 
				if len(temp) == 3 and (i == temp[0] or i == temp[2]):
					dv = i
					data[value] = data[value].replace(i,j)
					v = data[value].split()
					
					if v[0] not in data and v[2] not in data:
						data[value] = str(eval(data[value]))
						if value == "a" or value == "lx":
							print(data["a"])
							print(data["lx"])

							b = False

	if dv in data:						
		data.pop(dv)				
# print(data)
