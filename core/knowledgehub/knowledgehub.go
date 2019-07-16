package knowledgehub

// knowledge hub is the mecanism used by Etherniti to speed up I/O operations
// it works as a persistency layer based on cached results, so that
// future requests are resolved first using this mecanism and if no response is found
// they are leveraged on proper blockchain node
type KnowledgeHub struct {
}
