package client_test

// You MUST NOT change these default imports.  ANY additional imports may
// break the autograder and everyone will be sad.

import (
	// Some imports use an underscore to prevent the compiler from complaining
	// about unused imports.
	_ "encoding/hex"
	_ "errors"
	_ "strconv"
	_ "strings"
	"testing"

	// A "dot" import is used here so that the functions in the ginko and gomega
	// modules can be used without an identifier. For example, Describe() and
	// Expect() instead of ginko.Describe() and gomega.Expect().
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	userlib "github.com/cs161-staff/project2-userlib"

	"github.com/cs161-staff/project2-starter-code/client"
)

func TestSetupAndExecution(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Tests")
}

// ================================================
// Global Variables (feel free to add more!)
// ================================================
const defaultPassword = "password"
const emptyString = ""
const contentOne = "Bitcoin is Nick's favorite "
const contentTwo = "digital "
const contentThree = "cryptocurrency!"

// ================================================
// Describe(...) blocks help you organize your tests
// into functional categories. They can be nested into
// a tree-like structure.
// ================================================

var _ = Describe("Client Tests", func() {

	// A few user declarations that may be used for testing. Remember to initialize these before you
	// attempt to use them!
	var alice *client.User
	var bob *client.User
	var charles *client.User
	// var doris *client.User
	// var eve *client.User
	var frank *client.User
	var grace *client.User
	var horace *client.User
	// var ira *client.User

	// These declarations may be useful for multi-session testing.
	var alicePhone *client.User
	var aliceLaptop *client.User
	var aliceDesktop *client.User
	var bobLaptop *client.User
	var charlesLaptop *client.User
	var graceLaptop *client.User
	var frankLaptop *client.User

	var err error

	// A bunch of filenames that may be useful.
	aliceFile := "aliceFile.txt"
	aliceFile2 := "aliceFile2.txt"
	bobFile := "bobFile.txt"
	charlesFile := "charlesFile.txt"
	charlesFile2 := "charlesFile2.txt"
	// dorisFile := "dorisFile.txt"
	// eveFile := "eveFile.txt"
	frankFile := "frankFile.txt"
	frankFile2 := "frankFile2.txt"
	graceFile := "graceFile.txt"
	graceFile2 := "graceFile2.txt"
	horaceFile := "horaceFile.txt"
	horaceFile2 := "horaceFile2.txt"

	// iraFile := "iraFile.txt"

	BeforeEach(func() {
		// This runs before each test within this Describe block (including nested tests).
		// Here, we reset the state of Datastore and Keystore so that tests do not interfere with each other.
		// We also initialize
		userlib.DatastoreClear()
		userlib.KeystoreClear()
	})

	Describe("Basic Tests", func() {

		Specify("Basic Test: Testing InitUser/GetUser on a single user.", func() {
			userlib.DebugMsg("Initializing user Alice.")
			alice, err = client.InitUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Getting user Alice.")
			aliceLaptop, err = client.GetUser("alice", defaultPassword)
			Expect(err).To(BeNil())
		})

		Specify("Basic Test: Testing Single User Store/Load/Append.", func() {
			userlib.DebugMsg("Initializing user Alice.")
			alice, err = client.InitUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Storing file data: %s", contentOne)
			err = alice.StoreFile(aliceFile, []byte(contentOne))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Appending file data: %s", contentTwo)
			err = alice.AppendToFile(aliceFile, []byte(contentTwo))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Appending file data: %s", contentThree)
			err = alice.AppendToFile(aliceFile, []byte(contentThree))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Loading file...")
			data, err := alice.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne + contentTwo + contentThree)))
		})

		Specify("Basic Test: Testing Create/Accept Invite Functionality with multiple users and multiple instances.", func() {
			userlib.DebugMsg("Initializing users Alice (aliceDesktop) and Bob.")
			aliceDesktop, err = client.InitUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			bob, err = client.InitUser("bob", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Getting second instance of Alice - aliceLaptop")
			aliceLaptop, err = client.GetUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("aliceDesktop storing file %s with content: %s", aliceFile, contentOne)
			err = aliceDesktop.StoreFile(aliceFile, []byte(contentOne))
			Expect(err).To(BeNil())

			userlib.DebugMsg("aliceLaptop creating invite for Bob.")
			invite, err := aliceLaptop.CreateInvitation(aliceFile, "bob")
			Expect(err).To(BeNil())

			userlib.DebugMsg("Bob accepting invite from Alice under filename %s.", bobFile)
			err = bob.AcceptInvitation("alice", invite, bobFile)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Bob appending to file %s, content: %s", bobFile, contentTwo)
			err = bob.AppendToFile(bobFile, []byte(contentTwo))
			Expect(err).To(BeNil())

			userlib.DebugMsg("aliceDesktop appending to file %s, content: %s", aliceFile, contentThree)
			err = aliceDesktop.AppendToFile(aliceFile, []byte(contentThree))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Checking that aliceDesktop sees expected file data.")
			data, err := aliceDesktop.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne + contentTwo + contentThree)))

			userlib.DebugMsg("Checking that aliceLaptop sees expected file data.")
			data, err = aliceLaptop.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne + contentTwo + contentThree)))

			userlib.DebugMsg("Checking that Bob sees expected file data.")
			data, err = bob.LoadFile(bobFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne + contentTwo + contentThree)))

			userlib.DebugMsg("Getting third instance of Alice - alicePhone.")
			alicePhone, err = client.GetUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Checking that alicePhone sees Alice's changes.")
			data, err = alicePhone.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne + contentTwo + contentThree)))
		})

		Specify("Basic Test: Testing Revoke Functionality", func() {
			userlib.DebugMsg("Initializing users Alice, Bob, and Charlie.")
			alice, err = client.InitUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			bob, err = client.InitUser("bob", defaultPassword)
			Expect(err).To(BeNil())

			charles, err = client.InitUser("charles", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice storing file %s with content: %s", aliceFile, contentOne)
			alice.StoreFile(aliceFile, []byte(contentOne))

			userlib.DebugMsg("Alice creating invite for Bob for file %s, and Bob accepting invite under name %s.", aliceFile, bobFile)

			invite, err := alice.CreateInvitation(aliceFile, "bob")
			Expect(err).To(BeNil())

			err = bob.AcceptInvitation("alice", invite, bobFile)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Checking that Alice can still load the file.")
			data, err := alice.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne)))

			userlib.DebugMsg("Checking that Bob can load the file.")
			data, err = bob.LoadFile(bobFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne)))

			userlib.DebugMsg("Bob creating invite for Charles for file %s, and Charlie accepting invite under name %s.", bobFile, charlesFile)
			invite, err = bob.CreateInvitation(bobFile, "charles")
			Expect(err).To(BeNil())

			err = charles.AcceptInvitation("bob", invite, charlesFile)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Checking that Bob can load the file.")
			data, err = bob.LoadFile(bobFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne)))

			userlib.DebugMsg("Checking that Charles can load the file.")
			data, err = charles.LoadFile(charlesFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne)))

			userlib.DebugMsg("Alice revoking Bob's access from %s.", aliceFile)
			err = alice.RevokeAccess(aliceFile, "bob")
			Expect(err).To(BeNil())

			userlib.DebugMsg("Checking that Alice can still load the file.")
			data, err = alice.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne)))

			userlib.DebugMsg("Checking that Bob/Charles lost access to the file.")
			_, err = bob.LoadFile(bobFile)
			Expect(err).ToNot(BeNil())

			_, err = charles.LoadFile(charlesFile)
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("Checking that the revoked users cannot append to the file.")
			err = bob.AppendToFile(bobFile, []byte(contentTwo))
			Expect(err).ToNot(BeNil())

			err = charles.AppendToFile(charlesFile, []byte(contentTwo))
			Expect(err).ToNot(BeNil())
		})
	})

	Describe("Custom Tests", func() {
		Specify("GetUser on a single user and none exist user.", func() {
			userlib.DebugMsg("Initializing user Alice.")
			alice, err := client.InitUser("alice", "password")
			Expect(err).To(BeNil())
			Expect(alice).ToNot(BeNil())

			userlib.DebugMsg("Getting user Alice.")
			aliceLaptop, err := client.GetUser("alice", "password")
			Expect(err).To(BeNil())
			Expect(aliceLaptop).ToNot(BeNil())

			userlib.DebugMsg("Getting none exist user.")
			bob, err := client.GetUser("bob", "password123")
			Expect(err).ToNot(BeNil())
			Expect(bob).To(BeNil())

			userlib.DebugMsg("Getting user Alice but wrong password.")
			aliceMobile, err := client.GetUser("alice", "password123")
			Expect(err).ToNot(BeNil())
			Expect(aliceMobile).To(BeNil())

			//test empty username
			userlib.DebugMsg("Initializing user with empty unername.")
			empty, err := client.InitUser("", "bananaLord")
			Expect(err).ToNot(BeNil())
			Expect(empty).To(BeNil())

		})

		Specify("Testing InitUser/GetUser on a single user, existed user, and none exist user.", func() {
			userlib.DebugMsg("Initializing user Alice.")
			alice, err = client.InitUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Getting user Alice.")
			alice, err = client.GetUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Getting none exist user.")
			bob, err := client.GetUser("bob", "password123")
			Expect(err).ToNot(BeNil())
			Expect(bob).To(BeNil())

			userlib.DebugMsg("Getting user Alice but wrong password.")
			alice, err = client.GetUser("alice", "password123")
			Expect(err).ToNot(BeNil())
			Expect(alice).To(BeNil())

			userlib.DebugMsg("Initializing user with empty username.")
			empty, err := client.InitUser("", "bananaLord")
			Expect(err).ToNot(BeNil())
			Expect(empty).To(BeNil())

			userlib.DebugMsg("Initializing user with existed username Alice.")
			alice2, err2 := client.InitUser("alice", defaultPassword)
			Expect(alice2).To(BeNil())
			Expect(err2).ToNot(BeNil())

		})

		Specify("Testing Multiple Users Store/Load/Append", func() {
			userlib.DebugMsg("Initializing user Alice.")
			alice, err = client.InitUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Getting user Alice on Laptop.")
			aliceLaptop, err = client.GetUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Getting user Alice on Phone.")
			alicePhone, err = client.GetUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice Storing 0 string filename with empty data: %s", "")
			err = alice.StoreFile("", []byte(""))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice Load 0 string file name data: %s", "")
			data, err := alice.LoadFile("")
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("")))

			userlib.DebugMsg("Alice on Laptop Storing file data: %s", contentOne)
			err = aliceLaptop.StoreFile(aliceFile, []byte(contentOne))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice on Laptop Loading file...")
			data, err = aliceLaptop.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne)))

			userlib.DebugMsg("Alice on Laptop Overwrite the file data: %s", contentTwo)
			err = aliceLaptop.StoreFile(aliceFile, []byte(contentTwo))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice on Phone Loading file that has been overwritten...")
			data, err = alicePhone.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentTwo)))

			userlib.DebugMsg("Alice on Phone Loading file that doesnot existed (aliceFile2)")
			data, err = alicePhone.LoadFile(aliceFile2)
			Expect(err).ToNot(BeNil())
			Expect(data).To(BeNil())

			userlib.DebugMsg("Alice on Phone Storing new file data (aliceFile2): %s", contentOne)
			err = alicePhone.StoreFile(aliceFile2, []byte(contentOne))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice on Laptop can Loading that new file (aliceFile2)")
			data, err = aliceLaptop.LoadFile(aliceFile2)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne)))

			userlib.DebugMsg("Alice Appending file data: %s", contentTwo)
			err = alice.AppendToFile(aliceFile2, []byte(contentTwo))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice Loading file that appended...")
			data, err = alice.LoadFile(aliceFile2)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne + contentTwo)))

			userlib.DebugMsg("Alice Appending file data: %s", contentThree)
			err = alice.AppendToFile(aliceFile2, []byte(contentThree))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice on Phone Loading file that appended...")
			data, err = alicePhone.LoadFile(aliceFile2)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne + contentTwo + contentThree)))

			userlib.DebugMsg("Alice Appending file data: %s", contentOne)
			err = alice.AppendToFile(aliceFile2, []byte(contentOne))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice on Phone Loading file that appended...")
			data, err = alicePhone.LoadFile(aliceFile2)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(contentOne + contentTwo + contentThree + contentOne)))

		})
		Specify("Custom Testing Share/Accept Invitation", func() {
			//SUMMARY OF TEST:
			//Alice is the creator of file. She shares to Charles and appends various messages to file. Both Charles and Alices see these changes.
			//Charles fileStores new contents and file and appends new stuff. Both charles and Alice see these changes.
			//Then charles invites Frank to file. Frank fileStores new contents to files and everyone sees changes (including charles and alices on laptop).

			//Note all user that are shared with file should be able to: (i believe i tested all of these - mark / but double check)
			//Read the file contents with LoadFile.
			//Overwrite the file contents with StoreFile.
			//Append to the file with AppendToFile.
			//Share the file with CreateInvitation.

			//This custom checks for errors such as return error when:
			// 1)The given filename does not exist in the personal file namespace of the caller.
			// 2)The given recipientUsername does not exist.
			// 3)The user already has a file with the chosen filename in their personal file namespace.

			//for accept invition, not sure how to test:
			//Something about the invitationPtr is wrong
			//(e.g. the value at that UUID on Datastore is corrupt or missing, or the user cannot verify that invitationPtr was provided by senderUsername).
			userlib.DebugMsg("Initializing user Alice.")
			alice, err = client.InitUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Initializing user Charles.")
			charles, err = client.InitUser("charles", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice stores 'Hello world' in file")
			err = alice.StoreFile(aliceFile, []byte("Hello World!"))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice created invitation for user Zoe. (this should return error since Zoe DNE)")
			_, err := alice.CreateInvitation(aliceFile, "zoe") //error (2)
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("Alice created invitation but wrong fileName (return error since filename DNE)")
			_, err = alice.CreateInvitation(aliceFile+"DNE", "charles") //error (1)
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("Alice created invitation for user Charles.")
			invite, err := alice.CreateInvitation(aliceFile, "charles")
			Expect(err).To(BeNil())

			userlib.DebugMsg("Charles accepted invitation")
			err = charles.AcceptInvitation("alice", invite, charlesFile)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Both charles and alice should see 'hello world!' in loadFile")
			data, err := alice.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Hello World!")))

			data, err = charles.LoadFile(charlesFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Hello World!")))

			userlib.DebugMsg("Alice appends 'This is Alice's File.")
			err = alice.AppendToFile(aliceFile, []byte(" This is Alice's File."))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice appends I'm sharing it with Charles.'")
			err = alice.AppendToFile(aliceFile, []byte(" I'm sharing it with Charles."))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Both charles and alice should see 'Hello world! This is Alice's File. I'm sharing it with Charles' in loadFile")
			data, err = alice.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Hello World! This is Alice's File. I'm sharing it with Charles.")))

			data, err = charles.LoadFile(charlesFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Hello World! This is Alice's File. I'm sharing it with Charles.")))

			userlib.DebugMsg("Charles appends ' Hi I'm Charles'")
			err = charles.AppendToFile(charlesFile, []byte(" Hi I'm Charles."))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Both charles and alice should see 'Hello world! This is Alice's File. I'm sharing it with Charles. Hi I'm Charles.' in loadFile")
			data, err = alice.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Hello World! This is Alice's File. I'm sharing it with Charles. Hi I'm Charles.")))

			data, err = charles.LoadFile(charlesFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Hello World! This is Alice's File. I'm sharing it with Charles. Hi I'm Charles.")))

			userlib.DebugMsg("Charles stores 'File reset.'")
			err = charles.StoreFile(charlesFile, []byte("File reset."))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice appends ' First edit.")
			err = alice.AppendToFile(aliceFile, []byte(" First edit."))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Charles appends ' Second edit.")
			err = charles.AppendToFile(charlesFile, []byte(" Second edit."))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Both charles and alice should see 'File reset. First edit. Second edit.'")
			data, err = alice.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("File reset. First edit. Second edit.")))

			data, err = charles.LoadFile(charlesFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("File reset. First edit. Second edit.")))

			userlib.DebugMsg("Getting user Alice on Laptop.")
			aliceLaptop, err = client.GetUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Getting user Charles on Laptop.")
			charlesLaptop, err = client.GetUser("charles", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Both charles and alice should see on Laptop 'File reset. First edit. Second edit.'")
			data, err = aliceLaptop.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("File reset. First edit. Second edit.")))

			data, err = charlesLaptop.LoadFile(charlesFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("File reset. First edit. Second edit.")))

			userlib.DebugMsg("Initializing user frank.")
			frank, err = client.InitUser("frank", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Charles Laptop created invitation for user Frank.")
			invite, err = charlesLaptop.CreateInvitation(charlesFile, "frank")
			Expect(err).To(BeNil())

			userlib.DebugMsg("Frank accepts invitation from zoe (return error since user zoe DNE)")
			err = frank.AcceptInvitation("zoe", invite, frankFile)
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("Frank stores 'Hello world' in frankFile")
			err = frank.StoreFile(frankFile, []byte("Hello World!"))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Frank accepts invitation from charles (but erroers since user already has a file with the chosen filename in their personal file namespace.")
			err = frank.AcceptInvitation("charles", invite, frankFile) //error 3
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("Frank accepts invitation from charles under frankFile2 instead")
			err = frank.AcceptInvitation("charles", invite, frankFile2)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Frank should see 'File reset. First edit. Second edit.'")
			data, err = frank.LoadFile(frankFile2)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("File reset. First edit. Second edit.")))

			userlib.DebugMsg("Frank stores empty string ''")
			err = frank.StoreFile(frankFile2, []byte(""))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice on laptop appends ' First edit again")
			err = alice.AppendToFile(aliceFile, []byte(" First edit again."))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Frank appends ' second edit again")
			err = alice.AppendToFile(aliceFile, []byte(" second edit again."))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice, Charles, Frank should all see: ' First edit again second edit again'")
			data, err = alice.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(" First edit again. second edit again.")))

			data, err = charles.LoadFile(charlesFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(" First edit again. second edit again.")))

			data, err = aliceLaptop.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(" First edit again. second edit again.")))

			data, err = charlesLaptop.LoadFile(charlesFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(" First edit again. second edit again.")))

			data, err = frank.LoadFile(frankFile2)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(" First edit again. second edit again.")))

			userlib.DebugMsg("Initializing user Grace.")
			grace, err = client.InitUser("grace", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice created invitation for user Grace.")
			invite, err = alice.CreateInvitation(aliceFile, "grace")
			Expect(err).To(BeNil())

			userlib.DebugMsg("Grace trys to load file, but should error since she hasn't accept invitation.")
			data, err = grace.LoadFile(graceFile)
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("Grace now accepts file")
			err = grace.AcceptInvitation("alice", invite, graceFile)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Grace can now succesfully load file.")
			data, err = grace.LoadFile(graceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte(" First edit again. second edit again.")))
		})
		Specify("Custom Testing Revoke Access", func() {
			userlib.DebugMsg("Initializing users Alice, Bob, Charlie, Frank, Grace and Horace.")
			alice, err = client.InitUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			bob, err = client.InitUser("bob", defaultPassword)
			Expect(err).To(BeNil())

			charles, err = client.InitUser("charles", defaultPassword)
			Expect(err).To(BeNil())

			frank, err = client.InitUser("frank", defaultPassword)
			Expect(err).To(BeNil())

			grace, err = client.InitUser("grace", defaultPassword)
			Expect(err).To(BeNil())

			horace, err = client.InitUser("horace", defaultPassword)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice stores 'Alices File' in file and direct shares to bob and charles")

			err = alice.StoreFile(aliceFile, []byte("Alices File"))
			Expect(err).To(BeNil())

			inviteBob, err := alice.CreateInvitation(aliceFile, "bob")
			Expect(err).To(BeNil())

			inviteCharles, err := alice.CreateInvitation(aliceFile, "charles")
			Expect(err).To(BeNil())

			err = bob.AcceptInvitation("alice", inviteBob, bobFile)
			Expect(err).To(BeNil())

			err = charles.AcceptInvitation("alice", inviteCharles, charlesFile)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Bob shares Alice File with frank and Grace")

			invite, err := bob.CreateInvitation(bobFile, "frank")
			Expect(err).To(BeNil())

			err = frank.AcceptInvitation("bob", invite, frankFile)
			Expect(err).To(BeNil())

			invite, err = bob.CreateInvitation(bobFile, "grace")
			Expect(err).To(BeNil())

			err = grace.AcceptInvitation("bob", invite, graceFile)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Charles shares Alice file with horace")

			invite, err = charles.CreateInvitation(charlesFile, "horace")
			Expect(err).To(BeNil())

			err = horace.AcceptInvitation("charles", invite, horaceFile)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Everyone will append one thing to file")

			err = alice.AppendToFile(aliceFile, []byte(". Alice Edit. "))
			Expect(err).To(BeNil())

			err = horace.AppendToFile(horaceFile, []byte("Horace Edit. "))
			Expect(err).To(BeNil())

			err = grace.AppendToFile(graceFile, []byte("Grace Edit. "))
			Expect(err).To(BeNil())

			err = frank.AppendToFile(frankFile, []byte("Frank Edit. "))
			Expect(err).To(BeNil())

			err = bob.AppendToFile(bobFile, []byte("Bob Edit. "))
			Expect(err).To(BeNil())

			err = charles.AppendToFile(charlesFile, []byte("Charles Edit. "))
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice revoking charles's access from Alice File")
			err = alice.RevokeAccess(aliceFile, "charles")
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice, Bob, frank, and Grace should see Alice File. Charles and Horace should not ")

			data, err := alice.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Alices File. Alice Edit. Horace Edit. Grace Edit. Frank Edit. Bob Edit. Charles Edit. ")))

			data, err = frank.LoadFile(frankFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Alices File. Alice Edit. Horace Edit. Grace Edit. Frank Edit. Bob Edit. Charles Edit. ")))

			data, err = grace.LoadFile(graceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Alices File. Alice Edit. Horace Edit. Grace Edit. Frank Edit. Bob Edit. Charles Edit. ")))

			data, err = bob.LoadFile(bobFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Alices File. Alice Edit. Horace Edit. Grace Edit. Frank Edit. Bob Edit. Charles Edit. ")))

			data, err = charles.LoadFile(charlesFile)
			Expect(err).ToNot(BeNil())

			data, err = horace.LoadFile(horaceFile)
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("Charles and Horace cannot append, or share file now since they are revoked")

			err = charles.AppendToFile(charlesFile, []byte("charles trying to append after revoked. "))
			Expect(err).ToNot(BeNil())

			err = horace.AppendToFile(horaceFile, []byte("horace trying to append after revoked. "))
			Expect(err).ToNot(BeNil())

			err = charles.StoreFile(charlesFile, []byte("trying to store after revoked"))
			Expect(err).ToNot(BeNil())

			err = horace.StoreFile(horaceFile, []byte("trying to store after revoked"))
			Expect(err).ToNot(BeNil())

			invite, err = charles.CreateInvitation(charlesFile, "horace")
			Expect(err).ToNot(BeNil())

			invite, err = horace.CreateInvitation(horaceFile, "charles")
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("Bob will now make a new file named AliceFile and invite charles who will invite Horace")

			err = bob.StoreFile(aliceFile, []byte("Bob's File named AliceFile.txt"))
			Expect(err).To(BeNil())

			invite, err = bob.CreateInvitation(aliceFile, "charles")
			Expect(err).To(BeNil())

			err = charles.StoreFile("existFile", []byte("create existedFile to test accept on existedFile"))
			Expect(err).To(BeNil())

			err = charles.AcceptInvitation("bob", invite, "existFile")
			Expect(err).ToNot(BeNil())

			err = charles.AcceptInvitation("bob", invite, charlesFile2)
			Expect(err).To(BeNil())

			data, err = charles.LoadFile(charlesFile2)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Bob's File named AliceFile.txt")))

			invite, err = charles.CreateInvitation(charlesFile2, "horace")
			Expect(err).To(BeNil())

			err = horace.AcceptInvitation("charles", invite, horaceFile2)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Bob logins laptop and will now share to frank and frank will share to grace. Before grace accpets invite, frank will be revoked. Therefore grace should not be able to accept invite")

			bobLaptop, err = client.GetUser("bob", defaultPassword)
			Expect(err).To(BeNil())

			invite, err = bobLaptop.CreateInvitation(aliceFile, "frank")
			Expect(err).To(BeNil())

			err = frank.AcceptInvitation("bob", invite, frankFile2)
			Expect(err).To(BeNil())

			invite, err = frank.CreateInvitation(frankFile2, "grace")
			Expect(err).To(BeNil())

			err = bobLaptop.RevokeAccess(aliceFile, "frank")
			Expect(err).To(BeNil())

			err = grace.AcceptInvitation("frank", invite, graceFile2)
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("Frank and Grace should not be able to access file, but Bob, charles and Horce should still be able to")

			data, err = grace.LoadFile(graceFile2)
			Expect(err).ToNot(BeNil())

			data, err = frank.LoadFile(frankFile2)
			Expect(err).ToNot(BeNil())

			data, err = bob.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Bob's File named AliceFile.txt")))

			data, err = bobLaptop.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Bob's File named AliceFile.txt")))

			data, err = horace.LoadFile(horaceFile2)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Bob's File named AliceFile.txt")))

			data, err = charles.LoadFile(charlesFile2)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Bob's File named AliceFile.txt")))

			userlib.DebugMsg("Frank and grace cannot log into new device and access file")

			frankLaptop, err = client.GetUser("frank", defaultPassword)
			Expect(err).To(BeNil())

			graceLaptop, err = client.GetUser("grace", defaultPassword)
			Expect(err).To(BeNil())

			data, err = grace.LoadFile(graceFile2)
			Expect(err).ToNot(BeNil())

			data, err = frank.LoadFile(frankFile2)
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("But they still should be able to access Alice's shared file")

			data, err = frankLaptop.LoadFile(frankFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Alices File. Alice Edit. Horace Edit. Grace Edit. Frank Edit. Bob Edit. Charles Edit. ")))

			data, err = graceLaptop.LoadFile(graceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Alices File. Alice Edit. Horace Edit. Grace Edit. Frank Edit. Bob Edit. Charles Edit. ")))

			userlib.DebugMsg("Bob invite alice to file twice (This is allowed).")

			invite, err = bob.CreateInvitation(aliceFile, "alice")
			Expect(err).To(BeNil())

			invite, err = bobLaptop.CreateInvitation(aliceFile, "alice")
			Expect(err).To(BeNil())

			err = alice.AcceptInvitation("bob", invite, aliceFile2)
			Expect(err).To(BeNil())

			userlib.DebugMsg("Alice should be able to access bob file and her own file on different devices as well")

			aliceLaptop, err = client.GetUser("alice", defaultPassword)
			Expect(err).To(BeNil())

			data, err = aliceLaptop.LoadFile(aliceFile2)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Bob's File named AliceFile.txt")))

			data, err = aliceLaptop.LoadFile(aliceFile)
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]byte("Alices File. Alice Edit. Horace Edit. Grace Edit. Frank Edit. Bob Edit. Charles Edit. ")))

			err = bob.RevokeAccess(aliceFile, "NullUser")
			Expect(err).ToNot(BeNil())

			userlib.DebugMsg("Alice laptop revokes bob")

			err = aliceLaptop.RevokeAccess(aliceFile, "bob")
			Expect(err).To(BeNil())

			userlib.DebugMsg("Now bob, frank, grace cannot access file. Even on their laptops ")

			data, err = bob.LoadFile(bobFile)
			Expect(err).ToNot(BeNil())

			data, err = bobLaptop.LoadFile(bobFile)
			Expect(err).ToNot(BeNil())

			data, err = frank.LoadFile(frankFile)
			Expect(err).ToNot(BeNil())

			data, err = frankLaptop.LoadFile(frankFile)
			Expect(err).ToNot(BeNil())

			data, err = grace.LoadFile(graceFile)
			Expect(err).ToNot(BeNil())

			data, err = graceLaptop.LoadFile(graceFile)
			Expect(err).ToNot(BeNil())
		})
	})

})
