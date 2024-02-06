package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/yrichika/gest/pkg/gt"
)

type FlagHolder struct {
	Verbose            bool
	VeryVerbose        bool
	Coverage           bool
	CoverageProfileDir string
	RunOnly            string
	RunInAllDirs       bool
}

func flags() *FlagHolder {
	var verbose bool
	var veryVerbose bool
	var coverage bool
	var coverageProfileDir string
	var runOnly string
	var runInAllDirs bool
	const defaultCoverageProfileDir = "gest_coverage"
	flag.BoolVar(&verbose, "v", false, "verbose: this also logs \"go test\" command outputs")
	flag.BoolVar(&veryVerbose, "vv", false, "very verbose: this also logs \"go test -v\" command outputs. If you don't see some output you implemented for debugging like Println(), use this option.")
	flag.BoolVar(&coverage, "cover", false, "create coverage profile")
	flag.StringVar(&coverageProfileDir, "coverprofile", defaultCoverageProfileDir, "specify name of coverage profile output directory")
	flag.StringVar(&runOnly, "run", "", "run only tests matching the regular expression.")
	flag.BoolVar(&runInAllDirs, "all-dirs", false, "run all tests including hidden directories and other directories usually not targeted by 'go test ./...'. This is useful when you want to test all directories from the project root.")
	flag.Parse()

	return &FlagHolder{
		verbose,
		veryVerbose,
		coverage,
		coverageProfileDir,
		runOnly,
		runInAllDirs,
	}
}

func (f *FlagHolder) vOrVV() bool {
	return f.Verbose || f.VeryVerbose
}

const warningNotProjectRootMsg = "RECOMMENDATION: Please execute this command at the project root where 'go.mod' file exists. Gest will recursively test all directories from the project root."
const errCoverageMustBeAtProjectRootMsg = "ERROR: go.mod file not found. Please execute this command at the project root where 'go.mod' file exists for coverage. Coverage uses simple 'go test ./... -cover' command from the project root."
const errUnexpectedMsg = "Failed to run gest: unexpected error: "

func main() {

	flags := flags()
	_, err := os.Stat("go.mod")

	if flags.Coverage && os.IsNotExist(err) {
		log.Fatal(gt.RedMsg(errCoverageMustBeAtProjectRootMsg))
	}

	if flags.Coverage {
		createCoverageProfile(flags)
		os.Exit(0)
	}

	if os.IsNotExist(err) {
		fmt.Println(
			gt.YellowMsg(warningNotProjectRootMsg),
		)
	} else if err != nil {
		log.Fatal(gt.RedMsg(errUnexpectedMsg), err)
	}

	dirNames := gt.GetAllTestFileDirectories(flags.RunInAllDirs)
	var failedTestResultLines []string
	var passedTestResultLines []string
	var anyOtherOutput []string
	var gestOutputPrinted []bool

	for _, dirName := range dirNames {
		lines := runTestAt(dirName, flags)
		failed, passed, other, hasGestResult := assortOutput(lines, flags)
		failedTestResultLines = append(failedTestResultLines, failed...)
		passedTestResultLines = append(passedTestResultLines, passed...)
		anyOtherOutput = append(anyOtherOutput, other...)
		gestOutputPrinted = append(gestOutputPrinted, hasGestResult)
	}
	isTestRun := gt.InArray(true, gestOutputPrinted)

	finalOutput(
		flags,
		passedTestResultLines,
		failedTestResultLines,
		anyOtherOutput,
		isTestRun,
	)
}

func assortOutput(lines []string, flags *FlagHolder) ([]string, []string, []string, bool) {

	var failedTestResultLines []string
	var passedTestResultLines []string
	var anyOtherOutput []string
	var gestOutputPrinted []bool

	for _, line := range lines {
		isLineStartsWithEsc := strings.HasPrefix(line, "\033")
		isLineStartsWithFail := strings.HasPrefix(line, "FAIL")
		isLineStartsWithOk := strings.HasPrefix(line, "ok")
		// "FAIL"だけの文字列が、上の判定でのとりこぼしがあるため、再度除外するための判定
		isLineContainsFail := strings.Contains(line, "FAIL")

		switch {
		case isLineStartsWithEsc:
			// output immediately so that users can also see the progress
			fmt.Println(line)
			gestOutputPrinted = append(gestOutputPrinted, true)
		// exclude line with "FAIL" string only, but allow prefixed with "FAIL" and any other string
		case line != "FAIL" && isLineStartsWithFail:
			failedTestResultLines = append(failedTestResultLines, line)
		case isLineStartsWithOk:
			passedTestResultLines = append(passedTestResultLines, line)
		case excludeLineCondition(line):
			// this vOrVV enables to output strings containing "FAIL", that users intentionally output with Println() etc.
			if flags.vOrVV() || !isLineContainsFail {
				anyOtherOutput = append(anyOtherOutput, line)
			}
		default:
			if flags.vOrVV() {
				anyOtherOutput = append(anyOtherOutput, line)
			}
		}
	}
	hasGestResult := gt.InArray(true, gestOutputPrinted)

	return failedTestResultLines, passedTestResultLines, anyOtherOutput, hasGestResult
}

// these messages are standard `go test` command outputs.
// these outputs may confuse users, so exclude them.
func excludeLineCondition(line string) bool {
	return line != "" &&
		line != "PASS" &&
		!strings.HasPrefix(line, "===") &&
		!strings.HasPrefix(line, "---") &&
		!strings.HasPrefix(line, "exit") &&
		!strings.Contains(line, "testing: warning: no tests to run")
}

// run go test at the specified directory
// and returns the multiline output as []string each line
func runTestAt(dirName string, flags *FlagHolder) []string {

	os.Setenv(gt.EnvName, "true")
	cmd := exec.Command("go", "test")
	if flags.VeryVerbose {
		cmd = exec.Command("go", "test", "-v")
	}
	if flags.RunOnly != "" {
		cmd = exec.Command("go", "test", "-run", flags.RunOnly)
	}
	cmd.Dir = dirName
	out, err := cmd.CombinedOutput()
	if err != nil {
		// err string is usually something like "exit status 1"
		// not really necessary for now
	}

	return strings.Split(string(out), "\n")
}

const startingCoverageMsg = "Creating coverage profile... Executing 'go test -cover'"

func coverageErrMsg(dirName string) string {
	return "Error at creating coverage directory: [" + dirName + "]"
}

func coverageCreatedMsg(dirName string) string {
	return "✔ Created coverage profile at [" + dirName + ".html]!\n"
}

func createCoverageProfile(flags *FlagHolder) {
	fmt.Println(startingCoverageMsg)

	outputDir := flags.CoverageProfileDir
	errMkDir := os.MkdirAll(outputDir, 0755)
	if errMkDir != nil {
		log.Fatal(gt.RedMsg(coverageErrMsg(outputDir)), errMkDir)
	}

	t := time.Now()
	fileName := outputDir + "/profile_" + t.Format("20060102_150405")
	cmdOut := exec.Command("go", "test", "./...", "-v", "-cover", "-coverprofile="+fileName+".out")
	// Ignore error because when user's test fails, .Output returns error.
	outProfileCreated, _ := cmdOut.Output()
	fmt.Println(string(outProfileCreated))

	cmdHtml := exec.Command("go", "tool", "cover", "-html="+fileName+".out", "-o="+fileName+".html")
	outHtmlCreated, _ := cmdHtml.Output()
	fmt.Println(string(outHtmlCreated))

	fmt.Println(gt.GreenMsg(coverageCreatedMsg(fileName)))
}

const separatorOtherMsg = "\n----------========== OUTPUT ==========----------"
const separatorFailedMsg = "\n----------========== FAILED ==========----------"
const separatorPassedMsg = "\n----------========== PASSED ==========----------"
const noTestRunMsg = "\nNo tests to run.\n"

func finalOutput(
	flags *FlagHolder,
	passedTestResultLines []string,
	failedTestResultLines []string,
	anyOtherOutput []string,
	isTestRun bool,
) {
	if len(anyOtherOutput) > 0 {
		otherMsg := strings.Join(anyOtherOutput, "\n")
		fmt.Println(separatorOtherMsg + "\n" + otherMsg)
	}

	if !isTestRun && len(failedTestResultLines) == 0 {
		fmt.Println(gt.YellowMsg(noTestRunMsg))
		return
	}

	if len(failedTestResultLines) > 0 {
		failedMsg := strings.Join(failedTestResultLines, "\n")
		fmt.Println(gt.RedMsg(separatorFailedMsg + "\n" + failedMsg))
	}

	// omit final output if `-run` option is specified
	if flags.RunOnly != "" {
		return
	}

	if len(passedTestResultLines) > 0 {
		passedMsg := strings.Join(passedTestResultLines, "\n")
		fmt.Println(gt.GreenMsg(separatorPassedMsg + "\n" + passedMsg))
	}
}
