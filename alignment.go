package main

import (
	"github.com/chez-shanpu/repo2tree/model"
	"math"
)

func alignmentDistance(sourceLayerNode *model.Node, targetLayerNode *model.Node) float64 {
	var dist float64 = 0

	sourceLayerLength := layerLength(sourceLayerNode)
	targetLayerLength := layerLength(targetLayerNode)
	layerLengthGap := math.Abs(float64(sourceLayerLength - targetLayerLength))
	if sourceLayerLength < targetLayerLength {
		tmpNode := sourceLayerNode
		sourceLayerNode = targetLayerNode
		targetLayerNode = tmpNode
	}

	for lg := layerLengthGap; lg > 0; lg-- {
		dist += nodeDataSum(sourceLayerNode)
		sourceLayerNode = sourceLayerNode.NextNode
	}
	for remainLength := targetLayerLength; remainLength > 0; remainLength-- {
		dist += math.Abs(nodeDataSum(sourceLayerNode) - nodeDataSum(targetLayerNode))
		sourceLayerNode = sourceLayerNode.NextNode
		targetLayerNode = targetLayerNode.NextNode
	}
	return dist
}

func layerLength(leftmostNode *model.Node) int {
	length := 0
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
