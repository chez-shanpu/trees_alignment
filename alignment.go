package main

import (
	"github.com/chez-shanpu/repo2tree/model"
	"math"
)

func layerAlignmentDistanceTotal(sourceLayerRootNode *model.Node, targetLayerRootNode *model.Node) (sum float64) {
	sum = 0
	sNode := sourceLayerRootNode
	tNode := targetLayerRootNode
	sum += alignmentDistance(sNode, tNode)
	for sNode != nil || tNode != nil {
		if (sNode != nil && sNode.ChildNode != nil) || (tNode != nil && tNode.ChildNode != nil) {
			if sNode == nil {
				sum += layerAlignmentDistanceTotal(nil, tNode.ChildNode)
			} else if tNode == nil {
				sum += layerAlignmentDistanceTotal(sNode.ChildNode, nil)
			} else {
				sum += layerAlignmentDistanceTotal(sNode.ChildNode, tNode.ChildNode)
			}
		}
		if sNode != nil {
			sNode = sNode.NextNode
		}
		if tNode != nil {
			tNode = tNode.NextNode
		}
	}
	return
}

// Calculate the total alignment distance for that layer
func alignmentDistance(sourceLayerRootNode *model.Node, targetLayerRootNode *model.Node) float64 {
	var dist float64 = 0

	sourceLayerLength := layerLength(sourceLayerRootNode)
	targetLayerLength := layerLength(targetLayerRootNode)
	layerLengthGap := math.Abs(float64(sourceLayerLength - targetLayerLength))
	if sourceLayerLength < targetLayerLength {
		tmpNode := sourceLayerRootNode
		sourceLayerRootNode = targetLayerRootNode
		targetLayerRootNode = tmpNode
	}

	for lg := layerLengthGap; lg > 0; lg-- {
		dist += nodeDataSum(sourceLayerRootNode)
		sourceLayerRootNode = sourceLayerRootNode.NextNode
	}
	for remainLength := targetLayerLength; remainLength > 0; remainLength-- {
		dist += math.Abs(nodeDataSum(sourceLayerRootNode) - nodeDataSum(targetLayerRootNode))
		sourceLayerRootNode = sourceLayerRootNode.NextNode
		targetLayerRootNode = targetLayerRootNode.NextNode
	}
	return dist
}

func layerLength(leftmostNode *model.Node) (length int) {
	length = 0
	for node := leftmostNode; node != nil; node = node.NextNode {
		length++
	}
	return
}

func nodeDataSum(node *model.Node) (sum float64) {
	sum = 0
	for key, val := range node.Data {
		sum += math.Pow(val, math.Pow(10, float64(key)))
	}
	return
}
