package redis

import (
	dotenv "marketplace/config/services/Dotenv"
	"testing"

	"github.com/garyburd/redigo/redis"
)

// Teste do crud do redis
func TestRedisConnection(t *testing.T) {
	dotenv.LoadEnv()

	connection, err := RedisConnect()

	if err != nil {
		t.Fatal(err.Error())
	}

	defer connection.connection.Close()

	_, err = connection.connection.Do("SET", "Teste", []byte("Teste legal man"))

	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = redis.Bytes(connection.connection.Do("GET", "Teste"))

	if err != nil {
		t.Fatal(err.Error())
	}

}

type UseTestCrudRedis struct {
	Key   string
	Value string
}

var MyTestCrud = [2]UseTestCrudRedis{
	{
		Key:   "json",
		Value: "{\"id\":12}",
	},
	{
		Key:   "word",
		Value: "HELLO WORLD",
	},
}

// Test do crud do service do redis, baseado em varios dados
func TestCrudRedis(t *testing.T) {
	dotenv.LoadEnv()

	connection, err := RedisConnect()

	if err != nil {
		t.Fatal(err.Error())
	}

	defer connection.connection.Close()

	for _, TestActual := range MyTestCrud {

		t.Logf("Executando create\nValue: %s| Key: %s", TestActual.Value, TestActual.Key)
		err := connection.Create(TestActual.Key, TestActual.Value)

		if err != nil {
			t.Errorf("Erro na criação\nR: %s\nData: key %s value %s", err.Error(), TestActual.Key, TestActual.Value)
		}

		t.Logf("Executando get\nValue: %s| Key: %s", TestActual.Value, TestActual.Key)
		data, err := connection.Get(TestActual.Key)

		if err != nil {
			t.Errorf("Erro no get\nR: %s\nData: key %s value %s", err.Error(), TestActual.Key, TestActual.Value)
		} else if data != TestActual.Value {
			t.Errorf("Utilização do get deu erro\n R: value resgatado %s\n R: value esperado %s", data, TestActual.Value)
		}

		t.Logf("Executando update\nValue: %s| Key: %s", TestActual.Value, TestActual.Key)
		err = connection.Update(TestActual.Key, TestActual.Value)

		if err != nil {
			t.Errorf("Erro na atualização\nR: %s\nData: key %s value %s", err.Error(), TestActual.Key, TestActual.Value)
		}

		t.Logf("Executando get pós update\nValue: %s| Key: %s", TestActual.Value, TestActual.Key)
		data, err = connection.Get(TestActual.Key)

		if err != nil {
			t.Errorf("Erro no get após atualização\nR: %s\nData: key %s value %s", err.Error(), TestActual.Key, TestActual.Value)
		} else if data != TestActual.Value {
			t.Errorf("Utilização do get deu erro após atualização\n R: value resgatado %s\n R: value esperado %s", data, TestActual.Value)
		}

		t.Logf("Executando Delete\nValue: %s| Key: %s", TestActual.Value, TestActual.Key)
		err = connection.Delete(TestActual.Key)

		if err != nil {
			t.Errorf("Erro no delete\nR: %s\nData: key %s value %s", err.Error(), TestActual.Key, TestActual.Value)
		}

		t.Logf("Executando get pós update\nValue: %s| Key: %s", TestActual.Value, TestActual.Key)
		_, err = connection.Get(TestActual.Key)

		if err == nil {
			t.Errorf("Erro no get pós delete\nR: Analisar se não foi deletado\nR: %s", err.Error())
		}
	}
}
