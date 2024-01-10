package main

import (
  "fmt"
  "errors"
)

func commandMapf(cfg *config) error {
  locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsUrl)
  if err != nil {
    return err
  }
  cfg.nextLocatoinsUrl = locationsResp.Next
  fgc.prevousLocationUrl = locationsResp.Previous

  for _, loc := range locationsResp.Results {
    fmt.Println(loc.Name)
  }
  return nil
}

func commandMapb(cfg *config) error {
  if cfg.prevousLocationUrl == nil {
    return errors.New("You are already on the first page")
  }

  locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevousLocationUrl)
  if err != nil {
    return err
  }

  cfg.nextLocationsUrl = locationsResp.Next
  cfg.prevousLocationUrl = locationsResp.Previous

  for _, loc in range locationsResp.Results {
    fmt.Println(loc.Name)
  }
  return nil
}
