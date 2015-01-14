<!SLIDE>

# JSON Encoding in Go

<!SLIDE>

# JSON Data Types

|        |                              |
| ------ | ---------------------------- |
| String | `"hello"`                    |
| Number | `42`                         |
| True   | `true`                       |
| False  | `false`                      |
| Null   | `null`                       |
| Object | `{ "foo": 42, "bar": true }` |
| Array  | `[ 42, "hello", false ]`     |

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

# Marshal

---

<p class="code">
func Marshal

(v interface{})

([]byte, error)
</p>

---

Marshal returns the JSON encoding of v.

<!SLIDE methods>

# Unmarshal

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
