package testfactory

import (
	"github.com/bxcodec/faker"
	"github.com/zaihui/ent-factory/factories"
	"github.com/zaihui/ent-factory/gen/entschema"
	"github.com/zaihui/ent-factory/service/app"
)

type TestFieldSetter func(*entschema.Test)

func SetID(iDGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.ID = iDGen
	}
}

func SetUID(uIDGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.UID = uIDGen
	}
}

func SetName(nameGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Name = nameGen
	}
}

func SetEnableTakeaway(enableTakeawayGen bool) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.EnableTakeaway = enableTakeawayGen
	}
}

func SetPhone(phoneGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Phone = phoneGen
	}
}

func SetAnnouncement(announcementGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Announcement = announcementGen
	}
}

func SetEnableAutoAccept(enableAutoAcceptGen bool) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.EnableAutoAccept = enableAutoAcceptGen
	}
}

func SetPrintTimes(printTimesGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.PrintTimes = printTimesGen
	}
}

func SetBusinessGroupID(businessGroupIDGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.BusinessGroupID = businessGroupIDGen
	}
}

func SetBusinessGroupUID(businessGroupUIDGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.BusinessGroupUID = businessGroupUIDGen
	}
}

func SetDishCategories(dishCategoriesGen []string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.DishCategories = dishCategoriesGen
	}
}

func SetMode(modeGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Mode = modeGen
	}
}

func SetServerID(serverIDGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.ServerID = serverIDGen
	}
}

func SetImage(imageGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Image = imageGen
	}
}

func SetAddress(addressGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Address = addressGen
	}
}

func SetLatitude(latitudeGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Latitude = latitudeGen
	}
}

func SetLongitude(longitudeGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Longitude = longitudeGen
	}
}

func SetSort(sortGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Sort = sortGen
	}
}

func SetPayMode(payModeGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.PayMode = payModeGen
	}
}

func New(s factories.TestSuite, opts ...TestFieldSetter) *entschema.Test {
	data := entschema.Test{}
	s.NoError(faker.FakeData(&data))
	for _, opt := range opts {
		opt(&data)
	}
	return app.EntClient.Test.Create().
		SetID(data.ID).
		SetUID(data.UID).
		SetName(data.Name).
		SetEnableTakeaway(data.EnableTakeaway).
		SetPhone(data.Phone).
		SetAnnouncement(data.Announcement).
		SetEnableAutoAccept(data.EnableAutoAccept).
		SetPrintTimes(data.PrintTimes).
		SetBusinessGroupID(data.BusinessGroupID).
		SetBusinessGroupUID(data.BusinessGroupUID).
		SetDishCategories(data.DishCategories).
		SetMode(data.Mode).
		SetServerID(data.ServerID).
		SetImage(data.Image).
		SetAddress(data.Address).
		SetLatitude(data.Latitude).
		SetLongitude(data.Longitude).
		SetSort(data.Sort).
		SetPayMode(data.PayMode).
		SaveX(s.Context())
}
