package main

import (
    "fmt"
)

type Astro struct {
    name      string
    age       int
    mission   string
    isneeeded bool
}

type nasaMission struct {
    people  []Astro
    number  int
    message string
}

func main() {
    astro1 := Astro{"John", 23, "DOJ", false}
    astro2 := Astro{"Mike", 22, "DOD", true}
    astro3 := Astro{"Tim", 24, "DOE", false}

    fmt.Println(astro1)
    fmt.Println(astro2)
    fmt.Println(astro3)

    astroSlice := []Astro{astro1, astro2, astro3}

    fmt.Println(astroSlice)

    fmt.Println(astroSlice[2].mission)

    missionSlice := nasaMission{astroSlice, 3, "success"}

    fmt.Println(missionSlice)

    fmt.Printf("%+v", missionSlice)

}
