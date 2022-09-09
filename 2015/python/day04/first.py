import hashlib

key = "ckczppom"
n = 100000

result = ""
while True:
	s = key+str(n)
	result = hashlib.md5(s.encode())
	if result.hexdigest().startswith("000000"):
		break
	n+=1	

print(n)
print(result.hexdigest())	
