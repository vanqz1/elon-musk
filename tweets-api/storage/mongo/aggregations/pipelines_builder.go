package aggregations

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IPipelinesBuilder interface {
	Build() []bson.D
	setProductStage()
	setGroupStage()
	setSortStage()
	setAddFieldsStage()
}

type pipelinesBuilder struct {
	groupBy        []string
	sortStage      bson.D
	productStage   bson.D
	groupStage     bson.D
	addFieldsStage bson.D
}

// NewBuilder creates new instance of pipelinesBuilder
func NewBuilder(groupBy []string) IPipelinesBuilder {
	return &pipelinesBuilder{groupBy: groupBy}
}

// Build builds pipelines for aggregations
func (b *pipelinesBuilder) Build() []bson.D {
	b.setProductStage()
	b.setGroupStage()
	b.setAddFieldsStage()
	b.setSortStage()

	return []bson.D{b.productStage, b.groupStage, b.addFieldsStage, b.sortStage}
}

// setProductStage sets product stage of pipeline
func (b *pipelinesBuilder) setProductStage() {
	projectStage := bson.D{{"$project",
		bson.D{
			bson.E{Key: "document", Value: "$$ROOT"},
			b.getProjectStageIndex("year"),
			b.getProjectStageIndex("month"),
			b.getProjectStageIndex("hour"),
		},
	}}

	b.productStage = projectStage
}

// setGroupStage sets group stage of pipeline
func (b *pipelinesBuilder) setGroupStage() {
	var groupIds bson.D

	for _, itemName := range b.groupBy {
		groupIds = append(groupIds, bson.E{Key: itemName, Value: fmt.Sprintf("$%s", itemName)})
	}

	groupStage := bson.D{{"$group",
		bson.D{
			{"_id",
				groupIds,
			},
			bson.E{Key: "totalLikes", Value: bson.D{{"$sum", "$document.likes_count"}}},
			bson.E{Key: "averageLikes", Value: bson.D{{"$avg", "$document.likes_count"}}},
			bson.E{Key: "totalRetweets", Value: bson.D{{"$sum", "$document.retweets_count"}}},
			bson.E{Key: "averageRetweets", Value: bson.D{{"$avg", "$document.retweets_count"}}},
			bson.E{Key: "totalReplies", Value: bson.D{{"$sum", "$document.replies_count"}}},
			bson.E{Key: "averageReplies", Value: bson.D{{"$avg", "$document.replies_count"}}},
			bson.E{Key: "tweetsCount", Value: bson.D{{"$sum", 1}}},
		},
	}}

	b.groupStage = groupStage
}

// setSortStage sets sort stage of pipeline
func (b *pipelinesBuilder) setSortStage() {
	var sortIndex bson.D
	for _, field := range b.groupBy {
		sortIndex = append(sortIndex, bson.E{Key: fmt.Sprintf("_id.%s", field), Value: 1})
	}

	sortStage := bson.D{{
		"$sort",
		sortIndex,
	}}

	b.sortStage = sortStage
}

// setAddFieldsStage sets AddFields stage of pipeline
func (b *pipelinesBuilder) setAddFieldsStage() {
	var addFieldsIndex bson.A

	for i, itemName := range b.groupBy {
		addFieldsIndex = append(addFieldsIndex, bson.D{{"$toString", fmt.Sprintf("$_id.%s", itemName)}})
		addFieldsIndex = append(addFieldsIndex, "-")
		addFieldsIndex = append(addFieldsIndex, itemName)

		if i != len(b.groupBy)-1 {
			addFieldsIndex = append(addFieldsIndex, "-")
		}
	}

	addFieldsStage := bson.D{{"$addFields",
		bson.D{{"label",
			bson.D{{
				"$concat",
				addFieldsIndex,
			}}},
		}}}

	b.addFieldsStage = addFieldsStage
}

// getProjectStageIndex construct bson for stage index from item
func (b *pipelinesBuilder) getProjectStageIndex(itemName string) bson.E {
	return bson.E{
		Key: itemName,
		Value: bson.D{{
			Key: fmt.Sprintf("$%s", itemName),
			Value: bson.D{{
				Key: "$add",
				Value: bson.A{
					primitive.DateTime(0),
					"$created_at",
				},
			}}},
		}}
}
