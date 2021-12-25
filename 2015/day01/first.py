f = open("input.txt")
str = f.read()
print(str.count('(')-str.count(')'))