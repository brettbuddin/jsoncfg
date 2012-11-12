package jsoncfg

import (
    "os"
    "io"
    "bytes"
    "bufio"
    "strings"
    "encoding/json"
    "errors"
)

func LoadString(input string) (*Config, error) {
    data, err := decode(strings.NewReader(input))

    if err != nil {
        return &Config{nil}, err
    }

    return &Config{data}, nil
}

func LoadFile(path string) (*Config, error) {
    var (
        file *os.File
        err  error
    )

    if file, err = os.Open(path); err != nil {
        return &Config{nil}, err
    }

    defer file.Close()

    data, err := decode(bufio.NewReader(file))

    if err != nil {
        return &Config{nil}, err
    }

    return &Config{data}, nil
}

func decode(reader io.Reader) (map[string]interface{}, error) {
    var payload map[string]interface{}

    decoder := json.NewDecoder(reader)

    if err := decoder.Decode(&payload); err != nil {
        return payload, err
    }

    return payload, nil
}

type Config struct {
    raw interface{}
}

func (c *Config) Get(key string) *Config {
    m, err := c.Map()

    if err == nil {
        if val, ok := m[key]; ok {
            return &Config{val}
        }
    }

    return &Config{nil}
}

func (c *Config) String() (string, error) {
    if s, ok := c.raw.(string); ok {
        return s, nil
    }

    return "", errors.New("type assertion to string failed")
}

func (c *Config) Bytes() ([]byte, error) {
    if s, ok := c.raw.(string); ok {
        return []byte(s), nil
    }

    return nil, errors.New("type assertion to bytes failed")
}

func (c *Config) Int() (int, error) {
    if f, ok := c.raw.(float64); ok {
        return int(f), nil
    }

    return -1, errors.New("type assertion to integer failed")
}

func (c *Config) Float() (float64, error) {
    if f, ok := c.raw.(float64); ok {
        return f, nil
    }

    return -1, errors.New("type assertion to float64 failed")
}

func (c *Config) Array() ([]interface{}, error) {
    if a, ok := c.raw.([]interface{}); ok {
        return a, nil
    }

    return nil, errors.New("type assertion to array failed")
}

func (c *Config) Map() (map[string]interface{}, error) {
    if m, ok := c.raw.(map[string]interface{}); ok {
        return m, nil
    }

    return nil, errors.New("type assertion to map failed")
}

func (c *Config) SaveToFile(path string) error {
    var (
        file *os.File
        err  error
    )

    if file, err = os.Create(path); err != nil {
        return err
    }

    defer file.Close()

    data, _ := json.Marshal(c.raw)

    if err != nil {
        return err
    }

    output := bytes.NewBuffer(make([]byte, 0))

    json.Indent(output, data, "", "    ")

    _, err = file.WriteString(output.String())

    if err != nil {
        return err
    }

    return nil
}
