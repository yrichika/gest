package gt

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

type Test struct {
	testingT *testing.T
	testName string
	messages []string
	passed   int
	subtotal int
	isGestOn bool
	// check if a running test is failed
	isThisTestFailed bool
	// check if any test in the suite is failed
	isAnyTestFailed bool
	isAsyncEnabled  bool
	isSkipping      bool
	beforeEach      func()
	afterEach       func()
	beforeAll       func()
	afterAll        func()
}

func CreateTest(t *testing.T) *Test {
	isGestOn := len(os.Getenv(EnvName)) > 0
	testFileName := getTestName(t)
	return &Test{
		testingT:         t,
		testName:         testFileName,
		passed:           0,
		subtotal:         0,
		isGestOn:         isGestOn,
		isThisTestFailed: false,
		isAnyTestFailed:  false,
		isAsyncEnabled:   false,
		isSkipping:       false,
		beforeEach:       func() {},
		afterEach:        func() {},
		beforeAll:        func() {},
		afterAll:         func() {},
	}
}

func getTestName(t *testing.T) string {
	// `2` should be where CreateTest() is called
	_, fileFullPath, _, ok := runtime.Caller(2)
	if !ok {
		t.Fatal("Failed to start gest: Failed to get runtime caller information")
	}
	path, err := os.Getwd()
	if err != nil {
		t.Fatal("Failed to start gest:", err)
	}

	fileProjectPath := strings.Replace(fileFullPath, path+"/", "", 1)

	fullName := t.Name()
	parts := strings.Split(fullName, "/")
	testFuncName := parts[0]
	return testFuncName + "@" + fileProjectPath
}

func (t *Test) BeforeAll(body func()) {
	t.beforeAll = body
}

func (t *Test) BeforeEach(body func()) {
	t.beforeEach = body
}

func (t *Test) AfterEach(body func()) {
	t.afterEach = body
}

func (t *Test) AfterAll(body func()) {
	t.afterAll = body
}

func (t *Test) Describe(description string, body func()) {
	if t.isSkipping {
		t.gestOutput(t.describeFuncSkipMsg(description))
		defer t.disableSkipping()
		return
	}

	start := time.Now()
	t.beforeAll()
	t.testingT.Run(description, func(testingT *testing.T) {
		body()
	})
	t.afterAll()
	elapsed := time.Since(start)

	// if not executed from gest command, just run test and NO Gest OUTPUT
	if !t.isGestOn {
		return
	}

	messageFunc := t.describeFuncPassMsg
	if t.isAnyTestFailed {
		messageFunc = t.describeFuncFailMsg
	}
	t.gestOutput(messageFunc(description, elapsed))

	for _, msg := range t.messages {
		t.gestOutput(msg)
	}
}

func (t *Test) describeFuncFailMsg(description string, elapsed time.Duration) string {
	timeInSeconds := fmt.Sprintf("%.3f", elapsed.Seconds())
	line1 := RedMsg(fmt.Sprintf("%s (Asserted: %d/%d)", t.testName, t.passed, t.subtotal))
	line2 := RedMsg(fmt.Sprintf(" ✘ FAIL: describe \"%s\"  (%ss)", description, timeInSeconds))
	return line1 + "\n" + line2
}

func (t *Test) describeFuncPassMsg(description string, elapsed time.Duration) string {
	timeInSeconds := fmt.Sprintf("%.3f", elapsed.Seconds())
	line1 := GreenMsg(fmt.Sprintf("%s (Asserted: %d/%d)", t.testName, t.passed, t.subtotal))
	line2 := GreenMsg(fmt.Sprintf(" ✔ PASS: describe \"%s\"  (%ss)", description, timeInSeconds))
	return line1 + "\n" + line2
}

func (t *Test) describeFuncSkipMsg(description string) string {
	line1 := YellowMsg(t.testName)
	line2 := YellowMsg(" ➔ SKIP: describe \"" + description + "\"")
	return line1 + "\n" + line2
}

func (t *Test) It(description string, body func()) {
	if t.isSkipping {
		t.messages = append(t.messages, itFuncSkipMsg(description))
		defer t.disableSkipping()
		return
	}

	start := time.Now()
	t.beforeEach()
	t.testingT.Run(description, func(testingT *testing.T) {
		if t.isAsyncEnabled {
			t.testingT.Parallel()
		}
		body()
	})
	t.disableAsync()
	t.afterEach()
	elapsed := time.Since(start)

	messageFunc := itFuncPassMsg
	if t.isThisTestFailed {
		messageFunc = itFuncFailMsg
	}
	t.messages = append(
		t.messages,
		messageFunc(description, elapsed),
	)

	// reset flag
	t.isThisTestFailed = false
}

// TEST: Parallel()が本当にできているかテスト
func (t *Test) Async() *Test {
	t.isAsyncEnabled = true
	return t
}

func (t *Test) Skip() *Test {
	t.isSkipping = true
	return t
}

func (t *Test) Todo(description string) {
	// FIXME: Describe()の"中"で呼ばなければメッセージが出力されない
	t.messages = append(t.messages, YellowMsg("    - todo: \""+description+"\""))
}

func (t *Test) disableAsync() {
	t.isAsyncEnabled = false
}

func (t *Test) disableSkipping() {
	t.isSkipping = false
}

// Output to console only when GESTON env is set
func (t *Test) gestOutput(msg ...string) (n int, err error) {
	if t.isGestOn {
		return fmt.Println(strings.Join(msg, ""))
	}
	return 0, nil
}

// ---------- ALIASES ----------

// Alias of testing.T.TempDir()
func (t *Test) TempDir() {
	t.testingT.TempDir()
}

// Alias of testing.T.Setenv()
func (t *Test) Setenv(key, value string) {
	t.testingT.Setenv(key, value)
}

// It's just an alias of `Skip()` function
// Do not call this within Describe() or It()
// Call this directory within Test function
func (t *Test) SkipAll() {
	t.testingT.Skip()
}

// TODO: Add other convenient aliases
