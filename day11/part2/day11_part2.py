import sys

def dfs_aux(u, pass_fft, pass_dac, memo, deviceMap):

	state = (u, pass_fft, pass_dac)
	if state in memo:
		return memo[state]

	if u == "out":
		return 1 if pass_fft and pass_dac else 0
	
	if u == "fft":
		pass_fft = True
	elif u == "dac":
		pass_dac = True

	paths = 0
	for v in deviceMap.get(u, []):
		paths += dfs_aux(v, pass_fft, pass_dac, memo, deviceMap)
	
	memo[state] = paths
	return paths

def numPaths(deviceMap):

	memo = {}

	return dfs_aux("svr", False, False, memo, deviceMap)


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
		
	paths = numPaths(deviceMap)

	print(f"Paths from svr to out passing through fft / dac: {paths}")

if __name__ == "__main__":
	main()