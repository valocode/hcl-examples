
test_run {
    tool = "gotest"
    type = "unit_test"
    tests {
        test_case {
            name = "first test case"
            result = true
        }
        test_case {
            name = "second test case"
            result = false
        }
        test_case {
            name = "first test case"
            result = true
        }
    }
}
