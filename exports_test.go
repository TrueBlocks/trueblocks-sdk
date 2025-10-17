package sdk

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/output"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// TestSDKChannelClosureBug demonstrates that SDK export functions
// complete successfully but never close their output channels
func TestSDKChannelClosureBug(t *testing.T) {
	// Create render context - this is what the SDK uses for output
	renderCtx := &output.RenderCtx{
		ModelChan: make(chan types.Modeler, 100),
		ErrorChan: make(chan error, 10),
		Ctx:       context.Background(),
	}

	// Track what happens
	start := time.Now()
	itemCount := 0
	channelsClosed := false

	// Start goroutine to monitor channels
	done := make(chan struct{})
	go func() {
		defer close(done)

		// Call the SDK function that should close channels when done
		exportOpts := ExportOptions{
			Globals: Globals{
				Cache:   false,
				Verbose: false,
				Chain:   "mainnet",
			},
			RenderCtx:  renderCtx,
			Addrs:      []string{"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"}, // vitalik.eth
			Articulate: false,
			FirstBlock: 18000000, // Recent block range
			LastBlock:  18000010, // Small range for test
		}

		// This call should succeed and close channels when done
		if _, _, err := exportOpts.ExportApprovals(); err != nil {
			renderCtx.ErrorChan <- fmt.Errorf("ExportApprovals failed: %w", err)
		}

		// At this point, SDK should close channels but doesn't!
	}()

	// Wait for data and monitor for channel closure
	timeout := time.NewTimer(5 * time.Second) // Give it 5 seconds max
	defer timeout.Stop()

	for {
		select {
		case item, ok := <-renderCtx.ModelChan:
			if !ok {
				// ModelChan closed - good!
				fmt.Printf("✅ ModelChan closed after receiving %d items\n", itemCount)
				channelsClosed = true
				// Check if ErrorChan also closes
				errorCloseTimeout := time.NewTimer(100 * time.Millisecond)
				defer errorCloseTimeout.Stop()

				select {
				case _, ok := <-renderCtx.ErrorChan:
					if !ok {
						fmt.Printf("✅ ErrorChan also closed properly\n")
						duration := time.Since(start)
						fmt.Printf("✅ SUCCESS: Test completed in %v with properly closed channels\n", duration)
						fmt.Printf("📈 Total items received: %d\n", itemCount)
						return
					}
				case <-errorCloseTimeout.C:
					fmt.Printf("🐛 BUG: ErrorChan also never closed\n")
					duration := time.Since(start)
					fmt.Printf("❌ FAILURE: Channel closure bug confirmed after %v\n", duration)
					fmt.Printf("📈 Items received before timeout: %d\n", itemCount)
					t.Fatalf("SDK bug: ExportApprovals() completed but did not close output channels")
				}
				return
			}
			itemCount++
			fmt.Printf("📦 Received item %d: %T\n", itemCount, item)

		case err, ok := <-renderCtx.ErrorChan:
			if !ok {
				// ErrorChan closed
				fmt.Printf("✅ ErrorChan closed\n")
				if channelsClosed {
					duration := time.Since(start)
					fmt.Printf("✅ SUCCESS: Test completed in %v with properly closed channels\n", duration)
					fmt.Printf("📈 Total items received: %d\n", itemCount)
					return
				}
			}
			fmt.Printf("❌ Error: %v\n", err)

		case <-done:
			fmt.Printf("🏁 SDK function completed in %v\n", time.Since(start))
			fmt.Printf("📊 Received %d items\n", itemCount)

			// Now check if channels close naturally
			closeCheckStart := time.Now()
			closeTimeout := time.NewTimer(100 * time.Millisecond)
			defer closeTimeout.Stop()

			fmt.Printf("⏳ Waiting for channels to close...\n")

		waitLoop:
			for {
				select {
				case _, ok := <-renderCtx.ModelChan:
					if !ok {
						fmt.Printf("✅ ModelChan closed %v after SDK completion\n", time.Since(closeCheckStart))
						channelsClosed = true
						break waitLoop
					}
					itemCount++

				case <-closeTimeout.C:
					fmt.Printf("🐛 BUG CONFIRMED: ModelChan never closed after %v\n", time.Since(closeCheckStart))
					fmt.Printf("🐛 SDK function completed but left channels open!\n")
					duration := time.Since(start)
					fmt.Printf("❌ FAILURE: Channel closure bug confirmed after %v\n", duration)
					fmt.Printf("📈 Items received before timeout: %d\n", itemCount)
					t.Fatalf("SDK bug: ExportApprovals() completed but did not close output channels")
				}
			}

			// Check if ErrorChan also closes
			errorCloseTimeout := time.NewTimer(100 * time.Millisecond)
			defer errorCloseTimeout.Stop()

			select {
			case _, ok := <-renderCtx.ErrorChan:
				if !ok {
					fmt.Printf("✅ ErrorChan also closed properly\n")
					duration := time.Since(start)
					fmt.Printf("✅ SUCCESS: Test completed in %v with properly closed channels\n", duration)
					fmt.Printf("📈 Total items received: %d\n", itemCount)
					return
				}
			case <-errorCloseTimeout.C:
				fmt.Printf("🐛 BUG: ErrorChan also never closed\n")
				duration := time.Since(start)
				fmt.Printf("❌ FAILURE: Channel closure bug confirmed after %v\n", duration)
				fmt.Printf("📈 Items received before timeout: %d\n", itemCount)
				t.Fatalf("SDK bug: ExportApprovals() completed but did not close output channels")
			}

		case <-timeout.C:
			fmt.Printf("⏰ TIMEOUT: SDK function never completed after 5 seconds\n")
			duration := time.Since(start)
			fmt.Printf("❌ FAILURE: Channel closure bug confirmed after %v\n", duration)
			fmt.Printf("📈 Items received before timeout: %d\n", itemCount)
			t.Fatalf("SDK bug: ExportApprovals() completed but did not close output channels")
		}
	}
}

// If you want to run this as a standalone program instead of test:
// func main() {
// 	fmt.Println("🧪 Testing SDK channel closure behavior...")
// 	t := &testing.T{} // Mock testing.T for standalone run
// 	TestSDKChannelClosureBug(t)
// }
