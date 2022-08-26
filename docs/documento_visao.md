# Documento de Visão

## Equipe e Definição de Papéis

Membro  |     Papel   |   E-mail   |
------- | ----------- | ---------- |
Brenno  | Gerente, Analista, Testador, Desenvolvedor | brennovictoralves@gmail.com
Lucas   | Gerente, Analista, Testador, Desenvolvedor | lucasvinicius0511@gmail.com
Felipe  | Gerente, Analista, Testador, Desenvolvedor | felipedavid.huh@gmail.com
Relyson | Gerente, Analista, Testador, Desenvolvedor | relyson.m@gmail.com

### Matriz de Competências

Membro     |     Competências   |
---------  | ----------- |
Brenno    | Desenvolvedor Python e C |
Lucas     | Desenvolvedor Python e Javascript, Designer |
Felipe    | Desenvolvedor C e Go |
Relyson   | Desenvolvedor C e Javascript, HTML/CSS |

## Perfis dos Usuários

O sistema poderá ser utilizado por diversos usuários. Temos os seguintes perfis/atores:

Perfil                                 | Descrição   |
---------                              | ----------- |
Administrador | Este usuário pode realizar os cadastros de jogos e editar comentários
Usuário Padrão | Este usuário pode checar a avaliação dos jogos, avaliá-los, adicionar comentários e visualizar críticas.
Usuário Crítico | É um usuário padrão com privilégios de adicionar críticas a jogos.
Visitante | Pode-se apenas visualizar a avaliação dos jogos, comentários e críticas.

## Lista de Requisitos Funcionais

Requisito                                 | Descrição   | Ator |
---------                                 | ----------- | ---------- |
RF001 - Adicionar jogo     | Um jogo terá os atributos nome, avaliação, descrição, críticas, comentários, data de lançamento.  | Administrador |
RF002 - Alterar jogo | Atualizar os atributos do jogo | Administrador |
RF003 - Excluir jogo | Remove o jogo do sistema | Administrador |
RF004 - Listar jogos | Lista todos os jogos do sistema | Usuário Padrão |
RF005 - Adicionar usuário | Um usuário possui username, email e senha | Usuário Padrão |
RF006 - Alterar usuário | Altera os dados do usuário (email, nome, senha) | Usuário Padrão |
RF007 - Consultar Usuário | Exibe as informações do usuário (email, nome, senha) | Usuário Padrão |
RF008 - Excluir Usuário | Exclui o usuário do sistema | Administrador |
RF009 - Desativar Usuário | Desativa a conta do usuário por tempo determinado | Usuário Padrão |
RF010 - Adicionar Comentário | Adiciona um comentário a um jogo | Usuário Padrão | 
RF011 - Alterar Comentário | Altera o comentário do usuário | Usuário Padrão |
RF012 - Excluir Comentário | Exclui o comentário do usuário | Usuário Padrão |
RF013 - Adicionar Crítica | Adiciona uma crítica a um jogo | Usuário Crítico |
RF014 - Excluir Crítica | Exclui uma crítica de um jogo | Usuário Crítico |
RF015 - Alterar Crítica | Alterar uma crítica de um jogo | Usuário Crítico |
RF016 - Adicionar Avaliação | Os usuários podem adicionar avaliações de 0 a 5 estrelas ao jogo | Usuário Padrão | 
RF017 - Alterar Avaliação | Os usuários podem alterar as avaliações de 0 a 5 estrelas ao jogo | Usuário Padrão | 

