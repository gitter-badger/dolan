package main

// IPTables interface for interacting with an iptables library. Declare it this
// way so that it is easy to dependency inject a mock.
type IPTables interface {
	ClearChain(string, string) error
	Append(string, string, ...string) error
	AppendUnique(string, string, ...string) error
	NewChain(string, string) error
}

// Setup creates a new iptables chain for holding peers and adds the chain and
// deny rules to the specified interface
func Setup(ipt IPTables, ipFace string) error {
	var err error

	err = ipt.NewChain("filter", "dolan-peers")
	if err != nil {
		if err.Error() != "exit status 1: iptables: Chain already exists.\n" {
			return err
		}
	}

	err = ipt.AppendUnique("filter", "INPUT", "-i", ipFace, "-j", "dolan-peers")
	if err != nil {
		return err
	}
	err = ipt.AppendUnique("filter", "INPUT", "-i", ipFace, "-j", "DROP")
	if err != nil {
		return err
	}
	return nil
}

// UpdatePeers updates the dolan-peers chain in iptables with the specified
// peers
func UpdatePeers(ipt IPTables, peers []string) error {
	// TODO(tam7t): prune `dolan-peers` in a way that doesnt cause downtime
	err := ipt.ClearChain("filter", "dolan-peers")
	if err != nil {
		return err
	}

	for _, peer := range peers {
		err := ipt.Append("filter", "dolan-peers", "-s", peer, "-j", "ACCEPT")
		if err != nil {
			return err
		}
	}
	return nil
}
