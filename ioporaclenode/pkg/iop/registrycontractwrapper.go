package iop

type RegistryContractWrapper struct {
	*RegistryContract
}

//func (r *RegistryContractWrapper) FindOracleNodes() ([]RegistryContractOracleNode, error) {
//	count, err := r.CountOracleNodes(nil)
//	if err != nil {
//		return nil, fmt.Errorf("count oracle nodes: %w", err)
//	}
//	nodeEntries := make([]RegistryContractOracleNode, count.Int64())
//	for i := int64(0); i < count.Int64(); i++ {
//		node, err := r.FindOracleNodeByIndex(nil, big.NewInt(i))
//		if err != nil {
//			return nil, fmt.Errorf("find oracle node by index %d: %w", i, err)
//		}
//		nodeEntries[i] = node
//	}
//	return nodeEntries, nil
//}
//
//func (r *RegistryContractWrapper) GetAllPk() ([][]byte ,error){
//	nodes, err := r.FindOracleNodes()
//	count, err := r.CountOracleNodes(nil)
//	if err != nil {
//		return nil, fmt.Errorf("count oracle nodes: %w", err)
//	}
//	if err != nil {
//		fmt.Errorf("find nodes: %w", err)
//	}
//	nodePks := make([][]byte,count.Int64())
//	for i:= int64(0) ; i< count.Int64();i++{
//		nodePks[i] = nodes[i].PubKey
//	}
//	return nodePks , nil
//}
