package services

import (
    "context"
    "encoding/json"
    "excel-import-api/models"
    "github.com/go-redis/redis/v8"
    "strconv"
)

// RedisService handles interactions with Redis
type RedisService interface {
    CacheEmployee(employee models.Employee) error
    CacheEmployees(employees []models.Employee) error
    GetEmployeesFromCache() ([]models.Employee, error)
    RemoveEmployeeFromCache(ctx context.Context, id string) error

    ClearCache(ctx context.Context) error                       
}

// redisService implements RedisService
type redisService struct {
    rdb *redis.Client
}

// NewRedisService creates a new instance of RedisService
func NewRedisService() RedisService {
    // Initialize Redis client
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    return &redisService{rdb: rdb}
}

// CacheEmployee caches a single employee in Redis
func (rs *redisService) CacheEmployee(employee models.Employee) error {
    // Serialize employee struct to JSON
    empJSON, err := json.Marshal(employee)
    if err != nil {
        return err
    }

    // Cache employee data in Redis
    ctx := context.Background() // Create context
    err = rs.rdb.Set(ctx, "employee:"+strconv.Itoa(int(employee.ID)), empJSON, 0).Err()
    if err != nil {
        return err
    }

    return nil
}

// CacheEmployees caches a slice of employees in Redis
func (rs *redisService) CacheEmployees(employees []models.Employee) error {
    // Iterate over each employee and cache it individually
    ctx := context.Background()
    for _, emp := range employees {
        // Serialize employee struct to JSON
        empJSON, err := json.Marshal(emp)
        if err != nil {
            return err
        }

        // Cache employee data in Redis with a unique key
        key := "employee:" + strconv.Itoa(int(emp.ID))
        err = rs.rdb.Set(ctx, key, empJSON, 0).Err()
        if err != nil {
            return err
        }
    }

    return nil
}


// GetEmployeesFromCache retrieves employees data from Redis cache
func (rs *redisService) GetEmployeesFromCache() ([]models.Employee, error) {
    // Get employees data from Redis cache
    ctx := context.Background() // Create context
    employeesJSON, err := rs.rdb.Get(ctx, "employees").Bytes()
    if err != nil {
        return nil, err
    }

    // Deserialize employees JSON into slice of employee models
    var employees []models.Employee
    err = json.Unmarshal(employeesJSON, &employees)
    if err != nil {
        return nil, err
    }

    return employees, nil
}

// RemoveEmployeeFromCache removes employee data from Redis cache
func (rs *redisService) RemoveEmployeeFromCache(ctx context.Context, id string) error {
    // Delete employee data from Redis cache
    if err := rs.rdb.Del(ctx, "employee:"+id).Err(); err != nil {
        return err
    }
    return nil
}

// ClearCache clears all data from the Redis cache
func (rs *redisService) ClearCache(ctx context.Context) error {
    // Flush all keys from the Redis cache
    if err := rs.rdb.FlushAll(ctx).Err(); err != nil {
        return err
    }
    return nil
}


