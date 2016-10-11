package scanner

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWalkLog(t *testing.T) {
	sampleLog := `
  ██████╗ ██╗████████╗██████╗ ██╗███████╗███████╗
  ██╔══██╗██║╚══██╔══╝██╔══██╗██║██╔════╝██╔════╝
  ██████╔╝██║   ██║   ██████╔╝██║███████╗█████╗
  ██╔══██╗██║   ██║   ██╔══██╗██║╚════██║██╔══╝
  ██████╔╝██║   ██║   ██║  ██║██║███████║███████╗
  ╚═════╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝

Version: 1.4.0

INFO[23:04:14] Running workflow: test

INFO[23:04:14] Switching to workflow: test

INFO[23:04:14] Step uses latest version -- Updating StepLib ...
INFO[23:04:14] Update StepLib (https://github.com/bitrise-io/bitrise-steplib.git)...
Already up-to-date.
+------------------------------------------------------------------------------+
| (0) first step                                                               |
+------------------------------------------------------------------------------+
| id: script                                                                   |
| version: 1.1.3                                                               |
| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
| toolkit: bash                                                                |
| time: 2016-10-11T23:04:16+02:00                                              |
+------------------------------------------------------------------------------+
|                                                                              |
first step
|                                                                              |
+----+--------------------------------------------------------------+----------+
| ✅  | first step                                                   | 2.2 sec  |
+----+--------------------------------------------------------------+----------+

                                          ▼

+------------------------------------------------------------------------------+
| (1) script                                                                   |
+------------------------------------------------------------------------------+
| id: script                                                                   |
| version: 1.1.3                                                               |
| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
| toolkit: bash                                                                |
| time: 2016-10-11T23:04:17+02:00                                              |
+------------------------------------------------------------------------------+
|                                                                              |
second step
|                                                                              |
+----+--------------------------------------------------------------+----------+
| ✅  | script                                                       | 0.47 sec |
+----+--------------------------------------------------------------+----------+


+------------------------------------------------------------------------------+
|                               bitrise summary                                |
+----+--------------------------------------------------------------+----------+
|    | title                                                        | time (s) |
+----+--------------------------------------------------------------+----------+
| ✅  | first step                                                   | 2.2 sec  |
+----+--------------------------------------------------------------+----------+
| ✅  | script                                                       | 0.47 sec |
+----+--------------------------------------------------------------+----------+
| Total runtime: 2.7 sec                                                       |
+------------------------------------------------------------------------------+

WARN[23:04:17]
WARN[23:04:17] New version (0.9.5) of plugin (analytics) available
INFO[23:04:17]
INFO[23:04:17] Submitting anonymized usage information...
INFO[23:04:17] For more information visit:
INFO[23:04:17] https://github.com/bitrise-core/bitrise-plugins-analytics/blob/master/README.md`

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
[BeforeFirstStep]Version: 1.4.0
[BeforeFirstStep]
[BeforeFirstStep]INFO[23:04:14] Running workflow: test
[BeforeFirstStep]
[BeforeFirstStep]INFO[23:04:14] Switching to workflow: test
[BeforeFirstStep]
[BeforeFirstStep]INFO[23:04:14] Step uses latest version -- Updating StepLib ...
[BeforeFirstStep]INFO[23:04:14] Update StepLib (https://github.com/bitrise-io/bitrise-steplib.git)...
[BeforeFirstStep]Already up-to-date.
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[StepInfoHeader]| (0) first step                                                               |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]| id: script                                                                   |
[StepInfoHeader]| version: 1.1.3                                                               |
[StepInfoHeader]| collection: https://github.com/bitrise-io/bitrise-steplib.git                |
[StepInfoHeader]| toolkit: bash                                                                |
[StepInfoHeader]| time: 2016-10-11T23:04:16+02:00                                              |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]|                                                                              |
[StepLog]first step
[StepInfoFooter]|                                                                              |
[StepInfoFooter]+----+--------------------------------------------------------------+----------+
[StepInfoFooter]| ✅  | first step                                                   | 2.2 sec  |
[StepInfoFooter]+----+--------------------------------------------------------------+----------+
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
[StepInfoHeader]| time: 2016-10-11T23:04:17+02:00                                              |
[StepInfoHeader]+------------------------------------------------------------------------------+
[StepInfoHeader]|                                                                              |
[StepLog]second step
[StepInfoFooter]|                                                                              |
[StepInfoFooter]+----+--------------------------------------------------------------+----------+
[StepInfoFooter]| ✅  | script                                                       | 0.47 sec |
[StepInfoFooter]+----+--------------------------------------------------------------+----------+
[BetweenSteps]
[BetweenSteps]
[StepInfoHeaderOrBuildSummarySectionStarter]+------------------------------------------------------------------------------+
[BuildSummary]|                               bitrise summary                                |
[BuildSummary]+----+--------------------------------------------------------------+----------+
[BuildSummary]|    | title                                                        | time (s) |
[BuildSummary]+----+--------------------------------------------------------------+----------+
[BuildSummary]| ✅  | first step                                                   | 2.2 sec  |
[BuildSummary]+----+--------------------------------------------------------------+----------+
[BuildSummary]| ✅  | script                                                       | 0.47 sec |
[BuildSummary]+----+--------------------------------------------------------------+----------+
[BuildSummary]| Total runtime: 2.7 sec                                                       |
[BuildSummary]+------------------------------------------------------------------------------+
[AfterBuildSummary]
[AfterBuildSummary]WARN[23:04:17]
[AfterBuildSummary]WARN[23:04:17] New version (0.9.5) of plugin (analytics) available
[AfterBuildSummary]INFO[23:04:17]
[AfterBuildSummary]INFO[23:04:17] Submitting anonymized usage information...
[AfterBuildSummary]INFO[23:04:17] For more information visit:
[AfterBuildSummary]INFO[23:04:17] https://github.com/bitrise-core/bitrise-plugins-analytics/blob/master/README.md
`
	require.Equal(t, expectedResult, result)
}
