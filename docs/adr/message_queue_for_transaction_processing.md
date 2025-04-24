# ADR - Architecture Decision Record

**Título:**  
Mensageria assíncrona para aumento de resiliência e desempenho no sistema autorizador

**Status:**  
Aceita

**Data:**  
23 de abril de 2025

**Contexto:**  
O sistema autorizador precisa de um mecanismo de mensageria para desacoplar o processamento de transações ISO-8583 da resposta ao cliente, permitindo processamento assíncrono e maior resiliência.

**Cenário atual:**  
Estamos desenvolvendo um sistema autorizador como prova de conceito que precisa:
- Receber e enviar mensagens ISO-8583 via TCP/IP  
- Processar essas mensagens de forma eficiente  
- Garantir que transações sejam persistidas no banco de dados  

**Problema a ser resolvido:**  
- O processamento de transações financeiras não deve bloquear a resposta ao cliente  
- Precisamos garantir que nenhuma transação seja perdida em caso de falhas  
- O sistema deve ser capaz de lidar com picos de carga sem degradação  
- Diferentes tipos de processamento (persistência, notificações) precisam acontecer independentemente  

**Decisão:**  
Adotar o RabbitMQ como mecanismo de mensageria assíncrona no sistema autorizador.

**Alternativas Consideradas:**  
- **Apache Kafka:**  
  Prós: Alta escalabilidade, forte durabilidade, ecossistema robusto  
  Contras: Complexidade operacional, curva de aprendizado elevada  

- **RabbitMQ:**  
  Prós: Simples de configurar, bom roteamento de mensagens, leve para POC  
  Contras: Desempenho limitado sob carga extrema, menor tolerância a partições que Kafka  

- **Processamento direto e síncrono com retries internos:**  
  Prós: Baixa complexidade inicial, fácil de prototipar  
  Contras: Fortemente acoplado, baixo desempenho sob carga, menos resiliente  

**Justificativa:**  
RabbitMQ foi escolhido por oferecer um equilíbrio entre simplicidade e robustez, sendo adequado ao estágio atual de desenvolvimento (POC). Suporta múltiplos padrões de troca e garante a entrega de mensagens, com menor custo operacional comparado ao Kafka. Trade-offs incluem a necessidade de infraestrutura adicional e capacitação da equipe.

**Consequências:**  
- Desacoplamento do fluxo principal com aumento de resiliência  
- Processos independentes para persistência e notificações  
- Introdução de um novo componente na infraestrutura (RabbitMQ)  
- Requisitos de observabilidade e confiabilidade das filas  
- Roadmap futuro inclui retry policies, dead-letter queues e possível reavaliação da solução conforme escala  

**Autores:**  
Bruno Gonzaga Santos