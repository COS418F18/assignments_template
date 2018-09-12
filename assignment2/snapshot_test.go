package chandy_lamport

import (
	"fmt"
	"math/rand"
	"testing"
)

func runTest(t *testing.T, topFile string, eventsFile string, snapFiles []string) {
	startMessage := fmt.Sprintf("Running test '%v', '%v'", topFile, eventsFile)
	if debug {
		bars := "=================================================================="
		startMessage = fmt.Sprintf("%v\n%v\n%v\n", bars, startMessage, bars)
	}
	fmt.Println(startMessage)

	// Initialize simulator
	rand.Seed(8053172852482175524)
	sim := NewSimulator()
	readTopology(topFile, sim)
	actualSnaps := injectEvents(eventsFile, sim)
	if len(actualSnaps) != len(snapFiles) {
		t.Fatalf("Expected %v snapshot(s), got %v\n", len(snapFiles), len(actualSnaps))
	}
	// Optionally print events for debugging
	if debug {
		sim.logger.PrettyPrint()
		fmt.Println()
	}
	// Verify that the number of tokens are preserved in the snapshots
	checkTokens(sim, actualSnaps)
	// Verify against golden files
	expectedSnaps := make([]*SnapshotState, 0)
	for _, snapFile := range snapFiles {
		expectedSnaps = append(expectedSnaps, readSnapshot(snapFile))
	}
	sortSnapshots(actualSnaps)
	sortSnapshots(expectedSnaps)
	for i := 0; i < len(actualSnaps); i++ {
		assertEqual(expectedSnaps[i], actualSnaps[i])
	}
}

func Test2NodesSimple(t *testing.T) {
	runTest(t, "2nodes.top", "2nodes-simple.events", []string{"2nodes-simple.snap"})
}

func Test2NodesSingleMessage(t *testing.T) {
	runTest(t, "2nodes.top", "2nodes-message.events", []string{"2nodes-message.snap"})
}

func Test3NodesMultipleMessages(t *testing.T) {
	runTest(t, "3nodes.top", "3nodes-simple.events", []string{"3nodes-simple.snap"})
}

func Test3NodesMultipleBidirectionalMessages(t *testing.T) {
	runTest(
		t,
		"3nodes.top",
		"3nodes-bidirectional-messages.events",
		[]string{"3nodes-bidirectional-messages.snap"})
}

func Test8NodesSequentialSnapshots(t *testing.T) {
	runTest(
		t,
		"8nodes.top",
		"8nodes-sequential-snapshots.events",
		[]string{
			"8nodes-sequential-snapshots0.snap",
			"8nodes-sequential-snapshots1.snap",
		})
}

func Test8NodesConcurrentSnapshots(t *testing.T) {
	runTest(
		t,
		"8nodes.top",
		"8nodes-concurrent-snapshots.events",
		[]string{
			"8nodes-concurrent-snapshots0.snap",
			"8nodes-concurrent-snapshots1.snap",
			"8nodes-concurrent-snapshots2.snap",
			"8nodes-concurrent-snapshots3.snap",
			"8nodes-concurrent-snapshots4.snap",
		})
}

func Test10NodesDirectedEdges(t *testing.T) {
	runTest(
		t,
		"10nodes.top",
		"10nodes.events",
		[]string{
			"10nodes0.snap",
			"10nodes1.snap",
			"10nodes2.snap",
			"10nodes3.snap",
			"10nodes4.snap",
			"10nodes5.snap",
			"10nodes6.snap",
			"10nodes7.snap",
			"10nodes8.snap",
			"10nodes9.snap",
		})
}
