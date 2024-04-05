# p2-mod9

Implementar um Producer (Produtor): Deve coletar dados simulados de sensores de qualidade do ar e publicá-los em um tópico do Kafka chamado qualidadeAr. Os dados devem incluir:
Id do sensor, timestamp, tipo de poluente e nivel da medida.

Implementar um Consumer (Consumidor): Deve assinar o tópico qualidadeAr e processar os dados recebidos, exibindo-os em um formato legível, além de armazená-los para análise posterior (escolha a forma de armazenamento que achar mais adequada).

Implementar testes de Integridade e Persistência: Criar testes automatizados que validem a integridade dos dados transmitidos (verificando se os dados recebidos são iguais aos enviados) e a persistência dos dados (assegurando que os dados continuem acessíveis para consulta futura, mesmo após terem sido consumidos).

# Como executar?

Em primeiro lugar, para rodar o kafka localmente é necessário executar o comando:
```
docker compose up -d
```

Ademais, para executar o producer e consumer:
```
go mod tidy

<!-- producer -->
cd producer
go run producer.go

<!-- consumer -->
cd consumer
go run consumer.go
```

Por fim, para rodar os testes é necessário executar:
```
cd testes
go test -v
```

# Demonstração
As imagens se encontram na pasta `/demonstração`.
