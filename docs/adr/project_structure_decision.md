# ADR - Architecture Decision Record

**Título:**  
Organização Modular de Diretórios em Sistema com RabbitMQ e ISO-8583

**Status:**  
Aceita

**Data:**  
23 de abril de 2025

**Contexto:**  
Estamos desenvolvendo um sistema autorizador em Go que processará mensagens ISO-8583, utilizará RabbitMQ para mensageria assíncrona e PostgreSQL para persistência. Precisamos definir uma estrutura de pastas que suporte adequadamente o desenvolvimento, manutenção e evolução do projeto.

**Problema enfrentado:**  
Determinar a organização do código e recursos do projeto de forma a promover:
- Clareza e legibilidade
- Separação adequada de responsabilidades
- Facilidade de manutenção e testabilidade
- Suporte para evolução futura

**Objetivos a serem atendidos:**  
- Definir uma estrutura que reflita a arquitetura modular com workers e filas  
- Facilitar a localização de código e recursos  
- Promover práticas de código limpo e desacoplamento  
- Suportar testes unitários e de integração  
- Permitir expansão ordenada do sistema  

**Restrições envolvidas:**  
- Priorizar convenções amplamente aceitas na comunidade Go  
- Balancear complexidade e clareza, dado que é uma prova de conceito  
- Facilitar o trabalho para desenvolvedores com diferentes níveis de experiência  
- Considerar as necessidades do sistema distribuído com RabbitMQ  

**Desafios técnicos já enfrentados:**  
- Determinar o nível adequado de modularização sem sobrecomplicar  
- Equilibrar entre módulos focados em domínio vs. focados em tecnologia  
- Definir limites claros entre componentes do sistema  
- Decidir como organizar código compartilhado vs. específico  
- Estruturar de forma que facilite a evolução para microsserviços no futuro, se necessário  

**Histórico da arquitetura ou decisões anteriores relevantes:**  
- Decisão inicial de utilizar uma arquitetura modular em Go com workers e fila interna  
- Escolha do RabbitMQ como sistema de mensageria  
- Adoção do PostgreSQL como banco de dados  
- Padrões de estruturação de projetos Go foram revisados, incluindo:  
  - Padrão de repositório Go padrão  
  - Abordagem de camadas Clean Architecture  
  - Estruturação por domínio vs. por funcionalidade  
  - Precedentes de projetos similares na indústria  

**Decisão:**  
Adotar a estrutura baseada por domínio (Domain-Driven Package Layout)

**Alternativas Consideradas:**  
- **Alternativa 1:** Estrutura baseada em camadas (Clean Architecture)  
  _Prós:_ Alta separação de responsabilidades, fácil substituição de dependências, boa para futura evolução.  
  _Contras:_ Complexidade maior, pode ser excessivo para uma POC.  

- **Alternativa 2:** Estrutura baseada por domínio (Domain-Driven Package Layout)  
  _Prós:_ Reflete o negócio, modularidade clara, facilita onboarding.  
  _Contras:_ Pode misturar infraestrutura e domínio, exige disciplina.  

- **Alternativa 3:** Estrutura por tipo de arquivo (Go idiomática)  
  _Prós:_ Familiaridade com padrão Go, baixa complexidade, muito documentada.  
  _Contras:_ Diluição da coesão funcional, menos aderente a arquitetura explícita.  

**Justificativa:**  
A estrutura baseada por domínio oferece o melhor equilíbrio entre clareza organizacional e suporte à evolução do sistema. Favorece a modularização por contexto de negócio, promove boas práticas de separação de responsabilidades e mantém a estrutura acessível a desenvolvedores com diferentes níveis de experiência. Permite também uma transição mais natural para microsserviços, se necessário.

**Consequências:**  
- Facilita testes e manutenção isolada por domínio  
- Suporte a CI/CD modular e pipelines por contexto  
- Organiza o trabalho por funcionalidades, o que ajuda na divisão por times  
- Mitiga riscos de acoplamento entre camadas técnicas  
- Garante escalabilidade da estrutura para futuras evoluções

**Autores:**  
Bruno Gonzaga Santos