package main

import "fmt"

type iClusterState interface {
	execute(*ClusterState)
	setNext(iClusterState)
}

type ClusterState struct {
	id              int
	writeInCluster  int
	readFromCluster int
	viewCluster     int
	allSet          int
	isFinalStep     bool
}

type WriteState struct {
	next iClusterState
}

func (w *WriteState) execute(c *ClusterState) {
	if c.writeInCluster == 0 {
		fmt.Println("Write in cluster")
		c.isFinalStep = true
		return
	}
	w.next.execute(c)
}

func (w *WriteState) setNext(next iClusterState) {
	w.next = next
}

type ReadState struct {
	next iClusterState
}

func (r *ReadState) execute(c *ClusterState) {
	if c.readFromCluster == 0 {
		fmt.Println("Read from cluster")
		c.isFinalStep = true
		return
	}
	r.next.execute(c)
}

func (r *ReadState) setNext(next iClusterState) {
	r.next = next
}

type ViewState struct {
	next iClusterState
}

func (view *ViewState) execute(c *ClusterState) {
	if c.viewCluster == 0 {
		fmt.Println("View cluster")
		c.isFinalStep = true
		return
	}
	view.next.execute(c)
}

func (view *ViewState) setNext(next iClusterState) {
	view.next = next
}

type AllSetState struct {
	next iClusterState
}

func (all *AllSetState) execute(c *ClusterState) {
	if c.allSet == 0 {
		fmt.Println("All set")
		c.isFinalStep = true
		return
	}
	all.next.execute(c)
}

func (all *AllSetState) setNext(next iClusterState) {
	all.next = next
}

func main() {
	cluster := &ClusterState{1, 0, 0, 0, 0, false}
	write := &WriteState{}
	read := &ReadState{}
	view := &ViewState{}
	all := &AllSetState{}

	write.setNext(read)
	read.setNext(view)
	view.setNext(all)

	write.execute(cluster)

	cluster = &ClusterState{1, 1, 0, 0, 0, false}
	write.execute(cluster)

	cluster = &ClusterState{1, 1, 1, 0, 0, false}
	write.execute(cluster)
}
