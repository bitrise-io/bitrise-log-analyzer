package scanner

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWalkLog(t *testing.T) {
	t.Log("Simple log")
	{
		sampleLog := `
  ██████╗ ██╗████████╗██████╗ ██╗███████╗███████╗
  ██╔══██╗██║╚══██╔══╝██╔══██╗██║██╔════╝██╔════╝
  ██████╔╝██║   ██║   ██████╔╝██║███████╗█████╗
  ██╔══██╗██║   ██║   ██╔══██╗██║╚════██║██╔══╝
  ██████╔╝██║   ██║   ██║  ██║██║███████║███████╗
  ╚═════╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝

Version: 1.4.3

INFO[19:17:57] Running workflow: wf-two

INFO[19:17:57] Switching to workflow: wf-two

INFO[19:17:58] Step uses latest version -- Updating StepLib ...
INFO[19:17:58] Update StepLib (https://github.com/bitrise-io/bitrise-steplib.git)...
Already up-to-date.
+------------------------------------------------------------------------------+
| (0) wf2 - first step                                                         |
+------------------------------------------------------------------------------+
| id: script                                                                   |
| version: 1.1.3                                                               |
| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
| toolkit: bash                                                                |
| time: 2016-11-01T19:17:59+01:00                                              |
+------------------------------------------------------------------------------+
|                                                                              |
wf2 - first step
|                                                                              |
+---+---------------------------------------------------------------+----------+
| ✓ | wf2 - first step                                              | 2.18 sec |
+---+---------------------------------------------------------------+----------+

                                          ▼

+------------------------------------------------------------------------------+
| (1) script                                                                   |
+------------------------------------------------------------------------------+
| id: script                                                                   |
| version: 1.1.3                                                               |
| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
| toolkit: bash                                                                |
| time: 2016-11-01T19:18:00+01:00                                              |
+------------------------------------------------------------------------------+
|                                                                              |
wf2 - second step
|                                                                              |
+---+---------------------------------------------------------------+----------+
| ✓ | script                                                        | 0.40 sec |
+---+---------------------------------------------------------------+----------+


+------------------------------------------------------------------------------+
|                               bitrise summary                                |
+---+---------------------------------------------------------------+----------+
|   | title                                                         | time (s) |
+---+---------------------------------------------------------------+----------+
| ✓ | wf2 - first step                                              | 2.18 sec |
+---+---------------------------------------------------------------+----------+
| ✓ | script                                                        | 0.40 sec |
+---+---------------------------------------------------------------+----------+
| Total runtime: 2.58 sec                                                      |
+------------------------------------------------------------------------------+

INFO[19:18:00]
INFO[19:18:00] Submitting anonymized usage information...
INFO[19:18:00] For more information visit:
INFO[19:18:00] https://github.com/bitrise-core/bitrise-plugins-analytics/blob/master/README.md`

		result := ""
		err := WalkLog(strings.NewReader(sampleLog), func(line string, lineType LogLineType) {
			result = result + fmt.Sprintf("[%s]%s\n", lineType, line)
		})
		require.NoError(t, err)

		expectedResult := `[BeforeFirstStep]
[BeforeFirstStep]  ██████╗ ██╗████████╗██████╗ ██╗███████╗███████╗
[BeforeFirstStep]  ██╔══██╗██║╚══██╔══╝██╔══██╗██║██╔════╝██╔════╝
[BeforeFirstStep]  ██████╔╝██║   ██║   ██████╔╝██║███████╗█████╗
[BeforeFirstStep]  ██╔══██╗██║   ██║   ██╔══██╗██║╚════██║██╔══╝
[BeforeFirstStep]  ██████╔╝██║   ██║   ██║  ██║██║███████║███████╗
[BeforeFirstStep]  ╚═════╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝
[BeforeFirstStep]
[BeforeFirstStep]Version: 1.4.3
[BeforeFirstStep]
[BeforeFirstStep]INFO[19:17:57] Running workflow: wf-two
[BeforeFirstStep]
[BeforeFirstStep]INFO[19:17:57] Switching to workflow: wf-two
[BeforeFirstStep]
[BeforeFirstStep]INFO[19:17:58] Step uses latest version -- Updating StepLib ...
[BeforeFirstStep]INFO[19:17:58] Update StepLib (https://github.com/bitrise-io/bitrise-steplib.git)...
[BeforeFirstStep]Already up-to-date.
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[StepInfoHeader]| (0) wf2 - first step                                                         |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]| id: script                                                                   |
[StepInfoHeader]| version: 1.1.3                                                               |
[StepInfoHeader]| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
[StepInfoHeader]| toolkit: bash                                                                |
[StepInfoHeader]| time: 2016-11-01T19:17:59+01:00                                              |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]|                                                                              |
[StepLog]wf2 - first step
[StepInfoFooter]|                                                                              |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[StepInfoFooter]| ✓ | wf2 - first step                                              | 2.18 sec |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[BetweenSteps]
[BetweenSteps]                                          ▼
[BetweenSteps]
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[StepInfoHeader]| (1) script                                                                   |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]| id: script                                                                   |
[StepInfoHeader]| version: 1.1.3                                                               |
[StepInfoHeader]| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
[StepInfoHeader]| toolkit: bash                                                                |
[StepInfoHeader]| time: 2016-11-01T19:18:00+01:00                                              |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]|                                                                              |
[StepLog]wf2 - second step
[StepInfoFooter]|                                                                              |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[StepInfoFooter]| ✓ | script                                                        | 0.40 sec |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[BetweenSteps]
[BetweenSteps]
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[BuildSummary]|                               bitrise summary                                |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]|   | title                                                         | time (s) |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]| ✓ | wf2 - first step                                              | 2.18 sec |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]| ✓ | script                                                        | 0.40 sec |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]| Total runtime: 2.58 sec                                                      |
[BuildSummary]+------------------------------------------------------------------------------+
[AfterBuildSummary]
[AfterBuildSummary]INFO[19:18:00]
[AfterBuildSummary]INFO[19:18:00] Submitting anonymized usage information...
[AfterBuildSummary]INFO[19:18:00] For more information visit:
[AfterBuildSummary]INFO[19:18:00] https://github.com/bitrise-core/bitrise-plugins-analytics/blob/master/README.md
`

		fmt.Println()
		fmt.Println("Expected result:")
		fmt.Println(expectedResult)
		fmt.Println("------------")
		fmt.Println()

		fmt.Println()
		fmt.Println("The result:")
		fmt.Println(result)
		fmt.Println("------------")
		fmt.Println()

		require.Equal(t, expectedResult, result)
	}

	t.Log("Bitrise run in Bitrise run")
	{
		sampleLog := `

  ██████╗ ██╗████████╗██████╗ ██╗███████╗███████╗
  ██╔══██╗██║╚══██╔══╝██╔══██╗██║██╔════╝██╔════╝
  ██████╔╝██║   ██║   ██████╔╝██║███████╗█████╗
  ██╔══██╗██║   ██║   ██╔══██╗██║╚════██║██╔══╝
  ██████╔╝██║   ██║   ██║  ██║██║███████║███████╗
  ╚═════╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝

Version: 1.4.3

INFO[19:04:27] Running workflow: wf-one

INFO[19:04:27] Switching to workflow: wf-one

INFO[19:04:27] Step uses latest version -- Updating StepLib ...
INFO[19:04:27] Update StepLib (https://github.com/bitrise-io/bitrise-steplib.git)...
Already up-to-date.
+------------------------------------------------------------------------------+
| (0) wf1 - first step                                                         |
+------------------------------------------------------------------------------+
| id: script                                                                   |
| version: 1.1.3                                                               |
| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
| toolkit: bash                                                                |
| time: 2016-11-01T19:04:30+01:00                                              |
+------------------------------------------------------------------------------+
|                                                                              |
wf1 - first step
|                                                                              |
+---+---------------------------------------------------------------+----------+
| ✓ | wf1 - first step                                              | 2.81 sec |
+---+---------------------------------------------------------------+----------+

                                          ▼

+------------------------------------------------------------------------------+
| (1) script                                                                   |
+------------------------------------------------------------------------------+
| id: script                                                                   |
| version: 1.1.3                                                               |
| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
| toolkit: bash                                                                |
| time: 2016-11-01T19:04:30+01:00                                              |
+------------------------------------------------------------------------------+
|                                                                              |
wf1 - second step
|                                                                              |
+---+---------------------------------------------------------------+----------+
| ✓ | script                                                        | 0.46 sec |
+---+---------------------------------------------------------------+----------+

                                          ▼

+------------------------------------------------------------------------------+
| (2) script                                                                   |
+------------------------------------------------------------------------------+
| id: script                                                                   |
| version: 1.1.3                                                               |
| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
| toolkit: bash                                                                |
| time: 2016-11-01T19:04:31+01:00                                              |
+------------------------------------------------------------------------------+
|                                                                              |
+ bitrise run wf-two

  ██████╗ ██╗████████╗██████╗ ██╗███████╗███████╗
  ██╔══██╗██║╚══██╔══╝██╔══██╗██║██╔════╝██╔════╝
  ██████╔╝██║   ██║   ██████╔╝██║███████╗█████╗
  ██╔══██╗██║   ██║   ██╔══██╗██║╚════██║██╔══╝
  ██████╔╝██║   ██║   ██║  ██║██║███████║███████╗
  ╚═════╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝

Version: 1.4.3

INFO[19:04:31] Running workflow: wf-two

INFO[19:04:31] Switching to workflow: wf-two

INFO[19:04:31] Step uses latest version -- Updating StepLib ...
INFO[19:04:31] Update StepLib (https://github.com/bitrise-io/bitrise-steplib.git)...
Already up-to-date.
+------------------------------------------------------------------------------+
| (0) wf2 - first step                                                         |
+------------------------------------------------------------------------------+
| id: script                                                                   |
| version: 1.1.3                                                               |
| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
| toolkit: bash                                                                |
| time: 2016-11-01T19:04:34+01:00                                              |
+------------------------------------------------------------------------------+
|                                                                              |
wf2 - first step
|                                                                              |
+---+---------------------------------------------------------------+----------+
| ✓ | wf2 - first step                                              | 2.79 sec |
+---+---------------------------------------------------------------+----------+

                                          ▼

+------------------------------------------------------------------------------+
| (1) script                                                                   |
+------------------------------------------------------------------------------+
| id: script                                                                   |
| version: 1.1.3                                                               |
| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
| toolkit: bash                                                                |
| time: 2016-11-01T19:04:34+01:00                                              |
+------------------------------------------------------------------------------+
|                                                                              |
wf2 - second step
|                                                                              |
+---+---------------------------------------------------------------+----------+
| ✓ | script                                                        | 0.46 sec |
+---+---------------------------------------------------------------+----------+

+------------------------------------------------------------------------------+
|                               bitrise summary                                |
+---+---------------------------------------------------------------+----------+
|   | title                                                         | time (s) |
+---+---------------------------------------------------------------+----------+
| ✓ | wf2 - first step                                              | 2.79 sec |
+---+---------------------------------------------------------------+----------+
| ✓ | script                                                        | 0.46 sec |
+---+---------------------------------------------------------------+----------+
| Total runtime: 3.24 sec                                                      |
+------------------------------------------------------------------------------+

INFO[19:04:35]
INFO[19:04:35] Submitting anonymized usage information...
INFO[19:04:35] For more information visit:
INFO[19:04:35] https://github.com/bitrise-core/bitrise-plugins-analytics/blob/master/README.md
|                                                                              |
+---+---------------------------------------------------------------+----------+
| ✓ | script                                                        | 4.63 sec |
+---+---------------------------------------------------------------+----------+

+------------------------------------------------------------------------------+
|                               bitrise summary                                |
+---+---------------------------------------------------------------+----------+
|   | title                                                         | time (s) |
+---+---------------------------------------------------------------+----------+
| ✓ | wf1 - first step                                              | 2.81 sec |
+---+---------------------------------------------------------------+----------+
| ✓ | script                                                        | 0.46 sec |
+---+---------------------------------------------------------------+----------+
| ✓ | script                                                        | 4.63 sec |
+---+---------------------------------------------------------------+----------+
| Total runtime: 7.91 sec                                                      |
+------------------------------------------------------------------------------+

INFO[19:04:35]
INFO[19:04:35] Submitting anonymized usage information...
INFO[19:04:35] For more information visit:
INFO[19:04:35] https://github.com/bitrise-core/bitrise-plugins-analytics/blob/master/README.md`

		result := ""
		err := WalkLog(strings.NewReader(sampleLog), func(line string, lineType LogLineType) {
			result = result + fmt.Sprintf("[%s]%s\n", lineType, line)
		})
		require.NoError(t, err)

		expectedResult := `[BeforeFirstStep]
[BeforeFirstStep]
[BeforeFirstStep]  ██████╗ ██╗████████╗██████╗ ██╗███████╗███████╗
[BeforeFirstStep]  ██╔══██╗██║╚══██╔══╝██╔══██╗██║██╔════╝██╔════╝
[BeforeFirstStep]  ██████╔╝██║   ██║   ██████╔╝██║███████╗█████╗
[BeforeFirstStep]  ██╔══██╗██║   ██║   ██╔══██╗██║╚════██║██╔══╝
[BeforeFirstStep]  ██████╔╝██║   ██║   ██║  ██║██║███████║███████╗
[BeforeFirstStep]  ╚═════╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝
[BeforeFirstStep]
[BeforeFirstStep]Version: 1.4.3
[BeforeFirstStep]
[BeforeFirstStep]INFO[19:04:27] Running workflow: wf-one
[BeforeFirstStep]
[BeforeFirstStep]INFO[19:04:27] Switching to workflow: wf-one
[BeforeFirstStep]
[BeforeFirstStep]INFO[19:04:27] Step uses latest version -- Updating StepLib ...
[BeforeFirstStep]INFO[19:04:27] Update StepLib (https://github.com/bitrise-io/bitrise-steplib.git)...
[BeforeFirstStep]Already up-to-date.
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[StepInfoHeader]| (0) wf1 - first step                                                         |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]| id: script                                                                   |
[StepInfoHeader]| version: 1.1.3                                                               |
[StepInfoHeader]| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
[StepInfoHeader]| toolkit: bash                                                                |
[StepInfoHeader]| time: 2016-11-01T19:04:30+01:00                                              |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]|                                                                              |
[StepLog]wf1 - first step
[StepInfoFooter]|                                                                              |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[StepInfoFooter]| ✓ | wf1 - first step                                              | 2.81 sec |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[BetweenSteps]
[BetweenSteps]                                          ▼
[BetweenSteps]
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[StepInfoHeader]| (1) script                                                                   |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]| id: script                                                                   |
[StepInfoHeader]| version: 1.1.3                                                               |
[StepInfoHeader]| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
[StepInfoHeader]| toolkit: bash                                                                |
[StepInfoHeader]| time: 2016-11-01T19:04:30+01:00                                              |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]|                                                                              |
[StepLog]wf1 - second step
[StepInfoFooter]|                                                                              |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[StepInfoFooter]| ✓ | script                                                        | 0.46 sec |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[BetweenSteps]
[BetweenSteps]                                          ▼
[BetweenSteps]
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[StepInfoHeader]| (2) script                                                                   |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]| id: script                                                                   |
[StepInfoHeader]| version: 1.1.3                                                               |
[StepInfoHeader]| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
[StepInfoHeader]| toolkit: bash                                                                |
[StepInfoHeader]| time: 2016-11-01T19:04:31+01:00                                              |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]|                                                                              |
[StepLog]+ bitrise run wf-two
[StepLog]
[StepLog]  ██████╗ ██╗████████╗██████╗ ██╗███████╗███████╗
[StepLog]  ██╔══██╗██║╚══██╔══╝██╔══██╗██║██╔════╝██╔════╝
[StepLog]  ██████╔╝██║   ██║   ██████╔╝██║███████╗█████╗
[StepLog]  ██╔══██╗██║   ██║   ██╔══██╗██║╚════██║██╔══╝
[StepLog]  ██████╔╝██║   ██║   ██║  ██║██║███████║███████╗
[StepLog]  ╚═════╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝
[StepLog]
[StepLog]Version: 1.4.3
[StepLog]
[StepLog]INFO[19:04:31] Running workflow: wf-two
[StepLog]
[StepLog]INFO[19:04:31] Switching to workflow: wf-two
[StepLog]
[StepLog]INFO[19:04:31] Step uses latest version -- Updating StepLib ...
[StepLog]INFO[19:04:31] Update StepLib (https://github.com/bitrise-io/bitrise-steplib.git)...
[StepLog]Already up-to-date.
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[StepInfoHeader]| (0) wf2 - first step                                                         |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]| id: script                                                                   |
[StepInfoHeader]| version: 1.1.3                                                               |
[StepInfoHeader]| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
[StepInfoHeader]| toolkit: bash                                                                |
[StepInfoHeader]| time: 2016-11-01T19:04:34+01:00                                              |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]|                                                                              |
[StepLog]wf2 - first step
[StepInfoFooter]|                                                                              |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[StepInfoFooter]| ✓ | wf2 - first step                                              | 2.79 sec |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[BetweenSteps]
[BetweenSteps]                                          ▼
[BetweenSteps]
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[StepInfoHeader]| (1) script                                                                   |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]| id: script                                                                   |
[StepInfoHeader]| version: 1.1.3                                                               |
[StepInfoHeader]| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
[StepInfoHeader]| toolkit: bash                                                                |
[StepInfoHeader]| time: 2016-11-01T19:04:34+01:00                                              |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]|                                                                              |
[StepLog]wf2 - second step
[StepInfoFooter]|                                                                              |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[StepInfoFooter]| ✓ | script                                                        | 0.46 sec |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[BetweenSteps]
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[BuildSummary]|                               bitrise summary                                |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]|   | title                                                         | time (s) |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]| ✓ | wf2 - first step                                              | 2.79 sec |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]| ✓ | script                                                        | 0.46 sec |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]| Total runtime: 3.24 sec                                                      |
[BuildSummary]+------------------------------------------------------------------------------+
[AfterBuildSummary]
[AfterBuildSummary]INFO[19:04:35]
[AfterBuildSummary]INFO[19:04:35] Submitting anonymized usage information...
[AfterBuildSummary]INFO[19:04:35] For more information visit:
[AfterBuildSummary]INFO[19:04:35] https://github.com/bitrise-core/bitrise-plugins-analytics/blob/master/README.md
[StepInfoFooter]|                                                                              |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[StepInfoFooter]| ✓ | script                                                        | 4.63 sec |
[StepInfoFooter]+---+---------------------------------------------------------------+----------+
[BetweenSteps]
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[BuildSummary]|                               bitrise summary                                |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]|   | title                                                         | time (s) |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]| ✓ | wf1 - first step                                              | 2.81 sec |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]| ✓ | script                                                        | 0.46 sec |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]| ✓ | script                                                        | 4.63 sec |
[BuildSummary]+---+---------------------------------------------------------------+----------+
[BuildSummary]| Total runtime: 7.91 sec                                                      |
[BuildSummary]+------------------------------------------------------------------------------+
[AfterBuildSummary]
[AfterBuildSummary]INFO[19:04:35]
[AfterBuildSummary]INFO[19:04:35] Submitting anonymized usage information...
[AfterBuildSummary]INFO[19:04:35] For more information visit:
[AfterBuildSummary]INFO[19:04:35] https://github.com/bitrise-core/bitrise-plugins-analytics/blob/master/README.md
`
		fmt.Println()
		fmt.Println("Expected result:")
		fmt.Println(expectedResult)
		fmt.Println("------------")
		fmt.Println()

		fmt.Println()
		fmt.Println("The result:")
		fmt.Println(result)
		fmt.Println("------------")
		fmt.Println()
		require.Equal(t, expectedResult, result)
	}
}
