f = open("input.txt")
lines = f.readlines()

sum = 0
for line in lines:
	nums = [int(n) for n in line.split('x')]
	nums.sort()
	ribbon = 2*nums[0] + 2*nums[1]
	ribbon += nums[0]*nums[1]*nums[2]
	sum +=ribbon

print(sum)
	