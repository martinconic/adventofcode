f = open("input.txt")
lines = f.readlines()

sum = 0
for line in lines:
	nums = [int(n) for n in line.split('x')]
	value = 2 * nums[0]* nums[1] + 2 * nums[1] * nums[2] + 2 * nums[0] * nums[2]
	nums.sort()
	value += nums[0]*nums[1]
	sum +=value

print(sum)
	



