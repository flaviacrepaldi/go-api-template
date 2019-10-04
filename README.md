# go-api-template
Exemplo de estrutura básica de api em Go, utilizando Docker e Mongo.

Para começar, iniciar os services utilizando o Docker. Para isso, dentro da pasta do projeto:

<b>docker-compose up -d --build</b>

Para acessar o container da api:

<b>docker exec -it api bash</b>

Para monitorar erros no container:

<b>docker logs api --follow</b>

# endpoints

<b>POST</b> <br>
<i>/post-endpoint</i><br>

Registra no mongo os campos <b>first_beautiful_field</b> e <b>second_beautiful_field</b>, sendo que o <b>second_beautiful_field</b> registra a data atual. <br>

  { <br>
    &nbsp;&nbsp;&nbsp;&nbsp;"first_beautiful_field": "my beautiful string" <br>
  } <br>


<b>GET</b> <br>
<i>/get-endpoint</i>

Retorna os registros do mongo, ordenados pelo campo <b>second_beautiful_field</b>
