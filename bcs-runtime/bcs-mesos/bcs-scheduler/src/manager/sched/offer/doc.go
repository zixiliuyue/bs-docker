

/*
Package offer provides mesos offers resources pool implements.

struct Offer is responsible for the managements of mesos's offers, including add, use and delete.

Offer need interface SchedManager's function to manage offers, so we need struct
SchedManager to new struct offer. Interface SchedManager have four functions, for
details please look interface.go

	para := &OfferPara{
		Sched: &SchedManager{},
	}

	offerPool := NewOfferPool(para)
	//...
	var mesosOffers []*mesos.Offer
	//...

	//add mesos's offers to offer pool
	offerPool.AddOffers(mesosOffers)

	//get the first offer
	offer := offerPool.GetFirstOffer()
	for{
		if offer == nil {
			break
		}

		// if offer is suitable, then use this offer
		ok := offerPool.UseOffer(offer)
		if ok {
			break
		}

		//else get the next offer, until you get the right one
		offer = offerPool.GetNextOffer(offer)
	}
	//...

*/
package offer
