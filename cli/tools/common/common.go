package common

const (
    Coin            int64 = 1e8
)

//将amount切分为1,2,5的组合，这样在进行amount混淆的时候就能够方便获取相同额度的utxo
func decomAmount2Nature(amount int64, order int64) []int64 {
	res := make([]int64, 0)
	if order == 0 {
		return res
	}
	mul := amount / order
	switch mul {
	case 3:
		res = append(res, order)
		res = append(res, 2*order)
	case 4:
		res = append(res, 2*order)
		res = append(res, 2*order)
	case 6:
		res = append(res, 5*order)
		res = append(res, order)
	case 7:
		res = append(res, 5*order)
		res = append(res, 2*order)
	case 8:
		res = append(res, 5*order)
		res = append(res, 2*order)
		res = append(res, 1*order)
	case 9:
		res = append(res, 5*order)
		res = append(res, 2*order)
		res = append(res, 2*order)
	default:
		res = append(res, mul*order)
		return res
	}
	return res
}

// 62387455827 -> 455827 + 7000000 + 80000000 + 300000000 + 2000000000 + 60000000000, where 455827 <= dustThreshold
//res:[455827, 7000000, 80000000, 300000000, 2000000000, 60000000000]
func DecomposeAmount2digits(amount, dustThreshold int64) []int64 {
	res := make([]int64, 0)
	if 0 >= amount {
		return res
	}

	isDustHandled := false
	var dust int64
	var order int64 = 1
	var chunk int64

	for 0 != amount {
		chunk = (amount % 10) * order
		amount /= 10
		order *= 10
		if dust+chunk < dustThreshold {
			dust += chunk //累加小数，直到超过dust_threshold为止
		} else {
			if !isDustHandled && 0 != dust {
				//1st 正常情况下，先把dust保存下来
				res = append(res, dust)
				isDustHandled = true
			}
			if 0 != chunk {
				//2nd 然后依次将大的整数额度进行保存
				goodAmount := decomAmount2Nature(chunk, order/10)
				res = append(res, goodAmount...)
			}
		}
	}

	//如果需要被拆分的额度 < dustThreshold，则直接将其进行保存
	if !isDustHandled && 0 != dust {
		res = append(res, dust)
	}

	return res
}
