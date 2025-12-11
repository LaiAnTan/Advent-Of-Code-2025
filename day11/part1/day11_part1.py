import sys


def main():
	filename =  sys.argv[1]

	deviceMap = {}

	with open(filename, 'r') as f:

		for line in f.readlines():
			
			line = line.rstrip('\n')


			s = line.split(': ')

			key = s[0]
			values = s[1].split(' ')

			deviceMap[key] = values
		
	queue = ["you"]
	paths = 0

	while len(queue) > 0:

		curr = queue.pop()

		if curr == "out":
			paths += 1
			continue
		
		queue += deviceMap[curr]

	print(f"Paths: {paths}")

if __name__ == "__main__":
	main()