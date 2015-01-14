<!SLIDE>

# JSON Encoding in Go

<!SLIDE>

# JSON Data Types

|        |                                            |
| ------ | ------------------------------------------ |
| String | `"hello"`                                  |
| Number | `42`                                       |
| True   | `true`                                     |
| False  | `false`                                    |
| Null   | `null`                                     |
| Object | `{ "foo": 42, "bar": true }`               |
| Array  | `[ 42, "hello", false, { "foo": "bar" } ]` |

<!SLIDE>

# Equivalent Go Types

|        |                                                  |
| ------ | ------------------------------------------------ |
| String | `"hello"`                                        |
| Number | `42`                                             |
| True   | `true`                                           |
| False  | `false`                                          |
| Null   | `nil`                                            |
| Object | `map[string]interface{}{"foo": 42, "bar": true}` |
| Array  | `[]interface{}{42, "hello", false}`              |

<!SLIDE methods>

# json.Marshal Method

---

<p class="code">
func Marshal

(v interface{})

([]byte, error)
</p>

---

Marshal returns the JSON encoding of v.

<!SLIDE methods>

# json.Unmarshal Method

---

<p class="code">
func Unmarshal

(data []byte, v interface{})

error
</p>

---

Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.

<!SLIDE>

# Examples

<!SLIDE methods>

# Encode Time

<p class="code">
t := time.Now()

str := t.Format(time.RFC3339)

data, err := json.Marshal(str)

=> "2015-01-13T16:02:27Z"
</p>

<!SLIDE methods>

# Decode Time

<p class="code">
b := []byte("2015-01-13T17:02:27Z")

var str string

err := json.Unmarshal(b, &str)

t := time.Parse(time.RFC3339, str)

=> "2015-01-13T17:02:27Z"
</p>

<!SLIDE methods>

# JSON Interfaces

---

<p class="code">
type Marshaler interface {
  MarshalJSON() ([]byte, error)
}
</p>

---

Marshaler is the interface implemented by objects that can marshal themselves into valid JSON.

<!SLIDE methods>

# JSON Interfaces

---

<p class="code">
type Unmarshaler interface {
  UnmarshalJSON([]byte) error
}
</p>

---

Unmarshaler is the interface implemented by objects that can unmarshal a JSON description of themselves.

<!SLIDE methods>

# Time

<p class="code" style="margin-top:0">
t := time.Now()

data, err := json.Marshal(t)

=> "2015-01-13T16:02:27Z"
</p>

---

<p class="code" style="margin-top:0">
b := []byte("2015-01-13T17:02:27Z")

err := json.Unmarshal(b, &t)

=> "2015-01-13T17:02:27Z"
</p>

<!SLIDE methods>

# Custom Type

<p class="code">
type Rect struct {
  X int
  Y int
}

func (r *Rect) MarshalJSON()
                ([]byte, error)

func (r *Rect) UnmarshalJSON([]byte)
                error
</p>

<!SLIDE methods>

# Tags

Struct fields can have tags containing arbitrary data.

By convention, tags are a concatenation of optionally space-separated key:"value" pairs

Tags can be read at runtime with the `reflect` package.

<!SLIDE methods>

# Tags

<p class="code">
type Rect struct {
  X int `side:"X"`
  Y int `side:"Y"`
}

typ := reflect.TypeOf(Rect{})
field := typ.Field(0)
field.Tag.Get("side")
=> "X"
</p>

<!SLIDE methods>

# JSON Tags

<p class="code">
type Rect struct {
  X int `json:"x"`
  Y int `json:"y"`
}

rect := Rect{10, 20}
data, err := json.Marshal(rect)
=> {"x":10,"y":20}

data = []byte(`{"x": 5, "y": 8}`)
err := json.Unmarshal(data, &rect)
=> Rect{X: 5, Y: 8}
</p>

<!SLIDE methods>

# JSON Tags

<p class="code" style="margin-top:0">
type RightCuboid struct {
  A Rect `json:"a"`
  B Rect `json:"b"`
  C Rect `json:"c"`
}

r := RightCuboid{
  Rect{10, 5}, Rect{5, 7}, Rect{7, 10},
}
data, err := json.Marshal(r)
=> {"a":{"x":10,"y":5},
"b":{"x":5,"y":7},"c":{"x":7,"y":10}}
</p>

<!SLIDE methods>

# JSON Encoder

An Encoder writes JSON objects to an output stream.

---

<p class="code">
func NewEncoder(w io.Writer)
               *Encoder

func (e *Encoder) Encode
                  (v interface{})
                  error
</p>

<!SLIDE methods>

# JSON Decoder

A Decoder reads and decodes JSON objects from an input stream.

---

<p class="code">
func NewDecoder(r io.Reader)
               *Decoder

func (dec *Decoder) Decode
                    (v interface{})
                    error
</p>

<!SLIDE methods>

# Delayed Decoding

`type RawMessage []byte`

RawMessage is a raw encoded JSON object. It implements Marshaler and Unmarshaler and can be used to delay JSON decoding or precompute a JSON encoding.

<!SLIDE methods>

# FIN
