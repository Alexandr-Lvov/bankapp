package card

import (
	"bank/pkg/bank/types"
)



func Total(cards []types.Card) types.Money{
	sum:=types.Money(0)
	for _,	total := range cards{
		
		if total.Balance<0{
			continue
		}
		if !total.Active{
			continue
		}
	
		sum+=total.Balance
	}


	return sum
}

	




