```markdown
# ADR - Architecture Decision Record

**Título:**  
Escolha do Banco de Dados para Persistência de Transações ISO-8583

**Status:**  
Aceita

**Data:**  
23/04/2025

**Contexto:**  
Estamos desenvolvendo um sistema autorizador de transações financeiras em Go que processa mensagens ISO-8583 recebidas via TCP/IP. O sistema precisa persistir essas transações em um banco de dados para fins de registro, auditoria e consultas futuras.

**Problema a ser resolvido:**  
Precisamos escolher um banco de dados que atenda aos requisitos de armazenamento de transações financeiras, garantindo integridade, confiabilidade e desempenho adequado.

**Objetivos do time:**
- Armazenar transações financeiras de forma segura e confiável  
- Garantir consistência e integridade dos dados (propriedades ACID)  
- Possibilitar consultas eficientes para análises e auditoria  
- Manter um sistema resiliente mesmo em caso de falhas  
- Implementar uma solução adequada para uma prova de conceito, mas que possa escalar  

**Restrições:**
- Simplicidade de implementação como prioridade (por ser uma PoC)  
- Recursos de infraestrutura limitados  
- Integração fácil com Go e RabbitMQ  
- Preferência por tecnologias bem documentadas e com boa comunidade  

**Desafios técnicos:**
- Modelagem de dados com campos variáveis do ISO-8583  
- Balanceamento entre performance de escrita e capacidade de consulta  
- Garantir durabilidade das transações sem prejudicar latência  
- Equilibrar integridade rígida com flexibilidade de estrutura  

**Tecnologias em uso:**
- Go, RabbitMQ, Docker, bibliotecas ISO-8583 em Go

**Decisão:**  
Utilizar PostgreSQL como banco de dados para persistência das transações ISO-8583.

**Alternativas Consideradas:**  
- **PostgreSQL**: Suporte completo a ACID, excelente para queries e integridade de dados, mas exige tuning para alto volume.  
- **MongoDB**: Flexível para campos variáveis e escalável, porém com menor consistência transacional por padrão.  
- **SQLite**: Muito simples e leve para PoC, mas com sérias limitações de concorrência e escalabilidade.

**Justificativa:**  
PostgreSQL oferece a melhor combinação entre confiabilidade, consistência transacional (ACID), suporte a queries complexas e integração madura com o ecossistema Go. Embora exija tuning para alta performance, é uma base sólida mesmo para escalabilidade futura. Garante um caminho seguro para a evolução do sistema sem comprometer os objetivos de segurança e auditoria.

**Consequências:**  
- Modelagem relacional adaptada ao formato ISO-8583 com possível uso de JSONB  
- Requer atenção à criação de índices para manter desempenho de leitura  
- A equipe deve se capacitar em boas práticas de uso do PostgreSQL  
- Excelente compatibilidade com o stack atual (Go, Docker)  
- Mitiga riscos de perda de integridade ou inconsistência nos dados  
- Permite evolução da PoC para produção com menor refatoração futura

**Autores:**  
Bruno Gonzaga Santos

```