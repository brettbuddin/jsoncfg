package jsoncfg

import (
    "testing"
    //"log"
)

var cfg *Config

func init() {
    var err error

    testData := `
        {
            "data": {
                "aString": "value",
                "anInteger": 5,
                "aFloat": 2.1,
                "anArray": [1, 4, 3, 2]
            }
        }`

    cfg, err = LoadString(testData)

    if err != nil {
        panic(err)
    }
}

func TestStringData(t *testing.T) {
    actual, err := cfg.Get("data").Get("aString").String()

    if (err != nil) {
        t.Error(err)
    }

    expect := "value"

    if actual != expect {
        t.Errorf("%s is not equal to %s", actual, expect)
    }
}

func TestIntData(t *testing.T) {
    actual, err := cfg.Get("data").Get("anInteger").Int()

    if (err != nil) {
        t.Error(err)
    }

    expect := 5

    if actual != expect {
        t.Errorf("%s is not equal to %s", actual, expect)
    }
}

func TestFloatData(t *testing.T) {
    actual, err := cfg.Get("data").Get("aFloat").Float()

    if (err != nil) {
        t.Error(err)
    }

    expect := 2.1

    if actual != expect {
        t.Errorf("%s is not equal to %s", actual, expect)
    }
}

func TestArrayData(t *testing.T) {
    actual, err := cfg.Get("data").Get("anArray").Array()

    if (err != nil) {
        t.Error(err)
    }

    expect  := []float64{1, 4, 3, 2}
    noMatch := false

    for i, item := range(actual) {
        if item != expect[i] {
            noMatch = true
        }
    }

    if noMatch {
        t.Errorf("%s is not equal to %s", actual, expect)
    }
}

func TestToAndFromFile(t *testing.T) {
    filepath := "/tmp/TestSaveToFile-Output.json"
    err      := cfg.SaveToFile(filepath)

    if err != nil {
        t.Errorf("could not write to file", err)
    }

    cfg, err = LoadFile(filepath)

    if err != nil {
        panic(err)
    }

    actual, _ := cfg.Get("data").Get("aString").String()
    expect    := "value"

    if actual != expect {
        t.Error("%s does not match %s", actual, expect)
    }
}
