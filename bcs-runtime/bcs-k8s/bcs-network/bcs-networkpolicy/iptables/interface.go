

package iptables

//Interface define iptables operation
type Interface interface {
	Exists(table, chain string, rulespec ...string) (bool, error)
	Insert(table, chain string, pos int, rulespec ...string) error

	//InsertUnique(table string, chain string, pos int, args ...string) error

	Append(table, chain string, rulespec ...string) error
	AppendUnique(table, chain string, rulespec ...string) error
	Delete(table, chain string, rulespec ...string) error
	List(table, chain string) ([]string, error)
	ListWithCounters(table, chain string) ([]string, error)
	ListChains(table string) ([]string, error)
	Stats(table, chain string) ([][]string, error)
	NewChain(table, chain string) error
	ClearChain(table, chain string) error
	RenameChain(table, oldChain, newChain string) error
	DeleteChain(table, chain string) error
	ChangePolicy(table, chain, target string) error
	HasRandomFully() bool
	GetIptablesVersion() (int, int, int)
}
