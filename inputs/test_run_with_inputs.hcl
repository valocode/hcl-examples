
input "tool" {
    type = object({
        name = string,
        type = string
    })
}

inputs {
    tool = {
        name = "gotest"
        type = "unit_test"
    }
}

test_run {
    tool = input.tool.name
    type = input.tool.type
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
