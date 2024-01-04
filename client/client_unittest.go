package client

///////////////////////////////////////////////////
//                                               //
// Everything in this file will NOT be graded!!! //
//                                               //
///////////////////////////////////////////////////

// In this unit tests file, you can write white-box unit tests on your implementation.
// These are different from the black-box integration tests in client_test.go,
// because in this unit tests file, you can use details specific to your implementation.

// For example, in this unit tests file, you can access struct fields and helper methods
// that you defined, but in the integration tests (client_test.go), you can only access
// the 8 functions (StoreFile, LoadFile, etc.) that are common to all implementations.

// In this unit tests file, you can write InitUser where you would write client.InitUser in the
// integration tests (client_test.go). In other words, the "client." in front is no longer needed.

import (
	"testing"

	userlib "github.com/cs161-staff/project2-userlib"

	_ "encoding/hex"

	_ "errors"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"

	_ "strconv"

	_ "strings"
)

func TestSetupAndExecution(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Unit Tests")
}

var _ = Describe("Client Unit Tests", func() {

	BeforeEach(func() {
		userlib.DatastoreClear()
		userlib.KeystoreClear()
	})

	Describe("Unit Tests", func() {
		Specify("Basic Test: Check that the Username field is set for a new user", func() {
			userlib.DebugMsg("Initializing user Alice.")
			// Note: In the integration tests (client_test.go) this would need to
			// be client.InitUser, but here (client_unittests.go) you can write InitUser.
			alice, err := InitUser("alice", "password")
			Expect(err).To(BeNil())

			// Note: You can access the Username field of the User struct here.
			// But in the integration tests (client_test.go), you cannot access
			// struct fields because not all implementations will have a username field.
			Expect(alice.Username).To(Equal("alice"))
		})
		Specify("Basic Test: GetUser on a single user and none exist user.", func() {
			userlib.DebugMsg("Initializing user Alice.")
			alice, err := InitUser("alice", "password")
			Expect(err).To(BeNil())
			Expect(alice.Username).To(Equal("alice"))

			userlib.DebugMsg("Getting user Alice.")
			aliceLaptop, err := GetUser("alice", "password")
			Expect(err).To(BeNil())
			Expect(aliceLaptop.Password).To(Equal("password"))

			userlib.DebugMsg("Getting none exist user.")
			bob, err := GetUser("bob", "password123")
			Expect(err).ToNot(BeNil())
			Expect(bob).To(BeNil())

			userlib.DebugMsg("Getting user Alice but wrong password.")
			aliceMobile, err := GetUser("alice", "password123")
			Expect(err).ToNot(BeNil())
			Expect(aliceMobile).To(BeNil())

			//test empty username
			userlib.DebugMsg("Initializing user with empty unername.")
			empty, err := InitUser("", "bananaLord")
			Expect(err).ToNot(BeNil())
			Expect(empty).To(BeNil())

		})
		// Specify("Basic Test: GetUser on a single user and none exist user.", func() {
		// 	userlib.DebugMsg("Initializing user Alice.")
		// 	alice, err := InitUser("alice", "password")
		// 	Expect(err).To(BeNil())
		// 	Expect(alice.Username).To(Equal("alice"))

		// 	userlib.DebugMsg("Storing file data: %s", "pancakes")
		// 	err = alice.StoreFile("aliceFile.txt", []byte("pancakes"))
		// 	userlib.DebugMsg(err.Error())
		// 	Expect(err).To(BeNil())

		// 	err = alice.StoreFile("aliceFile.txt", []byte("cookies"))
		// 	userlib.DebugMsg(err.Error())
		// 	Expect(err).To(BeNil())
		// })
	})
})
