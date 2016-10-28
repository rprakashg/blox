package e2einstancesteps

import (
	"github.com/aws/amazon-ecs-event-stream-handler/internal/features/wrappers"
	. "github.com/gucumber/gucumber"
)

func init() {

	eshWrapper := wrappers.NewESHWrapper()

	When(`^I list instances$`, func() {
		eshInstances, err := eshWrapper.ListInstances()
		if err != nil {
			T.Errorf(err.Error())
		}
		for _, i := range eshInstances {
			eshContainerInstanceList = append(eshContainerInstanceList, *i)
		}
	})

	Then(`^the list instances response contains all the registered instances$`, func() {
		// eshContainerInstanceList can have instances from other clusters too
		if len(ecsContainerInstanceList) < len(eshContainerInstanceList) {
			T.Errorf("Unexpected number of instances in the list instances response")
		}
		for _, ecsInstance := range ecsContainerInstanceList {
			err := ValidateListContainsInstance(ecsInstance, eshContainerInstanceList)
			if err != nil {
				T.Errorf(err.Error())
			}
		}
	})
}