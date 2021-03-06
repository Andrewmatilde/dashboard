package versions

import (
	"fmt"
	"testing"
)

func TestParseVersionFromTokens(t *testing.T) {
	fmt.Println(GetVersions(&testBody))
}

var testBody = `## Please edit this comment to complete the following information

### Not a bug

1. Remove the 'type/bug' label
2. Add notes to indicate why it is not a bug

### Duplicate bug

1. Add the 'type/duplicate' label
2. Add the link to the original bug

### Bug

Note: Make Sure that 'component', and 'severity' labels are added
Example for how to fill out the template: https://github.com/pingcap/tidb/issues/20100

#### 1. Root Cause Analysis (RCA) (optional)
<!-- Write down the reason why this bug occurs -->

#### 2. Symptom (optional)

<!-- What will the user see when this bug occurs. The error message may be in the terminal, log or monitoring -->

#### 3. All Trigger Conditions (optional)

<!-- All the user scenarios that may trigger this bug -->

#### 4. Workaround (optional)

#### 5. Affected versions

<!--
In the format of [start_version:end_version], multiple version ranges are
accepted. If the bug only affects the unreleased version, please input:
"unreleased". For example:

Notes:
  1. Do not use any white spaces in '[]'.
  2. The range in '[]' is a closed interval
  3. The version format is ` + "`v$Major.$Minor.$Patch`" + `, the $Majoy and $Minor
	 number in a version range should be the same. [v3.0.1:v3.1.2] is
	 invalid because the $Minor number of the version range is different.

Example 1: [v3.0.1:v3.0.5], [v4.0.1:v4.0.5]
Example 2: unreleased
-->
[v3.0.0:v3.0.19] unreleased
#### 6. Fixed versions
v3.0.20master
<!--
The first released version that contains this fix in each minor version. If the bug's affected version has been released, the fixed version should be a detailed version number; If the bug doesn't affect any released version, the fixed version can be "master";

Example 1: v3.0.13, v4.0.5
Example 2: master
-->`
