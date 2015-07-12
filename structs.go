package main

import "time"

type Bucket struct {
  Name         string     `json:"name"`
  CreationDate *time.Time `json:"subdomain"`
}

type R53 struct {
  Name   string  `json:"name"`
  Type   string  `json:"type"`
  Ttl    int64 `json:"ttl"`
  Value  string  `json:"value"`
}

type EC2InstanceTags struct {
  Name   string `json:"name"`
  Value  string `json:"value"`
}
