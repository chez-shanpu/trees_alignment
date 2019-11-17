package main

import (
	"github.com/chez-shanpu/repo2tree/model"
	"math"
)

//TODO 片方がnilだったときの処理について考える
func alignmentDistance(sourceLayerNode *model.Node, targetLayerNode *model.Node) float64 {
	var res float64 = 0

	sourceLayerLength := layerLength(sourceLayerNode)
	targetLayerLength := layerLength(targetLayerNode)
	layerLengthGap := math.Abs(float64(sourceLayerLength - targetLayerLength))
	if sourceLayerLength < targetLayerLength {
		tmpNode := sourceLayerNode
		sourceLayerNode = targetLayerNode
		targetLayerNode = tmpNode
		//tmpLength := sourceLayerLength
		//sourceLayerLength = targetLayerLength
		//targetLayerLength = tmpLength
	}

	for lg := layerLengthGap; lg > 0; lg-- {
		res += nodeDataSum(sourceLayerNode)
		sourceLayerNode = sourceLayerNode.NextNode
	}
	for remainLength := targetLayerLength; remainLength > 0; remainLength-- {
		res += math.Abs(nodeDataSum(sourceLayerNode) - nodeDataSum(targetLayerNode))
		sourceLayerNode = sourceLayerNode.NextNode
		targetLayerNode = targetLayerNode.NextNode
	}
	return res
}

func layerLength(leftmostNode *model.Node) int {
	var length int

	for node := leftmostNode; node != nil; node = node.NextNode {
		length++
	}
	return length
}

func nodeDataSum(node *model.Node) (sum float64) {
	sum = 0
	for key, val := range node.Data {
		sum += math.Pow(val, math.Pow(10, float64(key)))
	}
	return
}
