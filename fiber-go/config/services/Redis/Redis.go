package redis

import (
	"fmt"
	dotenv "marketplace/config/services/Dotenv"

	"github.com/garyburd/redigo/redis"
)

// Struct para representar conexão do redis
type RedisConnection struct {
	connection redis.Conn
}

// Conectar e retornar a conexão
func RedisConnect() (RedisConnection, error) {
	redisURL := fmt.Sprintf("%s:%s", dotenv.MyEnvironmentApp.Redis_host, dotenv.MyEnvironmentApp.Redis_port)

	connection, err := redis.Dial("tcp", redisURL)

	if err != nil {
		return RedisConnection{}, err
	}

	return RedisConnection{connection: connection}, nil
}

// Criar uma key ou value no redis
func (connec RedisConnection) Create(Key string, Value string) error {

	// Comando set vai fazer que vai setar um valor em uma chave
	_, err := connec.connection.Do("SET", Key, Value)
	if err != nil {
		return err
	}

	return nil
}

// Resgatar uma key ou value no redis
func (connec RedisConnection) Get(Key string) (string, error) {

	data, err := redis.Bytes(connec.connection.Do("GET", Key))

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Deletar uma value no redis
func (connec RedisConnection) Delete(Key string) error {
	_, err := connec.connection.Do("DEL", Key)
	if err != nil {
		return err
	}

	return nil
}

// Atualizar uma key ou value no redis
func (connec RedisConnection) Update(Key string, Value string) error {

	err := connec.Create(Key, Value)

	if err != nil {
		return err
	}

	return nil
}
