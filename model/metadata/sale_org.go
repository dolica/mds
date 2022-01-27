package metadata

import "github.com/dolica/mds/global"

type SaleOrg struct {
	global.HTF_MODEL
	Name              string          `bson:"name" json:"name"`
	License           string          `bson:"license" json:"license"`
	Type              string          `bson:"type" json:"type"`
	Province          string          `bson:"province" json:"province"`
	City              string          `bson:"city" json:"city"`
	District          string          `bson:"district" json:"district"`
	Address           string          `bson:"address" json:"address"`
	LicenseCreateTime int             `bson:"license_create_time" json:"licenseCreateTime"`
	LicenseExpireTime int             `bson:"license_expire_time" json:"licenseExpireTime"`
	BusinessScope     string          `bson:"business_scope" json:"businessScope"`
	Location          global.Location `bson:"location" json:"location"`
}
