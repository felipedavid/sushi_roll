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

## Lista de Requisitos Não-Funcionais

Requisito                                 | Descrição   |
---------                                 | ----------- |
RNF001 - Logs | Toda requisição feita ao servidor web deve ser logada, como também erros e crashs |
RNF002 - Configurações | Coisas como porta do servidor web, endereço do banco de dados, devem ser passados por argumentos de linha de comando ao iniciar a aplicação |
RNF003 - Suporte para mobile | A aplicação deve abrir de forma responsiva em browsers mobile |

## Riscos

Tabela com o mapeamento dos riscos do projeto, as possíveis soluções e os responsáveis.

Data | Risco | Prioridade | Responsável | Status | Providência/Solução |
------ | ------ | ------ | ------ | ------ | ------ |
26/08/2022 | Não aprendizado das ferramentas utilizadas pelos componentes do grupo | Alta | Todos | Vigente | Reforçar estudos sobre as ferramentas e aulas com a integrante que conhece a ferramenta |
26/08/2022 | Ausência por qualquer motivo do cliente | Média | Gerente | Vigente | Planejar o cronograma tendo em base a agenda do cliente |
26/08/2022 | Divisão de tarefas mal sucedida | Baixa | Gerente | Vigente | Acompanhar de perto o desenvolvimento de cada membro da equipe |
26/08/2022 | Implementação de protótipo com as tecnologias | Alto | Todos | Resolvido | Encontrar tutorial com a maioria da tecnologia e implementar um caso base do sistema |



