package hidden_home

import(
  "fmt"
  "log"
  "testing"
  "os"
  "os/exec"
)

type TestConfig struct {
	Name string
}

func CreateUUID() (string, error) {
  //see http://stackoverflow.com/questions/15130321/is-there-a-method-to-generate-a-uuid-with-go-language
  //I don't need to be fast and I don't want to introduce a package
  //dependency
  if out, err := exec.Command("uuidgen").Output(); err == nil {
    result := fmt.Sprintf("%s", out)
    return result, err
  } else {
    return "", err
  }
}

func CleanUp(dirname string) error {
  return os.RemoveAll(dirname)
}

func TestCreate(t *testing.T) {
  //Case 1 - try to create ~/.polka by passing in "polka"
  fmt.Println("blah")
  if uuid1, err := CreateUUID(); err == nil {
    polka1 := fmt.Sprintf("%s-%s", "polka", uuid1 )
    log.Println("create 1")
    if err := Create(polka1); err != nil {
      t.Fail()
      t.Log(err)
    }
    CleanUp(ExpandTilde(polka1))
  } else {
    t.Fail()
  }

  log.Println("create 2")
  //Case 2 - try to create ~/.polka by passing in "~/.polka"
  if uuid2, err := CreateUUID(); err == nil {
    polka2 := fmt.Sprintf("~/.%v-%v", "polka", uuid2)
    if err := Create(polka2); err != nil {
      t.Fail()
    }
    CleanUp(ExpandTilde(polka2))
  } else {
    t.Fail()
  }

}

func TestSave(t *testing.T) {
  //Save will save the json marshal output of config into the file
  //named ~/.dirname/configFilename.json
  if uuid, err := CreateUUID(); err == nil {
    dir := fmt.Sprintf("%v-%v", "polka", uuid)
    if err := Create(dir); err != nil {
      t.Fail()
    }
    config := TestConfig {
        Name: "testConfig",
    }

    if err := Save(dir, "polka", config); err != nil {
      t.Fail()
    }
    if err := CleanUp(ExpandTilde(dir)); err != nil {
      t.Fail()
      t.Log(err)
    }
  }
}
