package problem

func TwoSum(nums []int, target int) (res []int){

	for key, value := range nums {
		for k, v := range nums{
			if k == key {
				continue
			}
			if value + v == target {
				res = append(res, key)
				res = append(res, k)
				return
			}
		}
	}
	return
}
