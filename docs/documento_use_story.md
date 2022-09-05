# Documento lista de User Stories

## Histórico de Revisão

Data | Versão |  Descrição |  Autor
---- | ------ | ---------- | -----
02/09/2022 | 0.0.1 | Detalhamento do User Story RF001 | Lucas

### User Story 001 - Manter jogo:

Descrição | O Sistema deve manter um cadastro do jogo. O jogo tem os atributos id, nome, descrição, data de lançamento. O administrador será quem poderá cadastrar um jogo, alterar ou remover, e usuário padrão/crítico poderão consultar os jogos.
--------- | -----------------------------------------------

Requisitos Envolvidos |       |
--------------------- | -------
RF001 | Cadastrar Jogo|
RF002 | Alterar Jogo|
RF003 | Visualizar Jogo |
RF004 | Excluir Jogo  |

Prioridade | Essencial
---------- | --------
Estimativa | 3h
Tempo Gasto (real): | ?
Tamanho Funcional | ?
Analista | Lucas
Desenvolvedor | Felipe
Revisor | Brenno
Testador | Relyson

### Testes de aceitação (TA) 
Código | Descrição
-------|----------
TA02.01 | O administrador precisa adicionar/atualizar/visualizar/deletar um jogo no sistema, então ele vai em "Cadastrar jogo”. Então será exibido uma tela onde terá uma tabela com os jogos já cadastrados e os botões de Adicionar, atualizar, visualizar e deletar. Caso não tenha nenhum cadastro, será exibida uma mensagem informando que não há nada cadastrado.
TA02.02| Ao clicar em "Adicionar Jogo" uma tela de cadastro será exibida com todos os campos para serem preenchidos. Logo após preencher as informações, o administrador clica em "Salvar", e em seguida é levado de volta à tela da tabela dos jogos cadastrados.
TA02.03| O Administrador não preenche todos os campos obrigatórios. Então na tela deverá ser exibida uma mensagem informando que o usuário não preencheu todos os campos obrigatórios.
TA02.04| O usuário deseja ver os dados do jogo cadastrado com mais detalhes, então ele clica no banner do jogo, onde será exibida uma tela com todos os dados do jogo.
TA02.05| O administrador deseja atualizar os dados de um cadastro. Então ele clica em "atualizar", onde será exibida uma tela com todos os dados aptos a serem atualizados. O administrador faz a atualização desejada e depois clica em "Salvar" onde sua alteração será salva e o administrador será levado de volta à tela principal(de cadastro).

TA01.06| No ato da edição, o administrador deixa um campo obrigatório em branco, então o sistema não permite a alteração e informa na tela que todos os campos precisam ser preenchidos corretamente.
TA01.07| O administrador deseja deletar um jogo previamente cadastrado, então ele clica em "Deletar". Então, o cadastro do jogo é apagado instantaneamente do sistema.


### User Story 002 -Manter Usuário:

Descrição | O sistema deve manter um usuário cadastrado, consistindo de seu nome, email, senha e data de nascimento.
--------- | -----------------------------------------------

Requisitos Envolvidos |       |
--------------------- | -------
RF001 | Cadastrar usuário |
RF002 | Login  |
RF003 | Logout |
RF004 | Modificar informações do usuário |

Prioridade | Essencial
---------- | --------
Estimativa | 10h
Tempo Gasto (real): | 5h
Tamanho Funcional | ?
Analista | Felipe
Desenvolvedor | Brenno
Revisor | Relyson
Testador | Lucas

### Testes de aceitação (TA) 
Código | Descrição
-------|----------
TA01.01| O usuário que não está logado e não possui conta no sistema deve se cadastrar antes de poder avaliar jogos e fazer comentários. Para realizar o cadastro o usuário irá clicar no botão “Cadastrar” no canto superior direito da tela, onde será direcionado para um formulário e irá informar seu nome, email e senha. Após informar os dados de maneira correta e enviar o formulário, uma conta será criada e o usuário será direcionado para a página principal do site, já logado.
TA01.02| O usuário que já está cadastrado no site pode fazer login clicando no botão “Login” no canto superior direito da tela, onde irá informar o seu email e senha.
TA01.03| Para fazer logout o usuário precisa já estar logado, depois disso o mesmo irá clicar no botão “Logout” no canto superior direito da tela, onde sua sessão será encerrada e será redirecionado para a página principal do site.
TA01.03| O usuário já logado pode atualizar suas informações como nome, email e senha, indo para o painel de controle da conta clicando no seu ícone de perfil.


### User Story 003 - Manter Avaliação:

Descrição | O Sistema deve manter uma avaliação de jogo. O jogo tem os atributos id, nome, descrição, data de lançamento. O usuário poderá avaliar um jogo.
--------- | -----------------------------------------------

Requisitos Envolvidos |       |
--------------------- | -------
RF001 | Cadastrar avaliação |
RF002 | Alterar avaliação  |
RF003 | Listar avaliação |

Prioridade | Essencial
---------- | --------
Estimativa | 10h
Tempo Gasto (real): | 5h
Tamanho Funcional | ?
Analista | Relyson
Desenvolvedor | Lucas
Revisor | Felipe
Testador | Brenno

### Testes de aceitação (TA) 
Código | Descrição
-------|----------
TA01.01 | O usuário pode avaliar um jogo, precisará clicar em um “botão” de avaliação para poder dar uma nota de 0 a 5 e em seguida confirmando a avaliação.
TA01.02| O usuário pode visualizar a média da avaliação de um jogo, ficando disponível ao acessar um jogo.
TA01.03| A avaliação de um jogo será calculada a partir da média de todas as avaliações dos usuários em cima daquele jogo.
TA01.04| O usuário pode alterar a avaliação de um jogo, precisará clicar em um “botão” de avaliação para poder dar uma nova nota de 0 a 5 e em seguida confirmando a alteração de nota.


### User Story 004 - Manter Comentário:

Descrição | O Sistema deve manter os comentários de jogos. O comentário tem os atributos id, conteúdo(descrição) e data. O usuário poderá realizar comentários em jogos.
--------- | -----------------------------------------------

Requisitos Envolvidos |       |
--------------------- | -------
RF010 | Adicionar Comentário|	
RF011 | Alterar Comentário|
RF012 | Excluir Comentário|


Prioridade | Essencial
---------- | --------
Estimativa | 10h
Tempo Gasto (real): | 5h
Tamanho Funcional | ?
Analista | Brenno
Desenvolvedor | Relyson
Revisor | Lucas
Testador | Felipe

### Testes de aceitação (TA) 
Código | Descrição
-------|----------
TA01.01 | O usuário pode realizar comentários em jogos. Ele precisa clicar em algum jogo (banner) de sua escolha, e então as informações do jogo irão aparecer na sua tela, com as opções (botões) de “avaliar” e “comentar”.
TA01.02| Ao clicar em “comentar”, irá abrir uma caixa de texto, no qual o usuário poderá digitar o seu comentário, e ao clicar em “enviar”, o comentário será salvo.
TA01.03| Ao enviar o comentário, o usuário poderá visualizar seus comentários feitos no topo dos comentários do jogo, e os comentários de outros usuários logo abaixo. Os comentários do usuário possuem a opção (botão) de “editar”. Ao clicar em editar, o usuário poderá alterar o seu comentário, e ao clicar em “enviar”, a alteração realizada será salva.
TA01.04| Os comentários do usuário possuem a opção (botão) de "excluir". Ao clicar em “excluir”, irá abrir a opção de confirmar a exclusão com a descrição: “Deseja excluir o comentário” e os botões “sim”, e “cancelar”. Ao clicar em “sim”, o comentário será excluído. Ao clicar em “cancelar”, a exclusão será cancelada.
