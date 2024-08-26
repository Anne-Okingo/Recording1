package function

func Target(nums []int, target int)[]int{
	indice:= []int{}

	for i := 0; i < len(nums); i++{
		for j := 1; j < len(nums); j++{
			if nums[i] + nums[j] == target{
				indice = append(indice, i)
				indice = append(indice,j)
			}
		}
	}
	return indice
}