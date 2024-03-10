package main

import (
	"bird/vendors"
	"fmt"
	"strconv"
)

/*

A customer wants to send a single message to a list of phone numbers. Our system uses various suppliers to deliver those messages, but each supplier supports a limited set of prefixes. For this exercise, the prefix is a plus sign followed by 2 digits. For reliability purposes we have some overlap in prefix coverage.

Implement a system that simulates sending messages to the given numbers. Each number must go through the most suitable supplier at that time. It must optimise for the best cost for the customer. It must return the total messages sent and the total cost for the batch.

We start with the following suppliers, but assume for this list to be extended:


TweetMobile supports all prefixes at € 0,04 per message.
Swanizon supports +32 and +33 at € 0,02 and € 0,03 per message.
BeakOn supports +31 and +32, both at € 0,01 per message.

The batch is as follows:

batch := []string {
	"+31000000001",
	"+32000000001",
	"+33000000001",
	"+32000000002",
	"+3@00000003",
	"+34000000001",
	"+44000000001",
	"+31000000002",
	"+44000000002",
	"+31000000003",
}

$batch = [
        "+31000000001",
	"+32000000001",
	"+33000000001",
	"+32000000002",
	"+3@00000003",
	"+34000000001",
	"+44000000001",
	"+31000000002",
	"+44000000002",
	"+31000000003"
];

Expected outcome is 9 messages sent, total cost being 20 cents.

*/

func main() {

	vendors.InitializeVendorInfo()
	batch := []string{
		"+31000000001",
		"+32000000001",
		"+33000000001",
		"+32000000002",
		"+3@00000003",
		"+34000000001",
		"+44000000001",
		"+31000000002",
		"+44000000002",
		"+31000000003",
	}
	newList := make([]string, 0)
	totalCost := 0.0
	for _, number := range batch {
		regionString := "" + string(number[1]) + string(number[2])
		region, err := strconv.Atoi(regionString)
		if err != nil {
			continue
		} else {
			lowestCostVendor, ok := vendors.VendorInfoMap[region]
			if ok {
				if len(lowestCostVendor) == 0 {
					continue
				}
				lowestCost := lowestCostVendor[0].Price
				totalCost += lowestCost
				newList = append(newList, number)
			} else {
				allVendorCheck, ok1 := vendors.VendorInfoMap[0]
				if ok1 {
					lowestCOst := allVendorCheck[0].Price
					totalCost += lowestCOst
					newList = append(newList, number)
				}
			}
		}
	}
	fmt.Println(totalCost*100, " Cents")
	fmt.Println(newList)
}
