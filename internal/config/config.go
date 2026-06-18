package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisHost      string
	RedisPort      string
	GRPCPort       string
	BucketCapacity int
	RefillRate     int
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(" .env not found !!", err)
	}

	capacity, err := strconv.Atoi(getEnv("Bucket_Capacity", "100"))
	if err != nil {
		log.Fatal("Invalid Bucket_Capacity")
	}

	refillRate, err := strconv.Atoi(getEnv("Refill_Rate", "10"))
	if err != nil {
		log.Fatal("Invalid Refill_Rate")
	}

	return &Config{
		RedisHost:      getEnv("Redis_Host", "localhost"),
		RedisPort:      getEnv("Redis_Port", "6379"),
		GRPCPort:       getEnv("GRPCPort", "5005"),
		BucketCapacity: capacity,
		RefillRate:     refillRate,
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}
	return value
}
