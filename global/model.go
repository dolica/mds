package global

import "go.mongodb.org/mongo-driver/bson/primitive"

type HTF_MODEL struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreateAt int                `bson:"create_at,omitempty" json:"createAt"`
	CreateBy string             `bson:"create_by,omitempty" json:"createBy"`
	ModifyAt int                `bson:"modify_at,omitempty" json:"modifyAt"`
	ModifyBy string             `bson:"modify_by,omitempty" json:"modifyBy"`
}

type Location struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}
