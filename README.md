## Sushi roll - Aplicação para avaliação de jogos

Aplicação que permite usuários avaliar jogos e consultar opinião de outros usuários e críticos sobre jogos.

* A tecnologia escolhida para o projeto foi a linguagem **Go**, utilizando a biblioteca padrão e third-party packages 
para funcionalidades não fornecidas nela.

* Tutoriais da tecnologia escolhida:

https://go-tour-br.appspot.com/basics/1  
https://www.youtube.com/watch?v=WiGU_ZB-u0w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg  
https://medium.com/baixada-nerd/criando-um-crud-simples-em-go-3640d3618a67  


## Documentos

[Documento de visão](https://github.com/felipedavid/sushi_roll/blob/main/docs/visao.md)<br/>
[Documento de User Story](https://github.com/felipedavid/sushi_roll/blob/main/docs/user_story.md)<br/>
[Diagrama de dados](https://github.com/felipedavid/sushi_roll/blob/main/docs/modelo_de_dados.md)<br/>
[Documento de Iteração](https://github.com/felipedavid/sushi_roll/blob/main/docs/iteracao.md)<br/>


## Running ([golang-migrate](https://github.com/golang-migrate/migrate/releases) e [Go](https://go.dev/dl/) precisam estar instalados)
```console
# Preparando banco de dados
$ make postgre
$ make createuser
$ make createdb
$ make migrateup

# Iniciando servidor
$ make run
```

## API
[Coleção Postman](https://mega.nz/file/xVAzgQqA#uaSTeDIqJ0T7rrPcXVCIBpxB8t_a6GdwcP_lEk-TKxM)