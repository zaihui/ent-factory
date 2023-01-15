package testfactory

import (
	"time"

	"github.com/bxcodec/faker"
	"github.com/zaihui/ent-factory/factories"
	"github.com/zaihui/ent-factory/gen/entschema"
	"github.com/zaihui/ent-factory/service/app"
	"github.com/zaihui/ent-factory/spec/schema"
)

type TestFieldSetter func(*entschema.Test)

// SetID Function Optional func for ID.
func SetID(iDGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.ID = iDGen
	}
}

// SetUID Function Optional func for UID.
func SetUID(uIDGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.UID = uIDGen
	}
}

// SetCreatedAt Function Optional func for CreatedAt.
func SetCreatedAt(createdAtGen time.Time) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.CreatedAt = createdAtGen
	}
}

// SetUpdatedAt Function Optional func for UpdatedAt.
func SetUpdatedAt(updatedAtGen time.Time) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.UpdatedAt = updatedAtGen
	}
}

// SetDeactivatedAt Function Optional func for DeactivatedAt.
func SetDeactivatedAt(deactivatedAtGen *time.Time) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.DeactivatedAt = deactivatedAtGen
	}
}

// SetName Function Optional func for Name.
func SetName(nameGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Name = nameGen
	}
}

// SetEnableTakeaway Function Optional func for EnableTakeaway.
func SetEnableTakeaway(enableTakeawayGen bool) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.EnableTakeaway = enableTakeawayGen
	}
}

// SetPhone Function Optional func for Phone.
func SetPhone(phoneGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Phone = phoneGen
	}
}

// SetAnnouncement Function Optional func for Announcement.
func SetAnnouncement(announcementGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Announcement = announcementGen
	}
}

// SetEnableAutoAccept Function Optional func for EnableAutoAccept.
func SetEnableAutoAccept(enableAutoAcceptGen bool) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.EnableAutoAccept = enableAutoAcceptGen
	}
}

// SetDeliveryConfigs Function Optional func for DeliveryConfigs.
func SetDeliveryConfigs(deliveryConfigsGen schema.DeliveryConfig) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.DeliveryConfigs = deliveryConfigsGen
	}
}

// SetPayConfigs Function Optional func for PayConfigs.
func SetPayConfigs(payConfigsGen schema.PayConfig) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.PayConfigs = payConfigsGen
	}
}

// SetPrintTimes Function Optional func for PrintTimes.
func SetPrintTimes(printTimesGen *int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.PrintTimes = printTimesGen
	}
}

// SetRingConfigs Function Optional func for RingConfigs.
func SetRingConfigs(ringConfigsGen schema.RingConfig) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.RingConfigs = ringConfigsGen
	}
}

// SetBusinessGroupID Function Optional func for BusinessGroupID.
func SetBusinessGroupID(businessGroupIDGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.BusinessGroupID = businessGroupIDGen
	}
}

// SetBusinessGroupUID Function Optional func for BusinessGroupUID.
func SetBusinessGroupUID(businessGroupUIDGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.BusinessGroupUID = businessGroupUIDGen
	}
}

// SetDishCategories Function Optional func for DishCategories.
func SetDishCategories(dishCategoriesGen []string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.DishCategories = dishCategoriesGen
	}
}

// SetEndOfTakeaway Function Optional func for EndOfTakeaway.
func SetEndOfTakeaway(endOfTakeawayGen time.Time) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.EndOfTakeaway = endOfTakeawayGen
	}
}

// SetMode Function Optional func for Mode.
func SetMode(modeGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Mode = modeGen
	}
}

// SetSelfPickupConfigs Function Optional func for SelfPickupConfigs.
func SetSelfPickupConfigs(selfPickupConfigsGen schema.SelfPickupConfig) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.SelfPickupConfigs = selfPickupConfigsGen
	}
}

// SetServerID Function Optional func for ServerID.
func SetServerID(serverIDGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.ServerID = serverIDGen
	}
}

// SetImage Function Optional func for Image.
func SetImage(imageGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Image = imageGen
	}
}

// SetAddress Function Optional func for Address.
func SetAddress(addressGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Address = addressGen
	}
}

// SetLatitude Function Optional func for Latitude.
func SetLatitude(latitudeGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Latitude = latitudeGen
	}
}

// SetLongitude Function Optional func for Longitude.
func SetLongitude(longitudeGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Longitude = longitudeGen
	}
}

// SetBanners Function Optional func for Banners.
func SetBanners(bannersGen schema.Banners) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Banners = bannersGen
	}
}

// SetSort Function Optional func for Sort.
func SetSort(sortGen int) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.Sort = sortGen
	}
}

// SetPayMode Function Optional func for PayMode.
func SetPayMode(payModeGen string) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.PayMode = payModeGen
	}
}

// SetDineInConfigs Function Optional func for DineInConfigs.
func SetDineInConfigs(dineInConfigsGen schema.DineInConfigs) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.DineInConfigs = dineInConfigsGen
	}
}

// SetDealsConfig Function Optional func for DealsConfig.
func SetDealsConfig(dealsConfigGen schema.DealsConfig) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.DealsConfig = dealsConfigGen
	}
}

// SetPrintConfig Function Optional func for PrintConfig.
func SetPrintConfig(printConfigGen schema.PrintConfig) TestFieldSetter {
	return func(testGen *entschema.Test) {
		testGen.PrintConfig = printConfigGen
	}
}

// New function for creating one Test instance.
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
		SetNillablePrintTimes(data.PrintTimes).
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
