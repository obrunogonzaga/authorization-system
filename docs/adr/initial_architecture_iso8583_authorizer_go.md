# ADR - Architecture Decision Record

**Título:**  
Arquitetura Inicial do Autorizador ISO-8583 em Go

**Status:**  
Aceita

**Data:**  
23 de abril de 2025

**Contexto:**  
O cenário atual é a necessidade de desenvolver um sistema autorizador como prova de conceito que possa processar mensagens financeiras no formato ISO-8583. O sistema precisa receber essas mensagens via conexão TCP/IP, processar adequadamente e retornar respostas no mesmo padrão, além de persistir os dados em um banco de dados.

Os objetivos principais são:
- Receber e processar mensagens financeiras no padrão ISO-8583
- Implementar uma arquitetura modular e desacoplada
- Utilizar algum broker para processamento assíncrono de mensagens
- Garantir persistência dos dados em SQL ou NoSQL
- Fornecer respostas em formato ISO-8583 aos clientes
- Criar uma base sólida que possa ser expandida posteriormente

As restrições incluem:
- Por ser uma POC, priorizar simplicidade em vez de uma implementação completa
- Foco em demonstrar o fluxo completo ao invés de todos os detalhes do protocolo
- Uso de Go como linguagem principal

Desafios técnicos:
- Parser ISO-8583 robusto e flexível
- Balanceamento entre resposta rápida e persistência
- Definição de estrutura de filas no broker
- Modelagem correta da base de dados

Este é o primeiro projeto com essa proposta, sem histórico de tentativas anteriores.

**Decisão:**  
Adotar uma arquitetura modular em Go com uso de fila de mensagens (ex: RabbitMQ ou Redis Streams) e processamento assíncrono por meio de workers especializados.

**Alternativas Consideradas:**  
- Alternativa 1: Arquitetura monolítica simples com handlers ISO-8583  
  - Prós: fácil de implementar, menor curva de aprendizado  
  - Contras: baixa escalabilidade, difícil manutenção, pouco flexível  

- Alternativa 2: Microsserviços com broker de mensagens (ex: Kafka, NATS)  
  - Prós: alta escalabilidade, processamento assíncrono, extensível  
  - Contras: alta complexidade, mais dependências, overhead para POC  

- Alternativa 3: Arquitetura modular com workers e fila interna  
  - Prós: modularidade, desacoplamento, paralelismo com simplicidade  
  - Contras: exige bom desenho de filas e concorrência, complexidade moderada

**Justificativa:**  
A opção modular com workers e fila oferece equilíbrio entre simplicidade e robustez para a fase de POC. Evita o overhead de microsserviços, mas já introduz boas práticas como processamento assíncrono e desacoplamento. Permite validação ágil e fácil evolução futura. Reduz risco de bloqueios e favorece testes e manutenção.

**Consequências:**  
- Técnicas: modelagem robusta de filas, concorrência segura, parser testável  
- Organizacionais: equipe deve se familiarizar com eventos e workers  
- Operacionais: escalabilidade horizontal por worker, desacoplamento de camadas  
- Riscos Mitigados: evita bloqueios no fluxo e lock-in prematuro  
- Impacto no roadmap: base sólida para expansão e novas integrações futuras

**Autores:**  
Bruno Gonzaga Santos